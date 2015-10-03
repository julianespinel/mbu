package stats

import (
	"github.com/mbu/Godeps/_workspace/src/github.com/shirou/gopsutil/cpu"
	"github.com/mbu/Godeps/_workspace/src/github.com/shirou/gopsutil/disk"
	"github.com/mbu/Godeps/_workspace/src/github.com/shirou/gopsutil/mem"
	"time"
)

func GetCPUStat() MultipleStat {
	perCpu := true
	usagePercentagePerCore, _ := cpu.CPUPercent(400*time.Millisecond, perCpu)
	return BuildCPUMultipleStat(usagePercentagePerCore)
}

func GetRAMStat() SingleStat {
	vm, _ := mem.VirtualMemory()
	return BuildRamStat(*vm)
}

func GetDiskStat() SingleStat {
	path := "/"
	diskUsage, _ := disk.DiskUsage(path)
	return BuildDiskStat(*diskUsage)
}

func GetAllStats(cpu MultipleStat, ram SingleStat, disk SingleStat) AllStat {
	return AllStat{cpu, ram, disk}
}
