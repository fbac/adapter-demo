package adapter

import (
	//"fmt"
	"github.com/fbac/adapter-demo/pkg/extlib"
)

type CompatibilityAdapter interface {
	Reader() interface{}
	Writer() interface{}
}

type Adapter struct{}

func (a Adapter) Run(rw any) {
	//fmt.Printf("%T\n", a)
	//fmt.Printf("%T\n", rw)
	rw.(*extlib.ExtClient).ReadBalance("adapter")
	rw.(*extlib.ExtClient).WriteBalance("adapter")
}
