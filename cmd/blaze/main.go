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
import {test} from 'test';

V8Engine.log("hello from JS");

V8Engine.log(test);
`

var testModule = `
export const test = 123;
export const abc = 'def';
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
		ret := engine.LoadModule(testModule, "test", resolver.ResolveModule)
		fmt.Println(ret)
		ret = engine.LoadModule(script, "main.js", resolver.ResolveModule)
		fmt.Println(ret)
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
