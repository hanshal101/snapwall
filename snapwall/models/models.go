package models

import (
	"time"
)

type Log struct {
	Time        time.Time `json:"time"`
	Type        string    `json:"type"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	Port        string    `json:"port"`
	Protocol    string    `json:"protocol"`
}
