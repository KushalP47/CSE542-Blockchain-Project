package p2p

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

func SendPING(ctx context.Context, host host.Host, peerID peer.ID) {

	msgPING := Message{
		ID:   rand.Uint64(),
		Code: uint(0),
		Want: uint(1),
		Data: []byte("PING"),
	}
	sendMessageWithCTX(ctx, host, peerID, msgPING)
}

func SendPONG(ctx context.Context, host host.Host, peerID peer.ID) {

	msgPONG := Message{
		ID:   rand.Uint64(),
		Code: uint(0),
		Want: uint(69),
		Data: []byte("PONG"),
	}
	sendMessageWithCTX(ctx, host, peerID, msgPONG)
}

// Message to send Signed Transaction
func SendSignedTransactionToPeer(signedTxn []byte) {
	for _, peer := range NODE.Network().Peers() {
		txnmsg := Message{
			ID:   rand.Uint64(),
			Code: uint(4),
			Data: signedTxn,
		}
		serializedmsg, err := SerializeMessage(txnmsg)
		if err != nil {
			fmt.Println("Error serializing message:", err)
			return
		}
		stream, err := NODE.NewStream(Ctxt, peer, "/NewTransaction")
		if err != nil {
			fmt.Println("Error opening stream to peer:", err)
			return
		}
		defer stream.Close()

		_, err = stream.Write(serializedmsg)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}
