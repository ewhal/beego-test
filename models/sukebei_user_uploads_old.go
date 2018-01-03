package models

type SukebeiUserUploadsOld struct {
	Username  string `orm:"column(username);null"`
	TorrentId int    `orm:"column(torrent_id);null"`
}
