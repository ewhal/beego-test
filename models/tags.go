package models

type Tags struct {
	TorrentId int     `orm:"column(torrent_id);null"`
	UserId    int     `orm:"column(user_id);null"`
	Tag       string  `orm:"column(tag);null"`
	Type      string  `orm:"column(type);null"`
	Weight    float64 `orm:"column(weight);null"`
	Accepted  bool    `orm:"column(accepted);null"`
}
