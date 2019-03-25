# Simple Go language https ReverseProxy server

This is a simple ReverseProxy server written in Go and using the Echo Framework

Just run the folowing

## Usage
```
go get golang.org/x/crypto/acme/autocert
go get github.com/labstack/echo/middleware
go get github.com/labstack/echo

go build ReverseProxy.go
sudo setcap CAP_NET_BIND_SERVICE=+eip todo
./ReverseProxy &

go run main.go

```
