package msg

import (
	"leaf/network/json"
)

var Processor =json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&People{})
}

type Hello struct {
	Name string
}

type People struct {
	Name string
	Sex string
	Age int
}
