package main

import (
	"./stats"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	mbu := router.Group("/mbu")
	api := mbu.Group("/api")
	admin := mbu.Group("/admin")

	api.GET("/cpu", func(context *gin.Context) {
		cpu := stats.GetCPUStat()
		context.IndentedJSON(200, cpu)
	})

	api.GET("/ram", func(context *gin.Context) {
		ram := stats.GetRAMStat()
		context.IndentedJSON(200, ram)
	})

	api.GET("/disk", func(context *gin.Context) {
		disk := stats.GetDiskStat()
		context.IndentedJSON(200, disk)
	})

	api.GET("/all", func(context *gin.Context) {
		allStats := stats.GetAllStats(stats.GetCPUStat(), stats.GetRAMStat(), stats.GetDiskStat())
		context.IndentedJSON(200, allStats)
	})

	admin.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
