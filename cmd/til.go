package cmd

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/russross/blackfriday/v2"
)

// RSS-related structs
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	LastBuildDate string `xml:"lastBuildDate"`
	Items         []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func Compile() {
	fmt.Printf("compiling...")

	postsDir := "/Users/liliwilson/Documents/GitHub/til/posts"       // TODO make this configurable
	siteURL := "https://raw.githubusercontent.com/liliwilson/til/main/rss.xml"

	// initialize rss feed
	rss := RSS{
		Version: "2.0",
		Channel: Channel{
			Title:         "lili/til", // TODO configure
			Link:          siteURL,
			Description:   "my learnings",
			LastBuildDate: time.Now().Format(time.RFC1123),
		},
	}

	// read all markdown files
	err := filepath.Walk(postsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".md" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// convert markdown content to HTML using black friday
			htmlContent := blackfriday.Run(content)

			// create an RSS item for this post
			item := Item{
				Title:       info.Name(),
                Link:        fmt.Sprintf("%s/%s", "https://github.com/liliwilson/til/main/posts", info.Name()),
				Description: string(htmlContent),
				PubDate:     info.ModTime().Format(time.RFC1123),
			}

			// Add the item to the RSS feed
			rss.Channel.Items = append(rss.Channel.Items, item)
		}

		return nil
	})

	if err != nil {
		fmt.Println("error reading folder:", err)
		return
	}

	// make xml file
	output, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		fmt.Println("error marshalling XML:", err)
		return
	}

	// write rss feed to a file
	err = os.WriteFile("rss.xml", output, 0644) // TODO don't hard code file name
	if err != nil {
		fmt.Println("error writing XML file:", err)
		return
	}

	fmt.Println("rss feed generated successfully!")
}


func Tui() {
	m := initialModel()

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error starting tui %v\n", err)
	}
}

// this function is called by the tea model when we submit a file
func Save(title string, content string) {
	filename := strings.ReplaceAll(title, " ", "_")
	postsDir := "/Users/liliwilson/Documents/GitHub/til/posts"       // TODO make this configurable

	err := os.MkdirAll(postsDir, os.ModePerm)
	if err != nil {
		fmt.Printf("error creating directory: %v\n", err)
		return
	}

	mdFile := filepath.Join(postsDir, filename+".md")

	file, err := os.Create(mdFile)
	if err != nil {
		fmt.Printf("error creating file: %v\n", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("error writing to file:", err)
		return
	}

}
