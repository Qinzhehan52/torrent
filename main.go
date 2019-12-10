package main

import (
	"torrent/router"
)

func main() {
	r := router.InitRouter()
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
