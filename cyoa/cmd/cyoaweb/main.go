package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	cyoa "github.com/akanumuri/learngo/cyoa/pkg"
)

func main() {
	jFile := flag.String("jFile", "gopher.json", "Path to the location on json file")
	port := flag.Int("port", 3000, " Port on which the web server is runnign")
	flag.Parse()

	f, err := os.Open(*jFile)
	if err != nil {
		log.Fatal(err)
	}
	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}

	// h := cyoa.NewHandler(story)
	fmt.Printf("Starting HTTP server on %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), story))
}

