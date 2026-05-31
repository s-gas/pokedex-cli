package main

import (
	"testing"
	"slices"
)

type Test struct {
	input	string
	want	[]string
}

var tests = []Test{
	{
		input: 	"Pikachu",
		want:		[]string{"Pikachu"},
	},
}

func TestCleanInput(t *testing.T) {
	for _, test := range tests {
		got := cleanInput(test.input)
		if !slices.Equal(got, test.want) {
			t.Errorf("got %v, want %v\n", got, test.want)
		}
	}
}
