package main

import (
	"github.com/donvito/hellomod"
	mod2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()

	mod2.SayHello("Mirko")
}
