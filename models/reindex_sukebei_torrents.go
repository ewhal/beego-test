package models

type ReindexSukebeiTorrents struct {
	ReindexTorrentsId int    `orm:"column(reindex_torrents_id)"`
	TorrentId         int    `orm:"column(torrent_id);null"`
	Action            string `orm:"column(action);null"`
}
