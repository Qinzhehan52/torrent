package response

type Response struct {
	Errno  int      `json:"errno"`
	ErrMsg string   `json:"err_msg"`
	Result struct{} `json:"result"`
}

type PieceInfo struct {
	PieceLength int `json:"piece_length"`
}

type TorrentInfo struct {
	Name         string    `json:"name"`
	MagnetUrl    string    `json:"magnet_url"`
	Encoding     string    `json:"encoding"`
	CreateTime   string    `json:"create_time"`
	CreateBy     string    `json:"create_by"`
	Hash         string    `json:"hash"`
	Note         string    `json:"note"`
	PieceInfo    PieceInfo `json:"piece_info"`
	AnnounceList []string  `json:"announce_list"`
	FileNameList []string  `json:"file_name_list"`
}
