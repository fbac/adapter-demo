package extlib

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/rpc"
)

type ExtClient struct {
	Id string
	RPC string
}

func (ec *ExtClient) ReadBalance(location string) {
	fmt.Printf("ExtClient %s reading balance from %s\n", ec.Id, location)
}

func (ec *ExtClient) WriteBalance(location string) {
	fmt.Printf("ExtClient %s writing balance from %s\n", ec.Id, location)
}

func NewClient(id, endpoint string) *ExtClient {
	return &ExtClient{
		Id: id,
		RPC: endpoint,
	}
}

type Reader interface {
	ReadBalance(location string)
}

type Writer interface {
	WriteBalance(location string)
}

type ReaderWriter interface {
	Reader
	Writer
}

func ReaderWriterReadWrite(r ReaderWriter) {
	r.ReadBalance("ReaderWriter interface")
	r.WriteBalance("ReaderWriter interface")
}

func (ec *ExtClient) Discover(endpoint string) error {
	client, err := rpc.Dial(endpoint)
	if err != nil {
		return err
	}

	var chainId string
	if err := client.Call(&chainId, "eth_chainId"); err != nil {
		return err
	}

	return nil
}

func (ec *ExtClient) Run() {
	fmt.Printf("Created Ethereum client with id: %v\n", ec.Id)
	client, err := rpc.Dial(ec.RPC)
	if err != nil {
		os.Exit(1)
	}

	var chainId string
	client.Call(&chainId, "eth_chainId")
	fmt.Printf("Client %v got chainId %v\n", ec.Id, chainId)
}
