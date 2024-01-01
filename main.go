package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
)

const (
	HTMLHeader = `<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html; charset=UTF-8">
	<title>Markdown Previewer</title>
	</head>
	<body>
	`
	HTMLFooter = `</body>
	</html>`
)

func main() {

	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := run(*filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func run(filename string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseMarkdown(input)
	htmlComplete := generateHTML(HTMLHeader, htmlData, HTMLFooter)
	baseFilename := strings.TrimSuffix(filepath.Base(filename), ".md")
	outName := fmt.Sprintf("%s.html", baseFilename)
	return saveHTML(outName, htmlComplete)
}

func parseMarkdown(input []byte) []byte {
	unsafeHTML := markdown.ToHTML(input, nil, nil)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML)

	return html
}

func saveHTML(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func generateHTML(header string, body []byte, footer string) []byte {
	var b bytes.Buffer
	b.WriteString(header)
	b.Write(body)
	b.WriteString(footer)

	return b.Bytes()
}
