package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"torrent/models"
)

func UploadAPI(c *gin.Context) {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errno":  -1,
			"errmsg": "上传文件错误" + err.Error(),
			"data":   struct{}{},
		})
		return
	}

	filename, err := filepath.Abs(file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errno":  -1,
			"errmsg": "上传文件错误" + err.Error(),
			"data":   struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data":   models.File{Name: file.Filename, Path: filename, Size: strconv.Itoa(int(file.Size))},
	})
}
