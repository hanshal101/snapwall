package models

import (
	"time"

	"gorm.io/gorm"
)

type Policy struct {
	gorm.Model
	Name  string `json:"name"`
	Type  string `json:"type"`
	IPs   []IP   `json:"ips" gorm:"foreignKey:PolicyID;constraint:OnDelete:CASCADE;"`
	Ports []Port `json:"ports" gorm:"foreignKey:PolicyID;constraint:OnDelete:CASCADE;"`
}

type IP struct {
	gorm.Model
	PolicyID uint   `json:"policy_id"`
	Address  string `json:"address"`
}

type Port struct {
	gorm.Model
	PolicyID uint   `json:"policy_id"`
	Number   string `json:"number"`
}

type Log struct {
	Time        time.Time `json:"time"`
	Type        string    `json:"type"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	Port        string    `json:"port"`
	Protocol    string    `json:"protocol"`
}
