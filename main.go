package main

import (
	"./stats"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	router.GET("/cpu", func(context *gin.Context) {
		cpu := stats.GetCPUStat()
		context.IndentedJSON(200, cpu)
	})

	router.GET("/ram", func(context *gin.Context) {
		ram := stats.GetRAMStat()
		context.IndentedJSON(200, ram)
	})

	router.GET("/disk", func(context *gin.Context) {
		disk := stats.GetDiskStat()
		context.IndentedJSON(200, disk)
	})

	router.GET("/all", func(context *gin.Context) {
		allStats := stats.GetAllStats(stats.GetCPUStat(), stats.GetRAMStat(), stats.GetDiskStat())
		context.IndentedJSON(200, allStats)
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
