package models

import "time"

type CommentsOld struct {
	TorrentId *Torrents `orm:"column(torrent_id);rel(fk)"`
	Username  string    `orm:"column(username)"`
	Content   string    `orm:"column(content)"`
	Date      time.Time `orm:"column(date);type(timestamp without time zone)"`
}
