package main

import (
	"fmt"
	"log"
	"os"
	"torrent/torrent"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: tracker-add {torrent file path}")
	}

	path := os.Args[1]

	//name := ""
	//
	//if len(os.Args) > 2 {
	//	name = os.Args[1]
	//}
	//
	//torrent.AddTrackers(path, name)

	info, err := torrent.NewTorrent(path)

	if err != nil {
		log.Printf("parse torrent failed %v", err)
		os.Exit(1)
	}

	fmt.Println(info.Data.Info.PieceLength)
	fmt.Println(info.Data.Info.Name)

	//fmt.Println(info.Data.Info.Files)

	for _, file := range info.Data.Info.Files {
		fmt.Print("length:")
		fmt.Println(file.Length)
		fmt.Print("md5sum:")
		fmt.Println(file.Md5sum)
		fmt.Print("path:")
		fmt.Println(file.Path)
	}
}
