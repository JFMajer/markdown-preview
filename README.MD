# Markdown Previewer

This Go program converts a Markdown file to HTML and then opens it in your default web browser.

## Installation

Before running this program, you need to install Go on your system. You can download Go from the [official Go Downloads Page](https://golang.org/dl/).

## Usage

To use this program, you need to pass the markdown file you want to convert as a command-line argument.

```bash
go run main.go -file=example.md
```

### How it works
The program first reads the markdown file passed as a command-line argument. It then converts the markdown to HTML and sanitizes it to remove any potentially unsafe content. The HTML is then wrapped in a basic HTML structure (defined in the `HTMLHeader` and `HTMLFooter` constants).

The program creates a temporary directory (`./tmp`) to store the output HTML file. The output file's name is derived from the input markdown file's name, with a timestamp appended to make it unique.

After writing the HTML to the output file, the program opens the file in your default web browser.

## Dependencies
This program uses the following external Go packages:
- `github.com/gomarkdown/markdown` : For parsing the markdown.
- `github.com/microcosm-cc/bluemonday` : For sanitizing the HTML.

You can install these packages with `go get`:
```
go get github.com/gomarkdown/markdown
go get github.com/microcosm-cc/bluemonday
```

## Error handling
If any errors occur during the process, the program will print the error message and exit with a non-zero status code.