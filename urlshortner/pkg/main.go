package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akanumuri/learngo/urlshortner"
)

func main() {
	mux := http.NewServeMux()
	filePath := flag.String("filepath", "Path2URL.yaml", "Yaml file input")
	flag.Parse()

	UrlMap := map[string]string{
		"/apple":  "https://apple.com",
		"/google": "https:/google.com",
	}
	maphandler := urlshortner.MapsHandler(UrlMap, mux)
	data, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatal("unable to read file", err)
	}
	yamlHandler, err := urlshortner.YamlHandler(data, maphandler)
	if err != nil {
		fmt.Println("Failed do decome yaml")
	}

	http.ListenAndServe(":8080", yamlHandler)
}
