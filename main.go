package main

import (
	"log"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/api"
)

func main() {
	// Define HTTP handlers
	http.HandleFunc("/addAccount", api.AddAccountHandler)
	http.HandleFunc("/getAccount", api.GetAccountHandler)
	http.HandleFunc("/getNonce", api.GetNonceHandler)
	http.HandleFunc("/addTxn", api.AddTxnHandler)
	http.HandleFunc("/getBalance", api.GetBalanceHandler)
	// gives latest block number
	http.HandleFunc("/getBlockNumber", api.GetBlockNumberHandler)
	// gives block details when block number is given
	http.HandleFunc("/getBlockWithNumber", api.GetBlockWithNumberHandler)
	// gives block details when block hash is given
	http.HandleFunc("/getBlockWithHash", api.GetBlockWithHashHandler)
	http.HandleFunc("/createGenesisBlock", api.CreateGenesisBlock)
	// Start the HTTP server
	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
	// ctx := context.Background()

	// // Predefined multiaddress
	// maddr, err := multiaddr.NewMultiaddr("/ip4/10.1.155.219/tcp/4747")
	// if err != nil {
	// 	panic(err)
	// }
	// node, err := libp2p.New(libp2p.ListenAddrs(maddr))
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Node Addrs: ", node.Addrs())

	// // configure our own ping protocol
	// pingService := &ping.PingService{Host: node}
	// node.SetStreamHandler(ping.ID, pingService.PingHandler)

	// // print the node's PeerInfo in multiaddr format
	// peerInfo := peerstore.AddrInfo{
	// 	ID:    node.ID(),
	// 	Addrs: node.Addrs(),
	// }
	// addrs, err := peerstore.AddrInfoToP2pAddrs(&peerInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("libp2p node address:", addrs[0])

	// // if a remote peer has been passed on the command line, connect to it
	// // and send it 5 ping messages, otherwise wait for a signal to stop
	// if len(os.Args) > 1 {
	// 	addr, err := multiaddr.NewMultiaddr(os.Args[1])
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	peer, err := peerstore.AddrInfoFromP2pAddr(addr)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	if err := node.Connect(context.Background(), *peer); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("sending 5 ping messages to", addr)
	// 	ch := pingService.Ping(context.Background(), peer.ID)
	// 	for i := 0; i < 5; i++ {
	// 		res := <-ch
	// 		fmt.Println("got ping response!", "RTT:", res.RTT)
	// 	}
	// } else {
	// 	// wait for a SIGINT or SIGTERM signal
	// 	ch := make(chan os.Signal, 1)
	// 	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	// 	<-ch
	// 	fmt.Println("Received signal, shutting down...")
	// }

	// // shut the node down
	// if err := node.Close(); err != nil {
	// 	panic(err)
	// }
}
