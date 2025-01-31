package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var path = "/Users/liliwilson/Documents/GitHub/til"

func Tui() {
	m := initialModel()

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error starting tui %v\n", err)
	}
}

// this function is called by the tea model when we submit a file
func Save(title string, content string) {
	filename := strings.ReplaceAll(title, " ", "-")
	postsDir := path + "/_posts" // TODO make this configurable

    loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	currentTime := time.Now().In(loc)

	formattedTime := currentTime.Format("2006-01-02 15:04:05 -0700")
	jekyll_header := fmt.Sprintf(`---
title: %s 
date: %s
layout: plain
---`, title, formattedTime)

	err = os.MkdirAll(postsDir, os.ModePerm)
	if err != nil {
		fmt.Printf("error creating directory: %v\n", err)
		return
	}

    currentDate := currentTime.Format("2006-01-02")
	mdFile := filepath.Join(postsDir, currentDate+"-"+filename+".md")

	file, err := os.Create(mdFile)
	if err != nil {
		fmt.Printf("error creating file: %v\n", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(jekyll_header + "\n" + content)
	if err != nil {
		fmt.Println("error writing to file:", err)
		return
	}

}

// TODO there is definitely a better way to do this. gh action maybe?
func Publish() {
	dir := "/Users/liliwilson/Documents/GitHub/til/"
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("error with chdir", err)
		return
	}

	// git add the "posts/" directory and "rss.xml" file
	cmd := exec.Command("git", "add", "posts/", "rss.xml")
	err = cmd.Run()
	if err != nil {
		fmt.Println("error with git add", err)
		return
	}

	// commit the changes
	cmd = exec.Command("git", "commit", "-m", "updating rss feed")
	err = cmd.Run()
	if err != nil {
		fmt.Println("error with git push", err)
		return
	}

	// push the changes
	cmd = exec.Command("git", "push")
	err = cmd.Run()
	if err != nil {
		fmt.Println("error with git push", err)
		return
	}

	fmt.Println("changes committed and pushed successfully!")
}
