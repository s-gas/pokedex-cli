package main

import (
	"testing"
	"slices"
)

type Test struct {
	info	string
	input	string
	want	[]string
}

var tests = []Test{
	{
		info:		"One word",
		input: 	"Pikachu",
		want:		[]string{"Pikachu"},
	},
	{
		info:		"Two words",
		input:	"Pikachu Ditto",
		want:		[]string{"Pikachu", "Ditto"},
	},
	{
		info:		"Empty string",
		input:	"",
		want:		[]string{},
	},
	{	
		info:		"All uppercase",
		input:	"PIKACHU",
		want:		[]string{"Pikachu"},
	},
}

func TestCleanInput(t *testing.T) {
	for _, test := range tests {
		got := cleanInput(test.input)
		if !slices.Equal(got, test.want) {
			t.Errorf("%q: got %v, want %v\n", test.info, got, test.want)
		}
	}
}
