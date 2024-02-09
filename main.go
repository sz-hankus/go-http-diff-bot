package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func listenForSignals() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	received := <-sig
	handleSignal(received)
}

func handleSignal(sig os.Signal) {

	switch sig {
	case syscall.SIGTERM:
		fmt.Println("Received a SIGTERM, shutting down")
		os.Exit(1)
	case syscall.SIGINT:
		fmt.Println("Received a SIGINT, shutting down")
		os.Exit(1)
	}
}

func main() {
	go listenForSignals()

	url1 := "https://www.rebel.pl/promocje/nemesis-edycja-polska-105353.html"

	attr, err := queryHtmlFromUrl(url1, "meta[property='product:price:amount']", "content")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("url: %s, attr: %s\n", url1, attr)
}
