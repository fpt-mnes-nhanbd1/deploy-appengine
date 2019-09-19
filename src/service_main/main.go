package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func createMux() *echo.Echo {
	e := echo.New()
	return e
}

func main() {
	//この辺は相変わらず必要
	e := echo.New()
	http.Handle("/", e)

	//ここから追加
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	//この部分で起動
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
