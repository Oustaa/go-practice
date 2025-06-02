package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/oustaa/go-practice/internal/app"
	"github.com/oustaa/go-practice/internal/routes"
)

func main() {
	var PORT int
	flag.IntVar(&PORT, "port", 9000, "This is the port, the app will be running on!!")
	flag.Parse()

	app := app.NewApplication()
	defer app.DB.Close()

	handler := routes.GetRoutes(app)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: handler,
	}

	log.Printf("[Info] Application running on port %d", PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error (server.ListenAndServe()): %v", err)
	}
}
