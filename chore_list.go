package main

import "time"

type chore struct {
	ID        string        `json:"id"` // TODO: UUID
	Name      string        `json:"name"`
	Frequency time.Duration `json:"frequency"` // TODO: custom type? Part of schedule??
	Schedule  string        `json:"schedule"`  // TODO: proper type - list of people
	Current   string        `json:"current"`   // TODO: Person. Part of schedule??
}
