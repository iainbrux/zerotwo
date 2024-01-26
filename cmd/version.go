package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Prints my current version number... even though I'm human.",
	Long: "All software has versions... but I'm human, aren't I darling?",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
`        ___   ___  
       / _ \ |__ \ 
__   _| | | |   ) |
\ \ / / | | |  / / 
 \ V /| |_| | / /_ 
  \_/  \___(_)____|
		`)
	},
}