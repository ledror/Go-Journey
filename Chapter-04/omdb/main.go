package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	APIKEY_VAR  = "OMDB_APIKEY"
	POSTER_PATH = "/tmp/posters"
)

type MovieJsonResult struct {
	Title     string `json:"Title"`
	PosterURL string `json:"Poster"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: omdb 'movie name'\n")
	}
	base, err := url.Parse("http://www.omdbapi.com")
	if err != nil {
		log.Fatal(err)
	}

	apikey, ok := os.LookupEnv(APIKEY_VAR)
	if !ok {
		log.Fatal("enviroment variable OMDB_APIKEY unset")
	}

	params := url.Values{}
	params.Add("t", strings.Join(os.Args[1:], " "))
	params.Add("apikey", apikey)
	base.RawQuery = params.Encode()
	resp, err := http.Get(base.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("query failed: %s\n", resp.Status)
	}
	var result MovieJsonResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("decoding failed: %s\n", err.Error())
	}
	if result.PosterURL == "N/A" {
		log.Fatalf("poster not available\n")
	}
	postRes, err := http.Get(result.PosterURL)
	if err != nil {
		log.Fatalf("invalid poser url: %s\n", err.Error())
	}
	defer postRes.Body.Close()
	if err := os.MkdirAll(POSTER_PATH, 0777); err != nil {
		log.Fatalf("error creating poster directory: %s\n", err.Error())
	}
	filePath := fmt.Sprintf("%s/%s%s", POSTER_PATH, result.Title, filepath.Ext(result.PosterURL))
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("error creating poster file: %s\n", err.Error())
	}
	defer file.Close()
	if _, err := file.ReadFrom(postRes.Body); err != nil {
		log.Fatalf("error reading from poster body: %s\n", err.Error())
	}
	fmt.Printf("downloaded poster to %s\n", filePath)
}
