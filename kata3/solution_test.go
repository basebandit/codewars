package main

import "testing"

func TestVowelCount(t *testing.T) {
	testCases := []struct {
		n        string
		expected int
	}{
		{"abracadabra", 5},
		{"elephant", 3},
		{"example", 3},
		{"consider", 3},
		{"Indivisibilities", 8},
	}

	for _, tt := range testCases {
		actual := vowelCount(tt.n)

		if actual != tt.expected {
			t.Errorf("vowelCount('%s') = %d; want %d", tt.n, actual, tt.expected)
		}
	}
}

func BenchmarkVowelCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		vowelCount("Indivisibilities")
	}
}
