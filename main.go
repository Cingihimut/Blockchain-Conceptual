package main

import (
	"blockchain-conceptual/blockchain"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	bc := blockchain.NewBlockchain()

	peerAddr := "localhost:3000"
	peer := network.Newpeer(peerAddr, bc)

	go func() {
		err := http.ListenAndServe(peerAddr, peer)
		if err != nil {
			log.Fatal()
		}
	}()

	waitExitSignal()

	fmt.Println("Blockchain ndoe is running on", peerAddr)

	select {}

}

func waitExitSignal() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		fmt.Println("\nShutting down...")
		os.Exit(0)
	}()
	wg.Wait()
}
