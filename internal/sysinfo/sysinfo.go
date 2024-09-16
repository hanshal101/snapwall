package sysinfo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall/models"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"golang.org/x/sys/unix"
)

func getSysInfo() (*models.SystemInfo, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	var diskInfo []disk.UsageStat
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err == nil {
			diskInfo = append(diskInfo, *usage)
		}
	}

	hostInfo, err := host.Info()
	if err != nil {
		return nil, err
	}

	uptime, err := host.Uptime()
	if err != nil {
		return nil, err
	}

	return &models.SystemInfo{
		CPUInfo:    cpuInfo,
		MemoryInfo: memInfo,
		DiskInfo:   diskInfo,
		HostInfo:   hostInfo,
		Uptime:     uptime,
	}, nil
}

func GetSystemInfo(c *gin.Context) {
	sysInfo, err := getSysInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch system information",
		})
		return
	}

	c.JSON(http.StatusOK, sysInfo)

}

// Node
type SystemInfo struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`
	Timestamp   int64   `json:"timestamp"`
}

func getCPUUsage() float64 {
	percentages, err := cpu.Percent(time.Second, false) // Get CPU usage percentage for all cores (false)
	if err != nil {
		fmt.Printf("Error fetching CPU usage: %v\n", err)
		return 0
	}
	return percentages[0] // Return the percentage of CPU usage for all cores
}

// Function to get Memory Usage
func getMemoryUsage() float64 {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error fetching memory usage: %v\n", err)
		return 0
	}
	return vmStat.UsedPercent // Percentage of used memory
}

// Function to get Disk Usage
func getDiskUsage() float64 {
	var stat unix.Statfs_t
	unix.Statfs("/", &stat)
	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bavail * uint64(stat.Bsize)
	return float64(total-free) / float64(total) * 100
}

// API handler to serve system info
func ServeNodeInfo(c *gin.Context) {
	info := SystemInfo{
		CPUUsage:    getCPUUsage(),
		MemoryUsage: getMemoryUsage(),
		DiskUsage:   getDiskUsage(),
		Timestamp:   time.Now().Unix(),
	}
	c.JSON(http.StatusOK, info)
}
