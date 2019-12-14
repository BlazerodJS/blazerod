package main

import (
	"fmt"
	"os"

	"github.com/BlazerodJS/blazerod/pkg/resolver"
	"github.com/BlazerodJS/blazerod/pkg/v8engine"
	"github.com/BlazerodJS/blazerod/pkg/version"
	"github.com/spf13/cobra"
)

var script = `
function uintToString(data) {
	return String.fromCharCode.apply(null, new Uint8Array(data))
}

V8Engine.cb((msg) => {
	V8Engine.log("Got a message from Go!");
	V8Engine.log(msg);
	V8Engine.log(uintToString(msg));
});

V8Engine.log("hello from JS");
`

func resolveModule(moduleName, referrerName string) (string, int) {
	fmt.Println("requested", moduleName, referrerName)
	return "", 0
}

var rootCmd = &cobra.Command{
	Use:               "blaze",
	Short:             "Blazerod JavaScript Runtime",
	Version:           version.Version(),
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		engine := v8engine.NewEngine()
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		r := resolver.NewResolver(engine, cwd)

		_, _ = engine.Run("var global = {};", "global")

		ret := engine.LoadModule(script, "main.js", r.ResolveModule)
		fmt.Println(ret)

		err = engine.Send([]byte("hello from Go"))
		if err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("{{.Version}}")
}

func main() {
	Execute()
}
