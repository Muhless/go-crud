package entities

import "time"

type Category struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}