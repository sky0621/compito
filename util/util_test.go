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

func TestSeparate(t *testing.T) {
	factors := []struct {
		in  input
		exp expect
	}{
		{
			in: input{
				chars: []string{"g", "o", "l", "a", "n", "g"},
				idx:   2,
			},
		},
	}

	for i, f := range factors {
		t.Run(fmt.Printf("case no %d", i), func(t *testing.T) {
			target, remains := Separate(f.in.chars, f.in.idx)
			assert.Equal(t, f.exp.target, target)
			assert.Equal(t, f.exp.remains, remains)
		})
	}
}
