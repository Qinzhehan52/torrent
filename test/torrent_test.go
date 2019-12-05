package test

import (
	"testing"
	"torrent/torrent"
)

func Test_NewTorrent(t *testing.T) {
	info, err := torrent.NewTorrent("../demo.torrent")

	if err != nil {
		t.Errorf("compute torrent hash %v", err)
	}

	files := info.Data.Info.Files

	t.Log(files[0].Path)
}

func Test_ComputeTorrentHash(t *testing.T) {
	hash, err := torrent.ComputeTorrentHash("../demo.torrent")

	if err != nil {
		t.Errorf("compute torrent hash %v", err)
	}

	t.Logf("%x", hash)
}
