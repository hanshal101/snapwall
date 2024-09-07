package sysinfo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall/models"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
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
