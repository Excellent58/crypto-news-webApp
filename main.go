package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type NewsResponse struct {
	Data []struct {
		ScreenData struct {
			News     []struct {
				NewsID           int    `json:"news_ID"`
				NewsProviderName string `json:"news_provider_name"`
				Type             string `json:"type"`
				HEADLINE         string `json:"HEADLINE"`
				NewsLink         string `json:"news_link"`
				ThirdPartyURL    string `json:"third_party_url"`
				RelatedImageBig  string `json:"related_image_big"`
			} `json:"news"`
		} `json:"screen_data"`
	} `json:"data"`
}

func main() {
	engine := html.New("./templates", ".html")
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	app.Static("/static", "./static")
	
	app.Get("/", func(c *fiber.Ctx) error {
		results := getCryptoNews()
		data := results.Data[0].ScreenData.News

		return c.Render("index", data)
	})


	app.Listen(":5000")
}


func getCryptoNews() NewsResponse {
	url := "https://investing-cryptocurrency-markets.p.rapidapi.com/coins/get-news?pair_ID=1057391&page=1&time_utc_offset=28800&lang_ID=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "dab64a933emshd4feed0a110bc98p1fa4f0jsncfc1e252bd18")
	req.Header.Add("X-RapidAPI-Host", "investing-cryptocurrency-markets.p.rapidapi.com")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	var newsResponse NewsResponse
	er := json.Unmarshal(jsonBytes, &newsResponse)
	if er != nil {
		log.Fatal(er)
	}

	return newsResponse
}
