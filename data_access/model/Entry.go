package model

import "time"

type Entry struct {
	Name      string
	Role      string
	DateAdded time.Time
}
