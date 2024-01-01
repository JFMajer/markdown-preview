package main

import (
	"bytes"
	"os"
	"testing"
)

const (
	inputFile  = "./testdata/example.md"
	resultFile = "example.html"
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
	if err := run(inputFile); err != nil {
		t.Fatal(err)
	}
	result, err := os.ReadFile(resultFile)
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
