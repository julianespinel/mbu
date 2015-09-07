package main

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"./resources"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	router.GET("/cpu", func(context *gin.Context) {
		perCpu := true
		cpuPercentArray, _ := cpu.CPUPercent(time.Second, perCpu)
		fmt.Println(cpuPercentArray)
		cpuUsage := resources.Stat{ "cpu", 123 }
		fmt.Println(cpuUsage)
		context.JSON(200, cpuUsage)
	})

	router.GET("/memory", func(context *gin.Context) {
		vm, _ := mem.VirtualMemory()
		fmt.Println(vm)
		context.String(200, "memoryUsage")
	})

	router.GET("/disk", func(context *gin.Context) {
		path := "/"
		diskUsage, _ := disk.DiskUsage(path)
		fmt.Println(diskUsage)
		context.String(200, "diskUsage")
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
