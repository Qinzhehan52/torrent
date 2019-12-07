package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"torrent/response"
	"torrent/torrent"
)

func main() {
	r := gin.Default()
	r.GET("/api/torrent/msg", func(c *gin.Context) {
		torrentFile, err := torrent.NewTorrent("demo.torrent")

		if err != nil {
			c.JSON(500, gin.H{
				"errmsg":  err,
				"message": struct{}{},
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

			torrentInfo := response.TorrentInfo{
				Name:         torrentFile.Data.Info.Name,
				MagnetUrl:    "",
				Encoding:     torrentFile.Data.Encoding,
				CreateTime:   createDate.String(),
				CreateBy:     torrentFile.Data.CreatedBy,
				PieceInfo:    response.PieceInfo{PieceLength: torrentFile.Data.Info.PieceLength},
				AnnounceList: announceList,
				FileNameList: fileNameList,
			}

			c.JSON(200, gin.H{
				"errmsg":  "",
				"message": torrentInfo,
			})
		}
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
