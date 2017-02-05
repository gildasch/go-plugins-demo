package main

import (
	"os"
	"plugin"
)

func main() {
	for _, pName := range os.Args[1:] {
		p, err := plugin.Open(pName)
		if err != nil {
			panic(err)
		}
		sayHello, err := p.Lookup("SayHello")
		if err != nil {
			panic(err)
		}
		sayHello.(func(string))("Gildas")
	}
}
