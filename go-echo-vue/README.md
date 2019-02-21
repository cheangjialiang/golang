# Simple Go language https webserver

This is a simple todo web app written in Go and using the Echo Framework

Just run the folowing

## Usage
```
go get golang.org/x/crypto/acme/autocert
go get github.com/labstack/echo/middleware
go get github.com/labstack/echo
go get github.com/mattn/go-sqlite3

go build todo.go
sudo setcap CAP_NET_BIND_SERVICE=+eip todo
./todo

```
<img src="https://user-images.githubusercontent.com/47512872/53147435-87e00f80-35e2-11e9-9e05-0c961fd4849e.png" height="200">

Then point your browser to https://localhost

<img src="https://user-images.githubusercontent.com/47512872/53147695-4ac84d00-35e3-11e9-8e26-dc5342359598.png" height="200">
