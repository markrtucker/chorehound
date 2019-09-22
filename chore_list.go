package main

import "time"

type chore struct {
	ID        string        `json:"id"`        // TODO: UUID
	Name      string        `json:"name"`      // Name of the chore (e.g., "Take out the trash")
	Frequency time.Duration `json:"frequency"` // TODO: custom type? Part of schedule??
	Schedule  string        `json:"schedule"`  // TODO: proper type - list of people
	Current   string        `json:"current"`   // TODO: Person. Part of schedule??
}

// Use cases:
// Create a new chore and store it in bucket (S3/OS??)
// Given a chore ID, read the chore JSON from the bucket
// Given a chore, mark it as done (advance the schedule)
// Given a chore, hound the person scheduled to do it
