package extlib2

import "fmt"

type ExtClient2 struct {
	Id string
}

func (ec *ExtClient2) ReadBalance(location string) {
	fmt.Printf("ExtClient2 %s reading balance from %s\n", ec.Id, location)
}

func (ec *ExtClient2) WriteBalance(location string) {
	fmt.Printf("ExtClient2 %s writing balance from %s\n", ec.Id, location)
}

func (ec *ExtClient2) SpecificMethod() {
	fmt.Printf("ExtClient2 %s called SpecificMethod\n", ec.Id)
}

func NewClient(id string) *ExtClient2 {
	return &ExtClient2{
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

func (ec *ExtClient2) Discover(endpoint string) error {
	return fmt.Errorf("not a real client")
}

func (ec *ExtClient2) Run() {
	fmt.Printf("- Run() started for client %v\n", ec.Id)
	ec.ReadBalance("run")
	ec.WriteBalance("run")
	ec.SpecificMethod()
}
