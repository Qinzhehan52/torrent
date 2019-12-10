package router

import (
	"github.com/gin-gonic/gin"
	"torrent/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/file/upload", apis.UploadAPI)
	router.GET("/api/torrent/info", apis.ParseTorrentAPI)

	return router
}
