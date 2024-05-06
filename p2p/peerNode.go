package p2p

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func ConnectToPeers(host host.Host) {
	for _, addr := range PeerAddrList {
		if addr != "" {

			fmt.Println(addr)
			peerMA, err := multiaddr.NewMultiaddr(addr)
			if err != nil {
				panic(err)
			}
			peerAddrInfo, err := peer.AddrInfoFromP2pAddr(peerMA)
			if err != nil {
				panic(err)
			}
			if err := host.Connect(context.Background(), *peerAddrInfo); err != nil {
				panic(err)
			}
			fmt.Println("Connected to", peerAddrInfo.String())
			SendPING(context.Background(), host, peerAddrInfo.ID)
		}
	}

}

func GetMultiAddr() { // []multiaddr.Multiaddr,[]peer.AddrInfo
	KnownHosturl := "http://10.1.155.219:8000/getKnownHosts"
	resp, err := http.Get(KnownHosturl)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %v", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Convert the byte slice to a string
	addrString := string(body)

	// Split the string into individual addresses based on newline delimiter
	addrList := strings.Split(addrString, "\n")

	// Print the list of addresses
	fmt.Println("List of addresses:")
	for _, addr := range addrList {
		fmt.Println(addr)
	}
	PeerAddrList = addrList

}

func StartPeerNode(ctx context.Context, peerAddr string) {
	host, err := StartNewNode()
	Ctxt = ctx
	NODE = host
	if err != nil {
		panic(err)
	}

	GetMultiAddr()
	ConnectToPeers(host)

	host.SetStreamHandler("/Ping", func(s network.Stream) {
		// fmt.Println("Received stream from:", s.Conn().RemotePeer())
		msg, err := ReceiveMessage(s)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(msg.Data))
		if msg.Want == uint(1) {
			SendPONG(ctx, host, s.Conn().RemotePeer())
		}

	})

	GetUpdatedPeerList := func() {
		fmt.Println("Peers IDs", host.Network().Peers())
		fmt.Println("Peers addrs", host.Network().Peerstore().PeersWithAddrs())
		fmt.Println("conns", host.Network().Conns())
		var addresses []string

		for _, conn := range host.Network().Conns() {
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
