// This is likely a CLI for GitHub issues.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var usage = "%s Usage:\n\tsearch QUERY\nOr:\n\t[read|edit|close|open] OWNER REPO ISSUE_NUMBER\n"

func usageDie() {
	fmt.Fprintf(os.Stderr, usage, os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usageDie()
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" {
		if len(args) < 1 {
			usageDie()
		}
		search(args)
		os.Exit(0)
	}
	if len(args) != 3 {
		usageDie()
	}
	owner, repo, number := args[0], args[1], args[2]
	switch cmd {
	case "read":
		readIssue(owner, repo, number)
	case "edit":
		editIssue(owner, repo, number)
	case "close":
		closeIssue(owner, repo, number)
	case "open":
		openIssue(owner, repo, number)
	}
}

func search(query []string) {
	result, err := SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		format := "#%-6d\t%20.20s\t%.100s\n"
		fmt.Printf(format, item.Number, item.User.Login, item.Title)
	}
}

func readIssue(owner, repo, number string) {
	issue, err := GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	body := issue.Body
	if body == "" {
		body = "<empty>\n"
	}
	fmt.Printf("repo: %s/%s\nnumber: %s\nuser: %s\ntitle: %s\n\n%s",
		owner, repo, number, issue.User.Login, issue.Title, body)
}

func editIssue(owner string, repo string, number string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}
	tempfile, err := ioutil.TempFile("", "issue_crud")
	if err != nil {
		log.Fatal(err)
	}
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	issue, err := GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(tempfile)
	err = encoder.Encode(map[string]string{
		"title": issue.Title,
		"state": issue.State,
		"body":  issue.Body,
	})
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tempfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	tempfile.Seek(0, 0)
	fields := make(map[string]string)
	if err = json.NewDecoder(tempfile).Decode(&fields); err != nil {
		log.Fatal(err)
	}

	_, err = EditIssue(owner, repo, number, fields)
	if err != nil {
		log.Fatal(err)
	}
}

func closeIssue(owner string, repo string, number string) {
	_, err := EditIssue(owner, repo, number, map[string]string{"state": "closed"})
	if err != nil {
		log.Fatal(err)
	}
}

func openIssue(owner string, repo string, number string) {
	_, err := EditIssue(owner, repo, number, map[string]string{"state": "open"})
	if err != nil {
		log.Fatal(err)
	}
}
