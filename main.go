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
		log.Fatalf("torrent not available: %v", err)
	}

	trackerList, err := tracker.GetTrackerList()
	if err != nil {
		log.Fatalf("get tracker list failed: %v", err)
	}

	torrentFile.Data.AnnounceList = append(torrentFile.Data.AnnounceList, trackerList...)
	log.Println(torrentFile.Data.AnnounceList)

	err = torrent.WriteToFile(torrentFile.Data, GetDstFileName(torrentFile.Data.Info.Name))

	if err != nil {
		log.Fatal(err)
	}
}

func GetDstFileName(name string) string {
	if len(os.Args) > 2 {
		return os.Args[1]
	} else {
		return name + ".tracked" + ".torrent"
	}
}
