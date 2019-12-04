package test

import (
	"testing"
	"torrent/torrent"
)

func Test_ComputeTorrentHash(t *testing.T) {
	hash, err := torrent.ComputeTorrentHash("../demo.torrent")

	if err != nil {
		t.Errorf("compute torrent hash %v", err)
	}

	t.Logf("%x", hash)
}
