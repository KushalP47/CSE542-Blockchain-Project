package p2p

import (
	"context"
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

func SendSignedTransaction(signedTxn []byte) {
	for _, peer := range NODE.Network().Peers() {
		msg := Message{
			ID:   rand.Uint64(),
			Code: uint(4),
			Data: signedTxn,
		}
		sendMessageWithCTX(Ctxt, NODE, peer, msg)
	}
}
