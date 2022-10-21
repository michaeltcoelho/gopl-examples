package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	// "os"
	"strings"
	"text/template"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    //Markdown formatted
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

const issuesListText = `{{.TotalCount}} issues:
{{range .Items}}-----------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

const issuesListHTML = `
<h1>Search issues</h1>
<form action="/" method="GET">
  <label>Query</label>
  <input type="text" name="terms" id="terms">
</form>
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
  <th>Age</th>
<tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td>{{.CreatedAt | daysAgo}} days</td>
</tr>
{{end}}
</table>
`

var issuesListHTMLTpl = template.Must(template.New("issueslisthtml").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(issuesListHTML))

var issuesListTextTpl = template.Must(template.New("issuelisttext").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(issuesListText))

func main() {
	// Send issues to Stdout
	// result, err := SearchIssues(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := issuesListHTMLTpl.Execute(os.Stdout, result); err != nil {
	// 	log.Fatal(err)
	// }

	// Without using templates
	// fmt.Printf("%d issues:\n", result.TotalCount)
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	// }
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		terms := r.URL.Query().Get("terms")
		if terms == "" || terms == " " {
			issuesListHTMLTpl.Execute(w, nil)
			return
		}
		result, err := SearchIssues(strings.Split(terms, " "))
		if err != nil {
			http.NotFound(w, r)
		}
		issuesListHTMLTpl.Execute(w, result)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
