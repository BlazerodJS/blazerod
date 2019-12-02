package main

import (
	"fmt"
	"os"

	"github.com/BlazerodJS/blazerod/pkg/v8engine"
	"github.com/BlazerodJS/blazerod/pkg/version"
	"github.com/spf13/cobra"
)

var script = `
V8Engine.log("hello from JS");
`

var rootCmd = &cobra.Command{
	Use:               "blaze",
	Short:             "Blazerod JavaScript Runtime",
	Version:           version.Version(),
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		engine := v8engine.NewEngine()
		val, err := engine.Run(script, "main.js")
		if err != nil {
			panic(err)
		}

		fmt.Println(val)
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
