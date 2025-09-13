package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetH1FromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "simple_and_small",
			inputBody: "<html><h1>scooby doo</h1></html>",
			expected:  "scooby doo",
		},
		{
			name:      "no_H1",
			inputBody: "<html><h2>not a h1</h2></html>",
			expected:  "",
		},
		{
			name:      "nested_html",
			inputBody: "<html><body><section><h1>nested</h1></section></body></html>",
			expected:  "nested",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.inputBody)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "simple",
			inputBody: "<html><body><p>simple</p></body></html>",
			expected:  "simple",
		},
		{
			name:      "missing",
			inputBody: "<html><body>simple</body></html>",
			expected:  "",
		},
		{
			name:      "multiple",
			inputBody: "<html><body><p>first</p><p>second</p></body></html>",
			expected:  "first",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {

	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "link",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a href="https://blog.boot.dev"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://blog.boot.dev"},
		},
		{
			name:      "images",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src="https://blog.boot.png"/></body></html>`,
			expected:  []string{"https://blog.boot.png"},
		},
		{
			name:      "relative_links",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a href="/foo/bar"><span>Boot.dev</span></a><img src="/favicon.png"/></body></html>`,
			expected:  []string{"https://blog.boot.dev/foo/bar", "https://blog.boot.dev/favicon.png"},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseUrl, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("couldn't parse input URL: %v", err)
				return
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseUrl)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
