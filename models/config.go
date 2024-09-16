package models

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"gorm.io/gorm"
)

type SEVERITY string

const (
	SEVERITY_LOW    SEVERITY = "LOW"
	SEVERITY_MEDIUM SEVERITY = "MEDIUM"
	SEVERITY_HIGH   SEVERITY = "HIGH"
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
	Severity    string    `json:"severity"`
}

type SystemInfo struct {
	CPUInfo    []cpu.InfoStat         `json:"cpu_info"`
	MemoryInfo *mem.VirtualMemoryStat `json:"memory_info"`
	DiskInfo   []disk.UsageStat       `json:"disk_info"`
	HostInfo   *host.InfoStat         `json:"host_info"`
	Uptime     uint64                 `json:"uptime"`
}

// APPLICATION Models
type Application struct {
	gorm.Model
	Name        string `json:"name"`
	Port        string `json:"port"`
	Description string `json:"description"`
	Tags        []Tags `json:"tags" gorm:"foreignKey:ApplicationID;constraint:OnDelete:CASCADE;"`
}

type Tags struct {
	ApplicationID uint   `json:"application_id"`
	Tag           string `json:"tag"`
}
