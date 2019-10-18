package chores

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/spf13/afero"
)

// var AppFs = afero.NewMemMapFs()

// FileSystemRepository TODO godoc
type FileSystemRepository struct {
	fs   afero.Fs
	base string
}

var _ Repository = (*FileSystemRepository)(nil)

// NewFileSystemRepository TODO godoc
func NewFileSystemRepository(base string) Repository {

	repo := FileSystemRepository{
		fs:   afero.NewOsFs(),
		base: base,
	}

	return &repo
}

// SaveChore TODO godoc
func (fs *FileSystemRepository) SaveChore(ctx context.Context, c Chore) error {

	return errors.New("save-chore-not-implemented")
}

// LoadChore TODO godoc
func (fs *FileSystemRepository) LoadChore(ctx context.Context, id string) (*Chore, error) {

	f := fs.getFs()

	// Load file: "<id>.json"
	file, err := f.Open(fs.baseDir() + id + ".json")
	if err != nil {
		// TODO Logging
		return nil, err // TODO wrap error
	}

	// TODO: Deserialize JSON into Chore struct
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		// TODO Logging
		return nil, err // TODO wrap error
	}

	c := Chore{}
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		// TODO Logging
		return nil, err // TODO wrap error
	}

	return &c, nil
}

func (fs *FileSystemRepository) baseDir() string {
	// TODO: verify that baseDir ends in "/"

	// Default to root base dir if another base dir has not been provided
	if fs.base == "" {
		fs.base = "/"
	}

	return fs.base
}

func (fs *FileSystemRepository) getFs() afero.Fs {
	// Default to OS file system if another FS has not been provided
	if fs.fs == nil {
		fs.fs = afero.NewOsFs()
	}

	return fs.fs
}
