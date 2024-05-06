package p2p

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

// Message struct
type Message struct {
	ID   uint64 `json:"id"`
	Code uint   `json:"code"`
	Want uint   `json:"want,omitempty"`
	Data []byte `json:"data"`
}

var PeerIdList []string
var PeerAddrList []string
var Ctxt context.Context
var HostPeerAddr string
var NODE host.Host

func SendMessage(stream network.Stream, msg Message) {
	// Serialize the message
	encodedMsg, err := SerializeMessage(msg)
	if err != nil {
		fmt.Println("Error serializing message:", err)
		return
	}

	// Write the serialized message to the stream
	_, err = stream.Write(encodedMsg)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

// SerializeMessage serializes the Message struct into a byte slice
func SerializeMessage(msg Message) ([]byte, error) {
	// Encode the message struct
	encodedMsg, err := rlp.EncodeToBytes(msg)
	if err != nil {
		return nil, err
	}
	return encodedMsg, nil
}

func DeserializeMessage(encodedMsg []byte, msg *Message) error {
	err := rlp.DecodeBytes(encodedMsg, msg)
	if err != nil {
		return err
	}
	return nil

}

func ReceiveMessage(stream network.Stream) (Message, error) {
	var msg Message

	// Read the bytes from the stream
	buf := make([]byte, 1024) // Adjust the buffer size as needed
	n, err := stream.Read(buf)
	if err != nil {
		return msg, err
	}

	// Deserialize the message using your custom deserialization function
	err = DeserializeMessage(buf[:n], &msg)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

func sendMessageWithCTX(ctx context.Context, host host.Host, peerID peer.ID, helloMessage Message) {
	stream, err := host.NewStream(ctx, peerID, "/Ping")
	if err != nil {
		fmt.Println("Error opening stream to peer:", err)
		return
	}
	defer stream.Close()
	SendMessage(stream, helloMessage)
}

func StartNewNode() (host.Host, error) {

	host2, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))

	if err != nil {
		return nil, err
	}

	fmt.Println("Addresses:", host2.Addrs())
	fmt.Println("ID:", host2.ID())
	// fmt.Println("Peer_ADDR:", os.Getenv("PEER_ADDR"))
	hostAddr := host2.Addrs()[0].String()
	peerID := host2.ID()
	peerAddr := hostAddr + "/p2p/" + peerID.String()
	fmt.Println("Host_ADDR:", peerAddr)
	HostPeerAddr = peerAddr
	return host2, nil
}
