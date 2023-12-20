package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	url2 "net/url"
)

type SomeStruct struct {
	Name string `json:"Person,omitempty"`
	URL  string `json:"URL,omitempty"`
}

type Page struct {
	Pageid    int    `json:"pageid"`
	Ns        int    `json:"ns"`
	Title     string `json:"title"`
	Thumbnail struct {
		Source string `json:"source"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"thumbnail"`
	Pageimage string `json:"pageimage"`
}

type Query struct {
	Normalized []struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"normalized"`
	Pages map[string]Page `json:"pages"`
}

type Response struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         Query  `json:"query"`
}

func main() {
	app := gin.Default()
	app.GET("/:name", func(c *gin.Context) {
		data := SomeStruct{
			Name: "Александр III",
			URL:  "https://ru.wikipedia.org/wiki/%D0%90%D0%BB%D0%B5%D0%BA%D1%81%D0%B0%D0%BD%D0%B4%D1%80_III",
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, HashSHA256")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.JSON(http.StatusOK, data)
	})
	app.GET("/imageWiki/:name", func(c *gin.Context) {
		url := fmt.Sprintf("https://ru.wikipedia.org/w/api.php?action=query&titles=%s&prop=pageimages&format=json&pithumbsize=100", url2.QueryEscape(c.Param("name")))
		client := resty.New()
		resp, err := client.R().Get(url)
		if err != nil {
			panic(err)
		}
		var response Response
		err = json.Unmarshal(resp.Body(), &response)
		if err != nil {
			panic(err)
		}

		page := response.Query.Pages["3101"] // Access the first page
		imageURL := page.Thumbnail.Source
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, HashSHA256")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		fmt.Println(imageURL)
		c.JSON(http.StatusOK, SomeStruct{URL: imageURL})
	})
	app.Run(":8080")
}
