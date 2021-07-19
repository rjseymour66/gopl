// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount	int `json:"total_count"`
	Items		[]*Issue
}

type Issue struct {
	Number		int
	HTMLURL		string `json:"html_url"`
	Title		string
	State		string
	User		*User
	CreatedAt	time.Time `json:created_at"`
	Body		string
}

type User struct {
	Login		string
	HTMLURL		string `json:"html_url"`
}

func main() {

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now   := time.Now()
	month := now.AddDate(0, -1, 0)
	year  := now.AddDate(-1, 0, 0)

	fmt.Printf("%d issues:\n", result.TotalCount)

	fmt.Println("In the past month:")
	for _, item := range result.Items {
		if item.CreatedAt.After(month) {
			printIssue(item)
		}
	}

	fmt.Println("In the past year:")
	for _, item := range result.Items {
		if item.CreatedAt.Before(month) && item.CreatedAt.After(year) {
			printIssue(item)
		}
	}

	fmt.Println("More than a year ago:")
	for _, item := range result.Items {
		if item.CreatedAt.Before(year) {
			printIssue(item)
		}
	}
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	// Create query string using the terms slice and check err
	qString := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + qString)
	if err != nil {
		return nil, err
	}

	// Check if the response was 200
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	// decode json from input stream, use Decode to store in result addr
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func printIssue(issue *Issue) {
	fmt.Printf("#%-5d %9.9s %.55s %v\n", issue.Number, issue.User.Login, issue.Title)
}
