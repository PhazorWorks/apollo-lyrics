package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Lyrics struct {
	Lyrics   string
	Explicit bool
	Artist   string
	Name     string
}

type Response struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Fuzzy bool   `json:"fuzzy"`
	Album struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Icon struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"icon"`
	} `json:"album"`
	Length   int  `json:"length"`
	Explicit bool `json:"explicit"`
	Artists  []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"artists"`
	Lyrics string `json:"lyrics"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/lyrics", func(c echo.Context) error {
		query := c.QueryParam("q")                                                              // get query from url parameters
		resp, err := http.Get("https://evan.lol/lyrics/search/top?q=" + url.QueryEscape(query)) // send off the new query
		if err != nil {
			log.Fatalln(err)
		}
		if err != nil {
			log.Fatalln(err)
		}
		var response Response
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			log.Println(err)
		}
		lyrics := Lyrics{
			Name:     response.Name,
			Lyrics:   response.Lyrics,
			Artist:   response.Artists[0].Name,
			Explicit: response.Explicit,
		}
		jsonData, err := json.Marshal(lyrics)
		if err != nil {
			log.Println(err)
		}
		return c.HTML(http.StatusOK, string(jsonData))
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
