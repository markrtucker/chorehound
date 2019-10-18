package chores

import (
	"context"
	"time"
)

// MockRepository TODO godoc
type MockRepository struct {
}

// Compile-time proof of interface implementation
var _ Repository = (*MockRepository)(nil)

// SaveChore TODO godoc
func (m MockRepository) SaveChore(ctx context.Context, c Chore) error {

	return nil
}

// LoadChore TODO godoc
func (m MockRepository) LoadChore(ctx context.Context, id string) (*Chore, error) {

	c := Chore{
		ID:        id,
		Name:      "Bring trash to curb",
		Frequency: time.Hour * 24 * 7, // 1 week
		Current:   "TestMe",           // TODO: Person; how to test sending messages?
	}

	return &c, nil

	// // Chore TODO godoc
	// type Chore struct {
	// 	ID        string        `json:"id"`        // TODO: UUID
	// 	Name      string        `json:"name"`      // Name of the chore (e.g., "Take out the trash")
	// 	Frequency time.Duration `json:"frequency"` // TODO: custom type? Part of schedule??
	// 	Schedule  string        `json:"schedule"`  // TODO: proper type - list of people
	// 	Current   string        `json:"current"`   // TODO: Person. Part of schedule??
	// }

	// return nil, errors.New("not-implemented")
}
