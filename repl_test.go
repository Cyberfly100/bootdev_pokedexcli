package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"basic":      {input: "hello world", want: []string{"hello", "world"}},
		"mixed case": {input: "Charmander Bulbasaur PIKACHU", want: []string{"charmander", "bulbasaur", "pikachu"}},
		"whitespace": {input: "  leading and trailing  whitespace  ", want: []string{"leading", "and", "trailing", "whitespace"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
