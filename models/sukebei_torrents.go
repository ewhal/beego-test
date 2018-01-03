package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SukebeiTorrents struct {
	Id           int       `orm:"column(torrent_id);pk"`
	TorrentName  string    `orm:"column(torrent_name);null"`
	TorrentHash  string    `orm:"column(torrent_hash);null"`
	Category     int       `orm:"column(category);null"`
	SubCategory  int       `orm:"column(sub_category);null"`
	Status       int       `orm:"column(status);null"`
	Date         time.Time `orm:"column(date);type(timestamp with time zone);null"`
	Uploader     int       `orm:"column(uploader);null"`
	Downloads    int       `orm:"column(downloads);null"`
	Stardom      int       `orm:"column(stardom);null"`
	Filesize     int64     `orm:"column(filesize);null"`
	Description  string    `orm:"column(description);null"`
	WebsiteLink  string    `orm:"column(website_link);null"`
	DeletedAt    time.Time `orm:"column(deleted_at);type(timestamp with time zone);null"`
	Trackers     string    `orm:"column(trackers);null"`
	Hidden       bool      `orm:"column(hidden);null"`
	AnidbId      string    `orm:"column(anidb_id);null"`
	Language     string    `orm:"column(language);null"`
	Anidbid      int       `orm:"column(anidbid);null"`
	Vndbid       int       `orm:"column(vndbid);null"`
	Vgmdbid      int       `orm:"column(vgmdbid);null"`
	Dlsite       string    `orm:"column(dlsite);null"`
	Videoquality string    `orm:"column(videoquality);null"`
	Tags         string    `orm:"column(tags);null"`
}

func (t *SukebeiTorrents) TableName() string {
	return "sukebei_torrents"
}

func init() {
	orm.RegisterModel(new(SukebeiTorrents))
}

// AddSukebeiTorrents insert a new SukebeiTorrents into database and returns
// last inserted Id on success.
func AddSukebeiTorrents(m *SukebeiTorrents) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSukebeiTorrentsById retrieves SukebeiTorrents by Id. Returns error if
// Id doesn't exist
func GetSukebeiTorrentsById(id int) (v *SukebeiTorrents, err error) {
	o := orm.NewOrm()
	v = &SukebeiTorrents{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSukebeiTorrents retrieves all SukebeiTorrents matches certain condition. Returns empty list if
// no records exist
func GetAllSukebeiTorrents(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SukebeiTorrents))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SukebeiTorrents
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSukebeiTorrents updates SukebeiTorrents by Id and returns error if
// the record to be updated doesn't exist
func UpdateSukebeiTorrentsById(m *SukebeiTorrents) (err error) {
	o := orm.NewOrm()
	v := SukebeiTorrents{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSukebeiTorrents deletes SukebeiTorrents by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSukebeiTorrents(id int) (err error) {
	o := orm.NewOrm()
	v := SukebeiTorrents{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SukebeiTorrents{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
