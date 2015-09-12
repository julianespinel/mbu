package stats

import (
	"time"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
)

func GetCPUStat() MultipleStat {
	perCpu := true
	usagePercentagePerCore, _ := cpu.CPUPercent(400 * time.Millisecond, perCpu)
	return buildCPUMultipleStat(usagePercentagePerCore)
}

func GetRAMStat() SingleStat {
	vm, _ := mem.VirtualMemory()
	return buildRamStat(*vm)
}

func GetDiskStat() SingleStat {
	path := "/"
	diskUsage, _ := disk.DiskUsage(path)
	return buildDiskStat(*diskUsage)
}

func GetAllStats(cpu MultipleStat, ram SingleStat, disk SingleStat) AllStat {
	return AllStat{ cpu, ram, disk }
}
