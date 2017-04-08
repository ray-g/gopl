// Issuesreport prints a report of issues matching the search terms.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/template"
	"time"
)

var stdout io.Writer = os.Stdout

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	printByDaysAgo(result)
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func printByDaysAgo(result *IssuesSearchResult) {
	now := time.Now()

	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	var dayIssues IssuesSearchResult
	var monthIssues IssuesSearchResult
	var yearIssues IssuesSearchResult
	var otherIssues IssuesSearchResult

	for _, issue := range result.Items {
		switch {
		case issue.CreatedAt.After(day):
			(dayIssues.TotalCount)++
			dayIssues.Items = append(dayIssues.Items, issue)
		case issue.CreatedAt.After(month):
			(monthIssues.TotalCount)++
			monthIssues.Items = append(monthIssues.Items, issue)
		case issue.CreatedAt.After(year):
			(yearIssues.TotalCount)++
			yearIssues.Items = append(yearIssues.Items, issue)
		default:
			(otherIssues.TotalCount)++
			otherIssues.Items = append(otherIssues.Items, issue)
		}
	}
	fmt.Fprint(stdout, "Last day:\t")
	if err := report.Execute(stdout, monthIssues); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(stdout, "Last month:\t")
	if err := report.Execute(stdout, monthIssues); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(stdout, "\nLast year:\t")
	if err := report.Execute(stdout, yearIssues); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(stdout, "\nLong long ago:\t")
	if err := report.Execute(stdout, otherIssues); err != nil {
		log.Fatal(err)
	}
}
