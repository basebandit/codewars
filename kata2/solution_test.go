package main

import "testing"

func TestDupCount(t *testing.T) {
	var testCases = []struct {
		n        string
		expected int
	}{
		{"abcde", 0},
		{"aabbcde", 2},
		{"aabBcde", 2},
		{"indivisibility", 1},
		{"Indivisibilities", 2},
		{"aA11", 2},
		{"ABBA", 2},
		{"abcdea", 1},
		{"abcdeaB11", 3},
	}

	for _, tt := range testCases {
		actual := dupCount(tt.n)
		if actual != tt.expected {
			t.Errorf("countDuplicates('%s') = '%d'; want '%d'", tt.n, actual, tt.expected)
		}
	}
}

func TestDupCount2(t *testing.T) {
	var testCases = []struct {
		n        string
		expected int
	}{
		{"abcde", 0},
		{"aabbcde", 2},
		{"aabBcde", 2},
		{"indivisibility", 1},
		{"Indivisibilities", 2},
		{"aA11", 2},
		{"ABBA", 2},
		{"abcdea", 1},
		{"abcdeaB11", 3},
	}

	for _, tt := range testCases {
		actual := dupCount2(tt.n)
		if actual != tt.expected {
			t.Errorf("countDuplicates('%s') = '%d'; want '%d'", tt.n, actual, tt.expected)
		}
	}
}

func BenchmarkDupCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dupCount("Indivisibilities")
	}
}

func BenchmarkDupCount2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dupCount2("Indivisibilities")
	}
}
