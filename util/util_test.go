package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	chars []string
	idx   int
}

type expect struct {
	target    string
	remaining []string
}

func TestSeparateOK(t *testing.T) {
	factors := []struct {
		in  input
		exp expect
	}{
		{
			in:  input{chars: []string{"c", "a", "t"}, idx: 0},
			exp: expect{target: "c", remaining: []string{"a", "t"}},
		},
		{
			in:  input{chars: []string{"c", "a", "t"}, idx: 1},
			exp: expect{target: "a", remaining: []string{"c", "t"}},
		},
		{
			in:  input{chars: []string{"c", "a", "t"}, idx: 2},
			exp: expect{target: "t", remaining: []string{"c", "a"}},
		},
		{
			in:  input{chars: []string{"c", "a", "t"}, idx: 3},
			exp: expect{target: "", remaining: []string{}},
		},
		{
			in:  input{chars: []string{"c", "a", "t"}, idx: -1},
			exp: expect{target: "", remaining: []string{}},
		},
		{
			in:  input{chars: []string{}, idx: 0},
			exp: expect{target: "", remaining: []string{}},
		},
		{
			in:  input{chars: nil, idx: 0},
			exp: expect{target: "", remaining: []string{}},
		},
	}

	for i, f := range factors {
		t.Run(fmt.Sprintf("case no %d", i), func(t *testing.T) {
			target, remaining := Separate(f.in.chars, f.in.idx)
			assert.Equal(t, f.exp.target, target)
			assert.Equal(t, f.exp.remaining, remaining)
		})
	}
}
