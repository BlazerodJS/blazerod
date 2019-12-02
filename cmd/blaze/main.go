package main

import (
	"fmt"

	"github.com/BlazerodJS/blazerod/pkg/v8engine"
)

var script = `
V8Engine.log("hello from JS");
`

func main() {
	fmt.Println(v8engine.Version())

	engine := v8engine.NewEngine()
	val, err := engine.Run(script, "main.js")
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
