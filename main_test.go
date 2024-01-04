package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	inputFile  = "./testdata/example.md"
	goldenFile = "./testdata/example.html"
)

func TestParseContent(t *testing.T) {
	md, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result := parseMarkdown(md)
	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(expected, result) {
		t.Errorf("expected and result not match")
	}
}

func TestRun(t *testing.T) {
	var buffer bytes.Buffer
	err := run(inputFile, &buffer, true)
	if err != nil {
		t.Fatal(err)
	}
	resultFile := strings.TrimSpace(buffer.String())
	fmt.Printf("result file is: %s\n", resultFile)
	result, err := os.ReadFile(filepath.Join("./tmp", resultFile))
	fmt.Printf("created path is: %s\n", filepath.Join("./tmp", resultFile))
	if err != nil {
		t.Fatal(err)
	}
	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expected, result) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("result content does not match golden file")
	}
	os.Remove(resultFile)
}
