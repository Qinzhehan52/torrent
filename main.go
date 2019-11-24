package main

import (
	"log"
	"os"
	"torrent/torrent"
	"torrent/tracker"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: tracker-add {torrent file path}")
	}

	path := os.Args[1]

	torrentFile, err := torrent.NewTorrent(path)

	if err != nil {
		log.Fatal("torrent not available")
	}

	trackerList, err := tracker.GetTrackerList()
	if err != nil {
		log.Fatalf("get tracker list failed: %v", err)
	}

	torrentFile.Data.AnnounceList = append(torrentFile.Data.AnnounceList, trackerList)
	log.Println(torrentFile.Data.AnnounceList)
}
