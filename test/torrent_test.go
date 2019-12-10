package test

import (
	"os"
	"testing"
	"torrent/torrent"
)

func Test_NewTorrent(t *testing.T) {
	info, err := torrent.NewTorrent("/Users/qinzhehan/torrent/demo.torrent")

	if err != nil {
		t.Errorf("compute torrent hash %v", err)
	}

	files := info.Data.Info.Files

	t.Log(info.Data.CreationDate)
	t.Log(files[0].Path)
}

func Test_ComputeTorrentHash(t *testing.T) {
	hash, err := torrent.ComputeTorrentHash("../demo.torrent")

	if err != nil {
		t.Errorf("compute torrent hash %v", err)
	}

	t.Logf("%x", hash)
}

func Test_AddTrackers(t *testing.T) {
	_ = os.Chdir("../")
	torrent.AddTrackers("demo.torrent", "")
}
