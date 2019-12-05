package torrent

import (
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/jackpal/bencode-go"
	"log"
	"net/http"
	"os"
	"strings"
	"torrent/tracker"
)

// File available as part of the torrent
type File struct {
	Length int      `bencode:"length"`
	Md5sum string   `bencode:"md5sum"`
	Path   []string `bencode:"path"`
}

// Data about the download itself
type MetaInfoData struct {
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
	Private     int    `bencode:"private"` //不需要
	Length      int    `bencode:"length"`  //单文件使用
	Md5sum      string `bencode:"md5sum"`  //单文件使用
	Files       []File `bencode:"files"`
}

// .torrent file description. Mostly metadata about the torrent
type MetaInfo struct {
	Announce     string       `bencode:"announce"`
	AnnounceList [][]string   `bencode:"announce-list"`
	Info         MetaInfoData `bencode:"info"`
	Encoding     string       `bencode:"encoding"`
	CreationDate int          `bencode:"creation date"`
	CreatedBy    string       `bencode:"created by"`
}

type Torrent struct {
	Path string
	Data MetaInfo
}

// NewTorrent builds a Torrent struct from the given .torrent file path
func NewTorrent(path string) (Torrent, error) {
	log.Printf("Opening %s", path)
	file, err := os.Open(path)
	if err != nil {
		return Torrent{}, errors.New("Failed to open torrent file: " + err.Error())
	}
	defer file.Close()

	log.Print("Decoding torrent file")
	info := MetaInfo{}
	err = bencode.Unmarshal(file, &info)
	if err != nil {
		return Torrent{}, errors.New("Failed to decode torrent file: " + err.Error())
	}

	log.Printf("Announce URL: %s", info.Announce)

	return Torrent{Path: path, Data: info}, nil
}

func ComputeTorrentHash(torrentPath string) (hash []byte, err error) {
	file, err := os.Open(torrentPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := bencode.Decode(file)
	if err != nil {
		return nil, err
	}

	torrentDict, ok := data.(map[string]interface{})
	if !ok {
		return nil, err
	}

	infoBuffer := bytes.Buffer{}
	err = bencode.Marshal(&infoBuffer, torrentDict["info"])
	if err != nil {
		return nil, err
	}

	h := sha1.New()
	h.Write(infoBuffer.Bytes())
	return h.Sum(nil), nil
}

func WriteToFile(info MetaInfo, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var buf bytes.Buffer

	if err = bencode.Marshal(&buf, info); err != nil {
		return err
	}

	_, err = file.Write(buf.Bytes())

	return err
}

func AddTrackers(path string, name string) {
	torrentFile, err := NewTorrent(path)
	if err != nil {
		log.Fatalf("torrent not available: %v", err)
	}

	trackerList, err := tracker.GetTrackerList()
	if err != nil {
		log.Fatalf("get tracker list failed: %v", err)
	}

	torrentFile.Data.AnnounceList = append(torrentFile.Data.AnnounceList, trackerList...)
	log.Println(torrentFile.Data.AnnounceList)

	if len(name) < 1 {
		name = torrentFile.Data.Info.Name
	}
	err = WriteToFile(torrentFile.Data, GetDstFileName(name))

	if err != nil {
		log.Fatal(err)
	}
}

func GetDstFileName(name string) string {
	return name + ".tracked" + ".torrent"
}

func GetSourceInfo(path string) (*http.Response, error) {
	info, err := NewTorrent(path)

	if err != nil {
		log.Fatalf("parse torrent failed: %v", err)
	}

	hash, err := ComputeTorrentHash(path)

	for _, announce := range info.Data.AnnounceList {
		for _, url := range announce {
			strings.Replace(url, "announce", "scrape", -1)
			url += "?info_hash=" + fmt.Sprintf("%x", hash)
			resp, err := http.Get(url)
			if err != nil {
				continue
			}
			return resp, err
		}
	}

	return nil, err
}
