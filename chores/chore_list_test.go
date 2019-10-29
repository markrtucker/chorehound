package chores

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SaveNewChore(t *testing.T) {
	c := Chore{
		ID:   "123",
		Name: "Test",
	}
	r := MockRepository{}

	err := c.SaveTo(context.TODO(), r)

	assert.Nil(t, err, "SaveTo should not return an error")
}
