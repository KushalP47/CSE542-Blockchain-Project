package p2p

import (
	"context"
	"fmt"
	"time"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/ethereum/go-ethereum/rlp"
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

	root.SetStreamHandler("/NewTransaction", func(s network.Stream) {
		// fmt.Println("Received stream from:", s.Conn().RemotePeer())
		msg, err := ReceiveMessage(s)
		if err != nil {
			panic(err)
		}

		fmt.Println("Received Transaction from", s.Conn().RemotePeer())

		// Deserialize the message
		var txn blockchain.SignedTx
		err = rlp.DecodeBytes(msg.Data, &txn)
		if err != nil {
			fmt.Println("Error deserializing message:", err)
			return
		}

		// Verify the transaction
		if !blockchain.VerifySignedTxn(txn) {
			fmt.Println("Transaction verification failed")
			return
		}

		// add Txn to the transactionsData
		txnNumber, err := database.GetLastTxnKey()
		if err != nil {
			fmt.Println("Error getting last txn key:", err)
			return
		}
		txnNumber++
		err = blockchain.AddTxn(txnNumber, txn)
		if err != nil {
			fmt.Println("Error adding transaction:", err)
			return
		}

		fmt.Println("Transaction added to the blockchain")
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
