package chores

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func Test_ReadEmptyJson_Success(t *testing.T) {
	fs := FileSystemRepository{
		fs: afero.NewMemMapFs(),
	}

	f, err := fs.fs.Create("/fred.json")
	assert.Nil(t, err, "File setup should not return an error")

	n3, err := f.WriteString("{}")
	assert.Nil(t, err, "File write should not return an error")
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()

	c, err := fs.LoadChore(context.TODO(), "fred")

	assert.Nil(t, err, "LoadChore should not return an error")
	assert.NotNil(t, c, "LoadChore should return a chore")
}

func Test_ReadInvalidJson_Error(t *testing.T) {
	fs := FileSystemRepository{
		fs: afero.NewMemMapFs(),
	}

	f, err := fs.fs.Create("/fred.json")
	assert.Nil(t, err, "File setup should not return an error")

	n3, err := f.WriteString("{")
	assert.Nil(t, err, "File write should not return an error")
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()

	c, err := fs.LoadChore(context.TODO(), "fred")

	assert.NotNil(t, err, "LoadChore should return an error")
	assert.Nil(t, c, "LoadChore should not return a chore")
}

func Test_ReadNoFile_Error(t *testing.T) {
	fs := FileSystemRepository{
		fs: afero.NewMemMapFs(),
	}

	c, err := fs.LoadChore(context.TODO(), "fred")

	assert.NotNil(t, err, "LoadChore should return an error")
	assert.Nil(t, c, "LoadChore should not return a chore")
}
