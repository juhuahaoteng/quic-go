package handshake

import (
	"crypto/tls"
	"fmt"

	"github.com/lucas-clemente/quic-go/fuzzing/internal/helper"
	"github.com/lucas-clemente/quic-go/internal/handshake"
	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/internal/testdata"
	"github.com/lucas-clemente/quic-go/internal/utils"
	"github.com/lucas-clemente/quic-go/internal/wire"
)

type messageType uint8

// TLS handshake message types.
const (
	typeClientHello         messageType = 1
	typeServerHello         messageType = 2
	typeNewSessionTicket    messageType = 4
	typeEncryptedExtensions messageType = 8
	typeCertificate         messageType = 11
	typeCertificateRequest  messageType = 13
	typeCertificateVerify   messageType = 15
	typeFinished            messageType = 20
)

func (m messageType) String() string {
	switch m {
	case typeClientHello:
		return "ClientHello"
	case typeServerHello:
		return "ServerHello"
	case typeNewSessionTicket:
		return "NewSessionTicket"
	case typeEncryptedExtensions:
		return "EncryptedExtensions"
	case typeCertificate:
		return "Certificate"
	case typeCertificateRequest:
		return "CertificateRequest"
	case typeCertificateVerify:
		return "CertificateVerify"
	case typeFinished:
		return "Finished"
	default:
		return fmt.Sprintf("unknown message type: %d", m)
	}
}

type chunk struct {
	data     []byte
	encLevel protocol.EncryptionLevel
}

type stream struct {
	chunkChan chan<- chunk
	encLevel  protocol.EncryptionLevel
}

func (s *stream) Write(b []byte) (int, error) {
	data := append([]byte{}, b...)
	select {
	case s.chunkChan <- chunk{data: data, encLevel: s.encLevel}:
	default:
		panic("chunkChan too small")
	}
	return len(b), nil
}

func initStreams() (chan chunk, *stream /* initial */, *stream /* handshake */) {
	chunkChan := make(chan chunk, 10)
	initialStream := &stream{chunkChan: chunkChan, encLevel: protocol.EncryptionInitial}
	handshakeStream := &stream{chunkChan: chunkChan, encLevel: protocol.EncryptionHandshake}
	return chunkChan, initialStream, handshakeStream
}

type handshakeRunner interface {
	OnReceivedParams(*wire.TransportParameters)
	OnHandshakeComplete()
	OnError(error)
	DropKeys(protocol.EncryptionLevel)
}

type runner struct {
	role    string // only used for logging
	errored bool

	client, server *handshake.CryptoSetup
}

var _ handshakeRunner = &runner{}

func newRunner(client, server *handshake.CryptoSetup, role string) *runner {
	return &runner{role: role, client: client, server: server}
}

func (r *runner) OnReceivedParams(*wire.TransportParameters) {}
func (r *runner) OnHandshakeComplete()                       {}
func (r *runner) OnError(err error) {
	r.errored = true
	(*r.client).Close()
	(*r.server).Close()
}
func (r *runner) DropKeys(protocol.EncryptionLevel) {}

const alpn = "fuzzing"

func toEncryptionLevel(n uint8) protocol.EncryptionLevel {
	switch n % 3 {
	default:
		return protocol.EncryptionInitial
	case 1:
		return protocol.EncryptionHandshake
	case 2:
		return protocol.Encryption1RTT
	}
}

// PrefixLen is the number of bytes used for configuration
const PrefixLen = 2

// Fuzz fuzzes the TLS 1.3 handshake used by QUIC.
//go:generate go run ./cmd/corpus.go
func Fuzz(data []byte) int {
	if len(data) < PrefixLen {
		return -1
	}
	enable0RTTClient := helper.NthBit(data[0], 0)
	enable0RTTServer := helper.NthBit(data[0], 1)
	useSessionTicketCache := helper.NthBit(data[0], 2)
	messageToReplace := data[1] % 32
	messageToReplaceEncLevel := toEncryptionLevel(data[1] >> 6)
	data = data[PrefixLen:]

	clientConf := &tls.Config{
		ServerName: "localhost",
		NextProtos: []string{alpn},
		RootCAs:    testdata.GetRootCA(),
	}
	if useSessionTicketCache {
		clientConf.ClientSessionCache = tls.NewLRUClientSessionCache(5)
	}
	cChunkChan, cInitialStream, cHandshakeStream := initStreams()
	var client, server handshake.CryptoSetup
	clientRunner := newRunner(&client, &server, "client")
	client, _ = handshake.NewCryptoSetupClient(
		cInitialStream,
		cHandshakeStream,
		protocol.ConnectionID{},
		nil,
		nil,
		&wire.TransportParameters{},
		clientRunner,
		clientConf,
		enable0RTTClient,
		utils.NewRTTStats(),
		nil,
		utils.DefaultLogger.WithPrefix("client"),
	)

	sChunkChan, sInitialStream, sHandshakeStream := initStreams()
	config := testdata.GetTLSConfig()
	config.NextProtos = []string{alpn}
	serverRunner := newRunner(&client, &server, "server")
	server = handshake.NewCryptoSetupServer(
		sInitialStream,
		sHandshakeStream,
		protocol.ConnectionID{},
		nil,
		nil,
		&wire.TransportParameters{},
		serverRunner,
		config,
		enable0RTTServer,
		utils.NewRTTStats(),
		nil,
		utils.DefaultLogger.WithPrefix("server"),
	)

	if len(data) == 0 {
		return -1
	}

	serverHandshakeCompleted := make(chan struct{})
	go func() {
		defer close(serverHandshakeCompleted)
		server.RunHandshake()
		// TODO: send session ticket
	}()

	clientHandshakeCompleted := make(chan struct{})
	go func() {
		defer close(clientHandshakeCompleted)
		client.RunHandshake()
	}()

	done := make(chan struct{})
	go func() {
		<-serverHandshakeCompleted
		<-clientHandshakeCompleted
		close(done)
	}()

messageLoop:
	for {
		select {
		case c := <-cChunkChan:
			b := c.data
			encLevel := c.encLevel
			if len(b) > 0 && b[0] == messageToReplace {
				fmt.Println("replacing message to the server", messageType(b[0]).String())
				b = data
				encLevel = messageToReplaceEncLevel
			}
			server.HandleMessage(b, encLevel)
		case c := <-sChunkChan:
			b := c.data
			encLevel := c.encLevel
			if len(b) > 0 && b[0] == messageToReplace {
				fmt.Println("replacing message to the client", messageType(b[0]).String())
				b = data
				encLevel = messageToReplaceEncLevel
			}
			client.HandleMessage(b, encLevel)
		case <-done: // test done
			break messageLoop
		}
		if clientRunner.errored || serverRunner.errored {
			break messageLoop
		}
	}

	<-serverHandshakeCompleted
	<-clientHandshakeCompleted

	return 1
}
