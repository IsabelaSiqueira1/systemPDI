package entities

import "time"

type Service struct {
	ID        string
	Name      string
	CreatedAt time.Time
}