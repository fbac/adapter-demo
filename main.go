package main

import (
	"github.com/fbac/adapter-demo/pkg/adapter"
	"github.com/fbac/adapter-demo/pkg/extlib"
)

func main() {
	ec := extlib.NewClient("main")
	ec.ReadBalance("ExtClient - main")
	ec.WriteBalance("ExtClient - main")
	extlib.ReaderWriterReadWrite(ec)

	var compat adapter.Adapter
	compat.Run(ec)
}
