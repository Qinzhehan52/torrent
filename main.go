package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"torrent/response"
	"torrent/torrent"
)

func main() {
	r := gin.Default()
	r.GET("/api/torrent/msg", func(c *gin.Context) {
		torrentFile, err := torrent.NewTorrent("demo.torrent")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errno":  -1,
				"errmsg": "服务器内部错误",
				"data":   struct{}{},
			})
		} else {
			createDate := time.Unix(int64(torrentFile.Data.CreationDate), 0)

			announceList := make([]string, 0)

			for _, sli := range torrentFile.Data.AnnounceList {
				announceList = append(announceList, sli...)
			}

			fileNameList := make([]string, 0)
			for _, file := range torrentFile.Data.Info.Files {
				fileNameList = append(file.Path)
			}

			hash, _ := torrent.ComputeTorrentHash("demo.torrent")

			torrentInfo := response.TorrentInfo{
				Name:         torrentFile.Data.Info.Name,
				MagnetUrl:    "magnet:?xt=urn:btih:" + strings.ToUpper(fmt.Sprintf("%x", hash)),
				Encoding:     torrentFile.Data.Encoding,
				CreateTime:   createDate.String(),
				CreateBy:     torrentFile.Data.CreatedBy,
				Hash:         fmt.Sprintf("%x", hash),
				PieceInfo:    response.PieceInfo{PieceLength: torrentFile.Data.Info.PieceLength},
				AnnounceList: announceList,
				FileNameList: fileNameList,
			}

			c.JSON(http.StatusOK, gin.H{
				"errno":  0,
				"errmsg": "",
				"data":   torrentInfo,
			})
		}
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
