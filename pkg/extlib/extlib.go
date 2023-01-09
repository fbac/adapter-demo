package extlib

import "fmt"

type ExtClient struct {
	Id string
}

func (ec *ExtClient) ReadBalance(location string) {
	fmt.Printf("ExtClient %s reading balance from %s\n", ec.Id, location)
}

func (ec *ExtClient) WriteBalance(location string) {
	fmt.Printf("ExtClient %s writing balance from %s\n", ec.Id, location)
}

func NewClient(id string) *ExtClient {
	return &ExtClient{
		Id: id,
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

func (ec *ExtClient) Run() {
	ec.ReadBalance("run")
	ec.WriteBalance("run")
}
