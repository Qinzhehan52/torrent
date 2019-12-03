package torrent

import (
	"bytes"
	"errors"
	"github.com/jackpal/bencode-go"
	"log"
	"os"
)

// File available as part of the torrent
type File struct {
	Length int    `bencode:"length"`
	Md5sum string `bencode:"md5sum"`
	Path   string `bencode:"path"`
}

// Data about the download itself
type MetaInfoData struct {
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
	//	Private     int    `bencode:"private"`
	//	Length      int    `bencode:"length"`
	//	Md5sum      string `bencode:"md5sum"`
	Files []File `bencode:"files"`
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
