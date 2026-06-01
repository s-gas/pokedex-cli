package input

import (
	"slices"
	"testing"
)

type Test struct {
	info  string
	input string
	want  []string
}

var tests = []Test{
	{
		info:  "One word",
		input: "Pikachu",
		want:  []string{"pikachu"},
	},
	{
		info:  "Two words",
		input: "Pikachu Ditto",
		want:  []string{"pikachu", "ditto"},
	},
	{
		info:  "Empty string",
		input: "",
		want:  []string{},
	},
	{
		info:  "All uppercase",
		input: "PIKACHU",
		want:  []string{"pikachu"},
	},
	{
		info:  "Mixed lower and upper case",
		input: "piKacHu dIttO",
		want:  []string{"pikachu", "ditto"},
	},
	{
		info:  "Digits",
		input: "piKacHu d1tt0",
		want:  []string{"pikachu", "d1tt0"},
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
