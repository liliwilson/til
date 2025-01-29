package cmd

import (
	"encoding/xml"
	// "errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/russross/blackfriday/v2"
	// "github.com/charmbracelet/bubbletea" // TODO make bubbletea editing ui
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

func New(filename string) {
	// TODO add flag to use or not use the template...
	templateFile := "posts/template.md" // TODO make this a configurable thing
	postsDir := "./posts"               // TODO make this a configurable thing

	err := os.MkdirAll(postsDir, os.ModePerm)
	if err != nil {
		fmt.Printf("error creating directory: %v\n", err)
		return
	}

	mdFile := filepath.Join(postsDir, filename+".md")

	if _, err := os.Stat(templateFile); os.IsNotExist(err) {
		fmt.Printf("template file does not exist: %s\n", templateFile)
		return
	}

	templateContent, err := os.ReadFile(templateFile)
	if err != nil {
		fmt.Printf("error reading template file: %v\n", err)
		return
	}

	err = os.WriteFile(mdFile, templateContent, os.ModePerm)
	if err != nil {
		fmt.Printf("error writing to new file: %v\n", err)
		return
	}
}

func Compile() {
	fmt.Printf("compiling...")

	postsDir := "./posts"              // TODO make this configurable
	siteURL := "http://localhost:8000" // TODO change to my site URL

	// initialize rss feed
	rss := RSS{
		Version: "2.0",
		Channel: Channel{
			Title:         "lili til", // TODO configure
			Link:          siteURL,
			Description:   "my posts",
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
				Link:        fmt.Sprintf("%s/%s", siteURL, info.Name()),
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
	err = os.WriteFile("rss_feed.xml", output, 0644) // TODO don't hard code file name
	if err != nil {
		fmt.Println("error writing XML file:", err)
		return
	}

	fmt.Println("rss feed generated successfully!")
}
