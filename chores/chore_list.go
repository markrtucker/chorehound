package chores

import (
	"context"
	"time"
)

// Chore TODO godoc
type Chore struct {
	ID        string        `json:"id"`        // TODO: UUID
	Name      string        `json:"name"`      // Name of the chore (e.g., "Take out the trash")
	Frequency time.Duration `json:"frequency"` // TODO: custom type? Part of schedule??
	Schedule  string        `json:"schedule"`  // TODO: proper type - list of people
	Current   string        `json:"current"`   // TODO: Person. Part of schedule??
}

// Repository TODO godoc
type Repository interface {
	LoadChore(ctx context.Context, id string) (*Chore, error)
	SaveChore(ctx context.Context, c Chore) error
}

// SaveTo TODO godoc
func (c Chore) SaveTo(ctx context.Context, r Repository) error {

	// TODO: populate ID (or verify uniqueness??)
	return nil
}
