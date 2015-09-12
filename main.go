package main

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"./stats"
)

func getAverage(array []float32) float32 {

	accumulator := float32(0)
	size := float32(len(array))

	for _, number := range array {
		accumulator += number
	}

	return (accumulator/size)
}

func buildCPUMultipleStat(usagePercentagePerCore []float32) stats.MultipleStat {

	averageUsagePercentage := getAverage(usagePercentagePerCore)
	return stats.MultipleStat{ averageUsagePercentage, usagePercentagePerCore }
}

func buildRamStat(vm mem.VirtualMemoryStat) stats.SingleStat {
	gb := float32(1000000000)
	totalGB := float32(vm.Total) / gb
	availableGB := float32(vm.Available) / gb
	usedGB := (totalGB - availableGB)
	usagePercentage := float32(vm.UsedPercent)
	return stats.SingleStat{ totalGB, usedGB, availableGB, usagePercentage }
}

func buildDiskStat(diskUsage disk.DiskUsageStat) stats.SingleStat {
	gb := float32(1000000000)
	totalGB := float32(diskUsage.Total) / gb
	availableGB := float32(diskUsage.Free) / gb
	usedGB := float32(diskUsage.Used) / gb
	usagePercentage := float32(diskUsage.UsedPercent)
	return stats.SingleStat{ totalGB, usedGB, availableGB, usagePercentage }
}

func getCPUStat() stats.MultipleStat {
	perCpu := true
	usagePercentagePerCore, _ := cpu.CPUPercent(time.Second, perCpu)
	return buildCPUMultipleStat(usagePercentagePerCore)
}

func getRAMStat() stats.SingleStat {
	vm, _ := mem.VirtualMemory()
	return buildRamStat(*vm)
}

func getDiskStat() stats.SingleStat {
	path := "/"
	diskUsage, _ := disk.DiskUsage(path)
	return buildDiskStat(*diskUsage)
}

func getAllStats(cpu stats.MultipleStat, ram stats.SingleStat, disk stats.SingleStat) stats.AllStat {
	return stats.AllStat{ cpu, ram, disk }
}

func main() {

	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	router.GET("/all", func(context *gin.Context) {
		allStats := getAllStats(getCPUStat(), getRAMStat(), getDiskStat())
		context.JSON(200, allStats)
	})

	router.GET("/cpu", func(context *gin.Context) {
		cpu := getCPUStat()
		context.JSON(200, cpu)
	})

	router.GET("/ram", func(context *gin.Context) {
		ram := getRAMStat()
		context.JSON(200, ram)
	})

	router.GET("/disk", func(context *gin.Context) {
		disk := getDiskStat()
		context.JSON(200, disk)
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
