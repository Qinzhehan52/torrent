package response

type Response struct {
	errno  int
	errmsg string
	result struct{}
}

type PieceInfo struct {
	PieceLength int
}

type TorrentInfo struct {
	Name         string
	MagnetUrl    string
	Encoding     string
	CreateTime   string
	CreateBy     string
	Hash         string
	Note         string
	PieceInfo    PieceInfo
	AnnounceList []string
	FileNameList []string
}
