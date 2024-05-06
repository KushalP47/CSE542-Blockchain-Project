package p2p

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
)

func StartRootNode(ctx context.Context) {
	root, err := StartNewNode()
	Ctxt = ctx
	NODE = root
	if err != nil {
		panic(err)
	}

	root.SetStreamHandler("/Ping", func(s network.Stream) {
		// fmt.Println("Received stream from:", s.Conn().RemotePeer())
		msg, err := ReceiveMessage(s)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(msg.Data))
		if msg.Want == uint(1) {
			SendPONG(ctx, root, s.Conn().RemotePeer())
		}
	})

	GetUpdatedPeerList := func() {
		fmt.Println("Peers IDs", root.Network().Peers())
		fmt.Println("Peers addrs", root.Network().Peerstore().PeersWithAddrs())
		fmt.Println("conns", root.Network().Conns())
		var addresses []string

		for _, conn := range root.Network().Conns() {
			// Extract peer addresses from the connection
			// localAddr := conn.LocalMultiaddr().String()
			remoteAddr := conn.RemoteMultiaddr().String()

			// Extract Peer ID
			remotePeerID := conn.RemotePeer()

			// Construct the multiaddress with Peer ID
			remoteAddrWithPeerID := fmt.Sprintf("%s/p2p/%s", remoteAddr, remotePeerID)

			// Add the peer addresses to the list
			addresses = append(addresses, remoteAddrWithPeerID)
		}
		fmt.Println("List of peer addresses:")
		addresses = append(addresses, HostPeerAddr)
		for _, addr := range addresses {
			fmt.Println(addr)
		}
		PeerAddrList = addresses

	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			GetUpdatedPeerList()
		}
	}()
	// SendPING(ctx, host, peerAddrInfo.ID)

	// Handle termination signals
	<-ctx.Done()

}

func GetPeerAddrs() []string {
	return PeerAddrList
}
