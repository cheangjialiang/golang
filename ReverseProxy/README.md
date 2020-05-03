# Simple Go language https ReverseProxy server Example

This is a simple ReverseProxy server written in Go and using the Echo Framework

Just run the following

## Usage
```
go get golang.org/x/crypto/acme/autocert
go get github.com/labstack/echo/middleware
go get github.com/labstack/echo

go build ReverseProxy.go
sudo setcap CAP_NET_BIND_SERVICE=+eip todo
./ReverseProxy &
```

## Gin Web server
This is a simple web server using the Gin Framework https://gin-gonic.com/
```
go get github.com/gin-gonic/gin
go get github.com/thinkerou/favicon
go run main.go
```
<img src="https://user-images.githubusercontent.com/47512872/54908985-19b1a400-4f24-11e9-9864-88988c31ce23.png" height="200">
