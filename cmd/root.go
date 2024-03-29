/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "zerotwo",
	Short: "I can help you with your games, darling.",
	Long:
`
                    _                 
  _______ _ __ ___ | |___      _____  
 |_  / _ \ '__/ _ \| __\ \ /\ / / _ \ 
  / /  __/ | | (_) | |_ \ V  V / (_) |
 /___\___|_|  \___/ \__| \_/\_/ \___/ 
                                      
Zerotwo is a CLI created by @iainbrux. It is a CLI that is intended
to be immersive, as if communicating directly to Zerotwo herself.

Well... she is human after all.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zerotwo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


