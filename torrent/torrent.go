package torrent

import (
	"bytes"
	"crypto/sha1"
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
	Private     int    `bencode:"private"`
	Length      int    `bencode:"length"`
	Md5sum      string `bencode:"md5sum"`
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
	Hash []byte
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

	log.Print("Computing torrent info hash")
	infoHash, err := computeInfoHash(path)
	if err != nil {
		return Torrent{}, errors.New("Failed to compute info hash: " + err.Error())
	}

	log.Printf("Announce URL: %s", info.Announce)
	log.Printf("Hash: %x", infoHash)

	return Torrent{Path: path, Data: info, Hash: infoHash}, nil
}

func computeInfoHash(torrentPath string) ([]byte, error) {

	file, err := os.Open(torrentPath)
	if err != nil {
		return nil, errors.New("Failed to open torrent: " + err.Error())
	}
	defer file.Close()

	data, err := bencode.Decode(file)
	if err != nil {
		return nil, errors.New("Failed to decode torrent file: " + err.Error())
	}

	torrentDict, ok := data.(map[string]interface{})
	if !ok {
		return nil, errors.New("(Torrent file is not a dictionary)")
	}

	infoBuffer := bytes.Buffer{}
	err = bencode.Marshal(&infoBuffer, torrentDict["info"])
	if err != nil {
		return nil, errors.New("Failed to encode info dict: " + err.Error())
	}

	hash := sha1.New()
	hash.Write(infoBuffer.Bytes())
	return hash.Sum(nil), nil
}
