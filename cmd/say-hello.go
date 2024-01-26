package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(sayHelloCmd)
}

var sayHelloCmd = &cobra.Command{
	Use: "hello",
	Short: "I'll always say hello for you. Whenever you like.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello darling.")
	},
}