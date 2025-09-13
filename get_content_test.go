package main

import (
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
