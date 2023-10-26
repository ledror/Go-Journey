package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type comicInfo struct {
	Num        int `json:"num"`
	URL        string
	Transcript string `json:"transcript"`
}

const (
	comicsFolder = "/tmp/xkcd_jsons"
	xkcdUrl      = "https://xkcd.com"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: xkcd [init|serach]")
	}
	if os.Args[1] == "init" {
		if err := initComics(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("xkcd: database initialized successfully")
	} else {
		if err := searchComics(os.Args[2:]); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done!")
	}
}

func initComics() error {
	resp, err := http.Get(xkcdUrl + "/info.0.json")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("get xkcd home page failed: %s", resp.Status)
	}
	var homeInfo comicInfo
	if err := json.NewDecoder(resp.Body).Decode(&homeInfo); err != nil {
		resp.Body.Close()
		return err
	}
	numComics := homeInfo.Num
	resp.Body.Close()
	if err := os.RemoveAll(comicsFolder); err != nil {
		return err
	}
	if err := os.MkdirAll(comicsFolder, 0777); err != nil {
		return err
	}
	errChan := make(chan error, 0)
	for i := 1; i <= numComics; i++ {
		go getComic(i, errChan)
	}
	for i := 1; i <= numComics; i++ {
		if err := <-errChan; err != nil {
			fmt.Printf("comic failed to download: %s", err)
			return err
		}
	}
	return nil
}

func getComic(idx int, errChan chan error) {
	urlIdx := xkcdUrl + fmt.Sprintf("/%d/info.0.json", idx)
	resp, err := http.Get(urlIdx)
	if err != nil {
		errChan <- err
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		errChan <- err
		return
	}
	var comic comicInfo
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		resp.Body.Close()
		errChan <- err
		return
	}
	resp.Body.Close()
	comic.URL = urlIdx
	data, err := json.Marshal(comic)
	if err != nil {
		errChan <- err
		return
	}
	if err := os.WriteFile(comicsFolder+fmt.Sprintf("/%d.json", idx), data, 0644); err != nil {
		errChan <- err
		return
	}
	errChan <- nil
}

func wordCountMap(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}

func searchComics(s []string) error {
	wordMap := wordCountMap(strings.Join(s, " "))
	files, err := os.ReadDir(comicsFolder)
	if err != nil {
		return err
	}
	errChan := make(chan error, 0)
	for _, file := range files {
		go searchComic(file.Name(), wordMap, errChan)
	}
	for i := 0; i < len(files); i++ {
		if err := <-errChan; err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func searchComic(
	name string,
	wordMap map[string]int,
	errChan chan<- error,
) {
	data, err := os.ReadFile(comicsFolder + "/" + name)
	if err != nil {
		errChan <- err
		return
	}
	var comic comicInfo
	if err := json.Unmarshal(data, &comic); err != nil {
		errChan <- err
		return
	}
	transcriptMap := wordCountMap(comic.Transcript)
	for key, val := range wordMap {
		if transcriptMap[key] < val {
			errChan <- nil
			return
		}
	}
	fmt.Printf("%s: transcript: %s\n", comic.URL, comic.Transcript)
	errChan <- nil
}
