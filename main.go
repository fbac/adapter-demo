package main

import (
	"fmt"
	"log"

	"github.com/fbac/adapter-demo/pkg/adapter"
	"github.com/fbac/adapter-demo/pkg/extlib"
	"github.com/fbac/adapter-demo/pkg/extlib2"
)

func main() {
	// Create client
	var client adapter.CompatibilityAdapter
	var endpoint string = "https://goerli.infura.io/v3/"
	//var endpoint string = "https://api.testnet.solana.com"

	if err := extlib.NewClient("ec1", endpoint).Discover(endpoint); err == nil {
		fmt.Println("extclient")
		client = extlib.NewClient("ec1", endpoint) 
	} else if err := extlib2.NewClient("ec2").Discover(endpoint); err == nil {
		fmt.Println("extclient2")
		client = extlib2.NewClient("ec2") 
	} else {
		log.Fatalf("%s is not a blockchain endpoint", endpoint)
	}

	client.Run()
}
