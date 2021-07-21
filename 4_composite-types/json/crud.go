package main

import (
	"bytes"
//	"flag"
	"fmt"
	"encoding/json"
	"log"
	"net/http"
//	"net/url"
	"os"
//	"strings"
	"time"
)

// CREATE:	"https://api.github.com/repos/{owner}/{repo}/issues
// READ:	"https://api.github.com/repos/{owner}/{repo}/issues/{issue_number}
// PATCH:	"https://api.github.com/repos/{owner}/{repo}/issues/{issue_number}
// DELETE:	"https://api.github.com/repos/{owner}/{repo}/issues/{issue_number}

type CreateResponse struct {
	Url		string
	Title		string
	Body		string
	State		string
	CreatedAt	time.Time `json:"created_at"`
}

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
	CreatedAt	time.Time `json:"created_at"`
	Body		string
}

type User struct {
	Login		string
	HTMLURL		string `json:"html_url"`
}

// enter flags on the cmdline to use issues api. If there 
func main() {

	// enter owner and repo from cmdline

	// test vars
	title := "Issue from golang"
	body  := "This is a test issue body comment"

	result, err := CreateIssue(title, body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Issue created:\n%v", result)
}

func CreateIssue(owner, repo, message string) (*CreateResponse, error) {
	buf := &bytes.Buffer{}
	// encode the fields map in json. This stores the encoded json in buf
	// or returns an error
	err := json.NewEncoder(buf).Encode(message)
	if err != nil {
		return nil, err
	}

	// create a client 
	client = &http.Client{}

	// create the url
	url := strings.Join([]string{
		"https://api.github.com", "repos", owner, repo, "issues"} "/")

	// create the request and set authorization using env vars
	req, err := http.NewRequest(http.MethodPost, url, buf)
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASS"))
	if err != nil {
		return nil, err
	}


}

func ReadIssue() {}

func UpdateIssue() {}

func CloseIssue() {}









func xCreateIssue(title string, body string) (*CreateResponse, error) {
	// create the URL
	url := "https://api.github.com/repos/rjseymour66/gopl/issues"

	// create the issue to fill in with the args
	issue := Issue {
		Title:	title,
		Body:	body,
	}

	// marshal it into json
	payload, err := json.Marshal(issue)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error during marshal %s", err)
	}

	// post method needs an io.Reader, which is a buffer
	// NewBuffer initializes the buf with the contents passed as an arg, so it 
	// is perfect for reading existing data

	// send marshaled payload in post req
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("post failed: %s", resp.Status)
	}

	var result CreateResponse
	// unmarshal json response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}



















