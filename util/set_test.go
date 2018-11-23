package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := NewSet()
	assert.Equal(t, []string{}, s.List())

	s.Add("Golang")
	assert.Equal(t, []string{"Golang"}, s.List())

	s.Add("Java")
	assert.ElementsMatch(t, []string{"Golang", "Java"}, s.List())

	// 追加済みの要素を再追加
	s.Add("Golang")
	assert.ElementsMatch(t, []string{"Golang", "Java"}, s.List())

	// 追加済みの要素を再追加
	s.Add("Java")
	assert.ElementsMatch(t, []string{"Golang", "Java"}, s.List())
}
