package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
)

type content struct {
	Title string
	Body  template.HTML
}

const (
	// Switching to go templates
	defaultTemplate = `<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html; charset=UTF-8">
	<title>{{ .Title }}</title>
	</head>
	<body>
		{{ .Body }}
	</body>
	</html>
	`
)

func main() {
	// Parse command line arguments
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	flag.Parse()
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Call the run function and handle any errors
	err := run(*filename, os.Stdout, *skipPreview)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(filename string, out io.Writer, skipPreview bool) error {
	// Create a temporary directory to store the output HTML files
	err := os.Mkdir("./tmp", 0755)
	// If the directory exists ignore the error
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	// Read the input file
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Convert the Markdown to HTML
	htmlData := parseMarkdown(input)
	// Combine the HTML header, body, and footer
	htmlComplete, err := generateHTML(htmlData)
	if err != nil {
		return err
	}
	// Create the output filename, adding a timestamp to make it unique
	baseFilename := strings.TrimSuffix(filepath.Base(filename), ".md")
	timestamp := time.Now().Format("20060102-150405")
	outName := fmt.Sprintf("%s-%s.html", baseFilename, timestamp)
	_, err = out.Write([]byte(outName + "\n"))
	if err != nil {
		return err
	}
	// Save the output HTML to a file in the temporary directory
	err = saveHTML(outName, htmlComplete, skipPreview)
	if err != nil {
		return err
	}
	return nil

}

func parseMarkdown(input []byte) []byte {
	// Convert the Markdown to HTML, then sanitize it to remove any potentially unsafe content
	unsafeHTML := markdown.ToHTML(input, nil, nil)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML)

	return html
}

func saveHTML(filename string, data []byte, skipPreview bool) error {
	// Write the HTML data to a file in the temporary directory
	err := os.WriteFile(filepath.Join("./tmp", filename), data, 0644)
	if err != nil {
		return err
	}
	// Open the HTML file in the default web browser
	fullPath := filepath.Join("./tmp", filename)
	openWebBrowser(fullPath, skipPreview)
	return nil
}

func generateHTML(body []byte) ([]byte, error) {
	// Parse the content of defaultTemplate into new template
	tmpl, err := template.New("default").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}
	// Create a new content struct with the title and body
	c := content{
		Title: "My Markdown Preview",
		Body:  template.HTML(body),
	}
	// Create a new buffer to hold the output HTML
	var buf bytes.Buffer
	// Execute the template with the content struct and write the output to the buffer
	if err := tmpl.Execute(&buf, c); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func openWebBrowser(filename string, skipPreview bool) error {
	// defer os.Remove(filename) // Remove the temporary HTML file after opening the browser)
	// Open the output HTML file in the default web browser
	if skipPreview {
		return nil
	}
	cmd := exec.Command("xdg-open", filename)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to open web browser: %w", err)
	}
	time.Sleep(time.Second)
	return nil

}
