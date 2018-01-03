package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type HydraOauth2Code struct {
	Id           int       `orm:"column(signature);pk"`
	RequestId    string    `orm:"column(request_id)"`
	RequestedAt  time.Time `orm:"column(requested_at);type(timestamp without time zone)"`
	ClientId     string    `orm:"column(client_id)"`
	Scope        string    `orm:"column(scope)"`
	GrantedScope string    `orm:"column(granted_scope)"`
	FormData     string    `orm:"column(form_data)"`
	SessionData  string    `orm:"column(session_data)"`
}

func (t *HydraOauth2Code) TableName() string {
	return "hydra_oauth2_code"
}

func init() {
	orm.RegisterModel(new(HydraOauth2Code))
}

// AddHydraOauth2Code insert a new HydraOauth2Code into database and returns
// last inserted Id on success.
func AddHydraOauth2Code(m *HydraOauth2Code) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHydraOauth2CodeById retrieves HydraOauth2Code by Id. Returns error if
// Id doesn't exist
func GetHydraOauth2CodeById(id int) (v *HydraOauth2Code, err error) {
	o := orm.NewOrm()
	v = &HydraOauth2Code{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHydraOauth2Code retrieves all HydraOauth2Code matches certain condition. Returns empty list if
// no records exist
func GetAllHydraOauth2Code(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HydraOauth2Code))
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

	var l []HydraOauth2Code
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

// UpdateHydraOauth2Code updates HydraOauth2Code by Id and returns error if
// the record to be updated doesn't exist
func UpdateHydraOauth2CodeById(m *HydraOauth2Code) (err error) {
	o := orm.NewOrm()
	v := HydraOauth2Code{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHydraOauth2Code deletes HydraOauth2Code by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHydraOauth2Code(id int) (err error) {
	o := orm.NewOrm()
	v := HydraOauth2Code{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HydraOauth2Code{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
