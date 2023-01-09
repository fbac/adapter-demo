package main

import (
	"fmt"
	"github.com/fbac/adapter-demo/pkg/adapter"
	"github.com/fbac/adapter-demo/pkg/extlib"
	"github.com/fbac/adapter-demo/pkg/extlib2"
)

func main() {
	// Create client
	ec := extlib.NewClient("ec1")
	fmt.Printf("created Client %v\n", ec.Id)
	ec2 := extlib2.NewClient("ec2")
	fmt.Printf("created Client %v\n", ec2.Id)

	// Call methods directly
	ec.ReadBalance("ExtClient main")
	ec.WriteBalance("ExtClient main")

	// Call methods through ReadWriter interface
	extlib.ReaderWriterReadWrite(ec)

	// Call methods through CompatibilityAdapter interface
	var compat adapter.Adapter
	compat.Run(ec)
	compat.Run(ec2)
}
