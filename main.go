package main

import (
	"fmt"
	"strconv"
	"github.com/mbu/stats"
	"github.com/mbu/config"
	"github.com/gin-gonic/gin"
	"github.com/BurntSushi/toml"
)

func main() {

	router := gin.Default()
	mbu := router.Group("/mbu")

	api := mbu.Group("/api")
	{
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
	}

	admin := mbu.Group("/admin")
	{
		admin.GET("/ping", func(context *gin.Context) {
			context.String(200, "pong")
		})
	}

	var config config.Config
	if _, err := toml.DecodeFile("mbu.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	router.Run(":" + strconv.Itoa(config.Server.Port)) // listen and serve on 0.0.0.0:port
}
