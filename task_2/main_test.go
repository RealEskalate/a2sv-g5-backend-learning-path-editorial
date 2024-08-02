package main

import (
	"testing"
)

// TestCheckPalindrome tests the checkPalindrome function.
func TestCheckPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"aba", true},
		{"abc", false},
		{"racecar", true},
		{"palindrome", false},
	}

	for _, test := range tests {
		result := checkPalindrome(test.input)
		if result != test.expected {
			t.Errorf("checkPalindrome(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// TestCountWordFrequency tests the countWordFrequency function.
func TestCountWordFrequency(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"", map[string]int{}},
		{"one", map[string]int{"one": 1}},
		{"one one", map[string]int{"one": 2}},
		{"one two", map[string]int{"one": 1, "two": 1}},
		{"one two two", map[string]int{"one": 1, "two": 2}},
	}

	for _, test := range tests {
		result := countWordFrequency(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("countWordFrequency(%q) = %v; want %v", test.input, result, test.expected)
		}
		for word, count := range test.expected {
			if result[word] != count {
				t.Errorf("countWordFrequency(%q)[%q] = %v; want %v", test.input, word, result[word], count)
			}
		}
	}
}
