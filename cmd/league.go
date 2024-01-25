package cmd

import (
	"log"
	"net/url"
	"zerotwo/scraper"

	"github.com/spf13/cobra"
)

func init() {
	getCmd.Flags().StringP("champion", "c", "", "Help message for champion")
	getCmd.MarkFlagRequired("champion")

	rootCmd.AddCommand(getCmd)
}

func validateUrl(u string) *url.URL {
	validUrl, err := url.ParseRequestURI(u)

	if err != nil {
		log.Fatalln("Error:", err)
	}

	return validUrl
}

var getCmd = &cobra.Command{
	Use: "league",
	Short: "Need help with winning your League of Legends games darling?",
	Long:
	`
Harnesses the power of ZeroTwo and obtains a quick synopsis of a LoL champion that you provide.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		champion, _ := cmd.Flags().GetString("champion")

		url := "https://www.op.gg/champions/" + champion + "/build"
		validUrl := validateUrl(url)

		scraper.Scrape(validUrl.String())
	},
}