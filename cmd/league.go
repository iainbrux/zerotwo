package cmd

import (
	"log"
	"net/url"
	"regexp"
	"strings"
	"zerotwo/scraper"

	"github.com/spf13/cobra"
)

func init() {
	getCmd.Flags().StringP("champion", "c", "", "Name of the LoL champion you wish ZeroTwo to search for")
	getCmd.MarkFlagRequired("champion")

	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "league",
	Short: "Need help with a League of Legends champion darling?",
	Long:`Harnesses the power of ZeroTwo and obtains a quick synopsis of a LoL champion that you provide.
Defaults to guides for Emerald+ and the champion's highest lane pick percentage.
Support may be added in the future to specify rank and lane.`,
	Run: func(cmd *cobra.Command, args []string) {
		champion, _ := cmd.Flags().GetString("champion")

		url := "https://www.op.gg/champions/" + formatChampionString(champion) + "/build"
		validUrl := validateUrl(url)

		scraper.Scrape(validUrl.String(), "league")
	},
}

func validateUrl(u string) *url.URL {
	validUrl, err := url.ParseRequestURI(u)

	if err != nil {
		log.Fatalln("Error:", err)
	}

	return validUrl
}

func formatChampionString(champion string) string {
	lowercaseChampion := strings.ToLower(champion)

	// Replace any apostrophes or whitespace in champion names with an
	// empty string for the url
	re := regexp.MustCompile(`['\s]`)

	return re.ReplaceAllString(lowercaseChampion, "")
}