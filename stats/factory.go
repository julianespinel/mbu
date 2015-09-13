package stats

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
)

const GB = float32(1000000000)

func GetAverage(array []float32) float32 {

	average := float32(0)
	size := float32(len(array))

	if (size > 0) {

		accumulator := float32(0)

		for _, number := range array {
			accumulator += number
		}

		average = (accumulator / size)
	}

	return average
}

func BuildCPUMultipleStat(usagePercentagePerCore []float32) MultipleStat {

	averageUsagePercentage := GetAverage(usagePercentagePerCore)
	return MultipleStat{ averageUsagePercentage, usagePercentagePerCore }
}

func BuildRamStat(vm mem.VirtualMemoryStat) SingleStat {
	totalGB := float32(vm.Total) / GB
	availableGB := float32(vm.Available) / GB
	usedGB := (totalGB - availableGB)
	usagePercentage := float32(vm.UsedPercent)
	return SingleStat{ totalGB, usedGB, availableGB, usagePercentage }
}

func BuildDiskStat(diskUsage disk.DiskUsageStat) SingleStat {
	totalGB := float32(diskUsage.Total) / GB
	availableGB := float32(diskUsage.Free) / GB
	usedGB := float32(diskUsage.Used) / GB
	usagePercentage := float32(diskUsage.UsedPercent)
	return SingleStat{ totalGB, usedGB, availableGB, usagePercentage }
}
