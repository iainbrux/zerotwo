package tarkov

import (
	"zerotwo/cmd"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(tarkovCmd)
}

var tarkovCmd = &cobra.Command{
	Use: "tarkov",
	Aliases: []string{"eft"},
	Short: "Let me help you with Escape from Tarkov",
	Long:`Zerotwo can find you almost anything to do with Tarkov. She is still 
discovering how to find most things, but for now she can help you with querying 
item lists and providing indepth data of specific items.`,
}