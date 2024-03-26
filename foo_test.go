package main

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{1, -1, 0},
		{0, 0, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("case(%d+%d=%d)", test.a, test.b, test.want),
			func(t *testing.T) {
				got := Foo(test.a, test.b)
				if got != test.want {
					t.Errorf("got %d, want %d", got, test.want)
				}
			})
	}
}
