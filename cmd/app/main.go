package main

import (
	"flag"
	"log"
	"os"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "HTTP server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development||staging||production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	logger.Printf("%s server started on port %d", cfg.env, cfg.port)

	err := app.serve()
	if err != nil {
		logger.Fatal(err, nil)
	}

}
