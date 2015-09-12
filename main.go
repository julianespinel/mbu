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

func main() {

	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	router.GET("/cpu", func(context *gin.Context) {
		perCpu := true
		cpuPercentArray, _ := cpu.CPUPercent(time.Second, perCpu)

		averageUsagePercentage := getAverage(cpuPercentArray);
		cpu := stats.MultipleStat{ averageUsagePercentage, cpuPercentArray }
		context.JSON(200, cpu)
	})

	router.GET("/ram", func(context *gin.Context) {
		vm, _ := mem.VirtualMemory()

		gb := float32(1000000000)
		totalGB := float32(vm.Total) / gb
		availableGB := float32(vm.Available) / gb
		usedGB := (totalGB - availableGB)
		usagePercentage := float32(vm.UsedPercent)
		ram := stats.SingleStat{ totalGB, usedGB, availableGB, usagePercentage }
		context.JSON(200, ram)
	})

	router.GET("/disk", func(context *gin.Context) {
		path := "/"
		diskUsage, _ := disk.DiskUsage(path)

		gb := float32(1000000000)
		totalGB := float32(diskUsage.Total) / gb
		availableGB := float32(diskUsage.Free) / gb
		usedGB := float32(diskUsage.Used) / gb
		usagePercentage := float32(diskUsage.UsedPercent)
		disk := stats.SingleStat{ totalGB, usedGB, availableGB, usagePercentage }
		context.JSON(200, disk)
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
