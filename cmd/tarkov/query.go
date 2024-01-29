package tarkov

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	queryTarkovCmd.Flags().StringP("item", "i", "", "")
	queryTarkovCmd.MarkFlagRequired("item")

	tarkovCmd.AddCommand(queryTarkovCmd)
}

type QueryResponse struct {
	Data struct {
		Items []QueryItem `json:"items"`
	} `json:"data"`
}

type QueryItem struct {
	Avg24hPrice int `json:"avg24hPrice"`
	Id string `json:"id"`
	Name string `json:"name"`
	Shortname string `json:"shortName"`
}

var queryTarkovCmd = &cobra.Command{
	Use: "query",
	Short: "Let me help you with Escape from Tarkov",
	Long:`Zerotwo will search through her deep knowledge and obtain a list of 
items that contain your query search term.

She then presents the item(s) in the form of "[ShortName] Full Item Name" which
can then be used to ask Zerotwo to further provide indepth data of the item by
passing the shortname to the "zerotwo tarkov item" command.`,
	Run: func(cmd *cobra.Command, args []string) {
		item, _ := cmd.Flags().GetString("item")

		cache := &tarkovCache{}
		res, ok := cache.read(item)

		if ok {
			if time.Now().Before(res.Timestamp.Add(time.Minute * 5)) {
				formatResponse(res.Items, true)
				return
			}
		}

		body := strings.NewReader(`{"query": "{ items(name: \"`+ item +`\") {avg24hPrice id name shortName } }"}`)
		req, err := http.NewRequest("POST", "https://api.tarkov.dev/graphql", body)
		if err != nil {
				log.Fatalln(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
				log.Fatalln(err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
				log.Fatalln(err)
		}

		defer resp.Body.Close()

		data := QueryResponse{}
		json.Unmarshal(bodyBytes, &data)

		if len(data.Data.Items) == 0 {
			fmt.Printf("Sorry darling, no results found for item '%s'\n", item)
			return
		}

		cache.update(item, data.Data.Items)
		formatResponse(data.Data.Items, false)
	},
}

func formatResponse(items []QueryItem, cached bool) {
	header := "D: [ShortName] Full Item Name // ₽ 24hr Avg Price"

	if cached {
		header = header + " -- CACHED"
	}

	fmt.Println(header)
	for index, item := range items {
		fmt.Printf("%d: " + "[" + item.Shortname + "] " + item.Name + " // ₽%d \n", index, item.Avg24hPrice)
	}
}