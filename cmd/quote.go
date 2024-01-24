package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(quoteCmd)
}

var quoteCmd = &cobra.Command{
	Use: "quote",
	Short: "I'll recall a quote from my fractured memories.",
	Long: "Will say a random quote from ZeroTwo, it may put you in touch with your emotions.",
	Run: func(cmd *cobra.Command, args []string) {
		quotes := make([]string, 0)

		quotes = append(quotes, 
			"You are now my darling!", 
			"Your taste makes my heart race. It bites and lingers...the taste of danger.", 
			"Its been a long time since I last saw a human cry.",
			"The distant skies. Beyond time and distance. An overwhelmingly long journey just for the two of us. You're a part of me. I'm a part of you. I'll remember your warmth, along with the memories we've made together. I'll never let you go again!",
			"I'm always alone, too. Thanks to these horns.",
		)

		randomIndex := rand.Intn(len(quotes))

		fmt.Println(quotes[randomIndex])
	},
}