package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"math/rand"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var result = ""

func main() {
	e := echo.New()

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 3,
	}))
	/*
		Logging
	*/
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip=${remote_ip}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", VPNStatus)
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Fatal(e.Start(":80")) //Change Port
}

func customHTTPErrorHandler(err error, c echo.Context) {
	c.NoContent(404)
}

func VPNStatus(c echo.Context) error {
	if rand.Int()%5 == 0 || result == "" {
		go func() {
			out, err := exec.Command("python", "status.py").Output()
			if err == nil {
				result = `<center>VPNGate Status<br>` + strings.Replace(string(out), "\n", "<br>", -1) + `<br><a href="https://github.com/x86taka/VPNGate-Web-Status">Source Code</a></center>`
			} else {
				log.Println(err)
			}
		}()
	}
	return c.HTML(http.StatusOK, result)
}
