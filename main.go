package main

import (
	"log"
	"os"
	"torrent/torrent"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: tracker-add {torrent file path}")
	}

	path := os.Args[1]

	name := ""

	if len(os.Args) > 2 {
		name = os.Args[1]
	}

	torrent.AddTrackers(path, name)
}
