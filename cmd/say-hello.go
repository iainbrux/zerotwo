package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sayHelloCmd)
}

var sayHelloCmd = &cobra.Command{
	Use: "hello",
	Short: "Says hello, because you know... you're lonely.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello darling.")
	},
}