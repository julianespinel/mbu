package stats

import (
	"github.com/mbu/Godeps/_workspace/src/github.com/shirou/gopsutil/disk"
	"github.com/mbu/Godeps/_workspace/src/github.com/shirou/gopsutil/mem"
)

const GB = float64(1000000000)

func GetAverage(array []float64) float64 {

	average := float64(0)
	size := float64(len(array))

	if size > 0 {

		accumulator := float64(0)

		for _, number := range array {
			accumulator += number
		}

		average = (accumulator / size)
	}

	return average
}

func BuildCPUMultipleStat(usagePercentagePerCore []float64) MultipleStat {

	numberOfCores := len(usagePercentagePerCore)
	averageUsagePercentage := GetAverage(usagePercentagePerCore)
	return MultipleStat{numberOfCores, averageUsagePercentage, usagePercentagePerCore}
}

func BuildRamStat(vm mem.VirtualMemoryStat) SingleStat {
	totalGB := float64(vm.Total) / GB
	availableGB := float64(vm.Available) / GB
	usedGB := (totalGB - availableGB)
	usagePercentage := float64(vm.UsedPercent)
	return SingleStat{totalGB, usedGB, availableGB, usagePercentage}
}

func BuildDiskStat(diskUsage disk.DiskUsageStat) SingleStat {
	totalGB := float64(diskUsage.Total) / GB
	availableGB := float64(diskUsage.Free) / GB
	usedGB := float64(diskUsage.Used) / GB
	usagePercentage := float64(diskUsage.UsedPercent)
	return SingleStat{totalGB, usedGB, availableGB, usagePercentage}
}
