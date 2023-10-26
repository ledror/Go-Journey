package main

import (
	"Chapter-04/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	monthOldItems := make([]*github.Issue, 0)
	yearOldItems := make([]*github.Issue, 0)
	moreThanYearOldItems := make([]*github.Issue, 0) // lmao
	now := time.Now()
	for _, item := range result.Items {
		switch hours := item.CreatedAt.Sub(now).Hours(); {
		case hours < 24*30:
			monthOldItems = append(monthOldItems, item)
		case hours < 24*365:
			yearOldItems = append(yearOldItems, item)
		default:
			moreThanYearOldItems = append(moreThanYearOldItems, item)
		}
	}
	fmt.Printf("%d one month old issues\n", len(monthOldItems))
	for _, item := range monthOldItems {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("%d one year old issues\n", len(yearOldItems))
	for _, item := range yearOldItems {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("%d more than one year old issues\n", len(moreThanYearOldItems))
	for _, item := range moreThanYearOldItems {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
