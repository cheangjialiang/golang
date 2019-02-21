package main

import (
	"database/sql"

	"./handlers"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/engine/standard"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/acme/autocert"
	"github.com/labstack/echo/middleware"
)

func main() {

	db := initDB("storage.db")
	migrate(db)

	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("saikang.tk")
	e.AutoTLSManager.Cache = autocert.DirCache(".cache")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))
	
	go func(c *echo.Echo){
		e.Logger.Fatal(e.Start(":80"))
	}(e)

	e.Logger.Fatal(e.StartAutoTLS(":443"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
	);
	`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
