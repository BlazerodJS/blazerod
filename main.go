package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BlazerodJS/blazerod/pkg/resolver"
	"github.com/BlazerodJS/blazerod/pkg/v8engine"
	"github.com/BlazerodJS/blazerod/pkg/version"
)

func main() {
	first := os.Args[1]

	switch first {
	case "run":
		if len(os.Args) != 3 {
			fmt.Println("Usage: blaze run script.js")
			os.Exit(1)
		}

		script, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}

		engine := v8engine.NewEngine()
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		r := resolver.NewResolver(engine, cwd)

		_, _ = engine.Run("var global = {};", "global")

		ret := engine.LoadModule(string(script), "main.js", r.ResolveModule)
		fmt.Println(ret)

		err = engine.Send([]byte("hello from Go"))
		if err != nil {
			panic(err)
		}

	case "version":
		fmt.Println(version.Version())
	}
}
