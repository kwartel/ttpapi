package ttpapi

type TorrentGroup struct {
	Group   GroupType     `json:"group"`
	Torrent []TorrentType `json:"torrents"`
}
