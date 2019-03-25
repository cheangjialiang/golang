package main

import (
	"net/url"
	"os"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache(".cache")
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("saikang.tk")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	// Setup proxy
	url1, err := url.Parse("http://localhost:5061")
	if err != nil {
		e.Logger.Fatal(err)
	}
	/*url2, err := url.Parse("http://localhost:80")
	if err != nil {
		e.Logger.Fatal(err)
	}*/
	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
		/*{
			URL: url2,
		},*/
	}
	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))
	/*e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} | ${status} | ${latency_human}\t | ${method} | ${remote_ip} | ${uri} \n",
	}))*/
	out, err := os.Create("ReverseProxy.log")
	if err != nil {
		panic(err)
	}
	e.Logger.SetOutput(out)
	e.Logger.Fatal(e.StartAutoTLS(":443"))
	
}