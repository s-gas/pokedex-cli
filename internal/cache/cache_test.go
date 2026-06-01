package cache

import (
	"testing"
	"time"
)

type Test struct {
	info 		string
	inputA	string
	inputB	[]byte
}

var tests = []Test{
	{
		info: 	"example.com",
		inputA:	"http://example.com",
		inputB:	[]byte("data"),
	},
	{
		info:		"pokeapi",
		inputA:	"https://pokeapi.co/api/v2/location-area",
		inputB:	[]byte("data"),
	},
}

func TestCache(t *testing.T) {
	for _, test := range tests {
		cache := New(1 * time.Second) 
		cache.Add(test.inputA, test.inputB)
		val, ok := cache.Get(test.inputA); 
		if !ok {
			t.Errorf("%q: expected to get entry\n", test.info)
		}
		if string(val) != string(test.inputB) {
			t.Errorf("%q: expected to get value\n", test.info)
		}
		time.Sleep(1 * time.Second)
		if _, ok := cache.Get(test.inputA); ok {
			t.Errorf("%q: expected not to get entry\n", test.info)
		}
	}
}
