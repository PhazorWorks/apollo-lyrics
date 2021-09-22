package main

import (
"github.com/labstack/echo/v4"
"github.com/labstack/echo/v4/middleware"
"io/ioutil"
"log"
"net/http"
"net/url"
"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/lyrics", func(c echo.Context) error {
		query := c.QueryParam("q") // get query from url parameters
		resp, err := http.Get("https://evan.lol/lyrics/search/top?q="+url.QueryEscape(query)) // send off the new query
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		return c.HTML(http.StatusOK, string(body))
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
