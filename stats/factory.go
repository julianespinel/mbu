package stats

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
)

const GB = float32(1000000000)

func getAverage(array []float32) float32 {

	accumulator := float32(0)
	size := float32(len(array))

	for _, number := range array {
		accumulator += number
	}

	return (accumulator/size)
}

func buildCPUMultipleStat(usagePercentagePerCore []float32) MultipleStat {

	averageUsagePercentage := getAverage(usagePercentagePerCore)
	return MultipleStat{ averageUsagePercentage, usagePercentagePerCore }
}

func buildRamStat(vm mem.VirtualMemoryStat) SingleStat {
	totalGB := float32(vm.Total) / GB
	availableGB := float32(vm.Available) / GB
	usedGB := (totalGB - availableGB)
	usagePercentage := float32(vm.UsedPercent)
	return SingleStat{ totalGB, usedGB, availableGB, usagePercentage }
}

func buildDiskStat(diskUsage disk.DiskUsageStat) SingleStat {
	totalGB := float32(diskUsage.Total) / GB
	availableGB := float32(diskUsage.Free) / GB
	usedGB := float32(diskUsage.Used) / GB
	usagePercentage := float32(diskUsage.UsedPercent)
	return SingleStat{ totalGB, usedGB, availableGB, usagePercentage }
}
