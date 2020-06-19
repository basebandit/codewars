package main

import "testing"

func TestToWeirdCase(t *testing.T) {

	var testCases = []struct {
		n        string
		expected string
	}{
		//This isatestLookslikeyoupassed
		{"String", "StRiNg"},
		{"Weird string case", "WeIrD StRiNg CaSe"},
		{"abc def", "AbC DeF"},
		{"ABC", "AbC"},
		{"This is a test Looks like you passed", "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"},
	}

	for _, tt := range testCases {
		actual := toWeirdCase(tt.n)
		if actual != tt.expected {
			t.Errorf("toWeirdCase('%s') = '%s'; want '%s'", tt.n, actual, tt.expected)
		}
	}

}

func TestToWeirdCase2(t *testing.T) {

	var testCases = []struct {
		n        string
		expected string
	}{
		//This isatestLookslikeyoupassed
		{"String", "StRiNg"},
		{"Weird string case", "WeIrD StRiNg CaSe"},
		{"abc def", "AbC DeF"},
		{"ABC", "AbC"},
		{"This is a test Looks like you passed", "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"},
	}

	for _, tt := range testCases {
		actual := toWeirdCase2(tt.n)
		if actual != tt.expected {
			t.Errorf("toWeirdCase('%s') = '%s'; want '%s'", tt.n, actual, tt.expected)
		}
	}

}

func BenchmarkToWeirdCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		toWeirdCase("This is a test Looks like you passed")
	}
}


func BenchmarkToWeirdCase2(b *testing.B){
	for n :=0;n <b.N; n++{
		toWeirdCase2("This is a test Looks like you passed")
	}
}