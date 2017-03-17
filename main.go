package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	init := flag.Bool("init", false, "Initialize database")
	dbPath := flag.String("db", "./recruiters.db", "Path to DB")
	port := flag.Int("port", 8080, "Port to listen on")
	staticPath := flag.String("static", "./static", "Static asset path")
	flag.Parse()

	handler, err := NewHandler(*dbPath, *init)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(*staticPath))))
	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/adduser", handler.AddUser)
	mux.HandleFunc("/getuser", handler.GetUser)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
