package models

import "time"

type Base struct {
	Id         int64     `orm:"pk;auto;unique;column(id)" json:"id"`
	UpdateTime time.Time `orm:"type(datetime);column(update_time);null;auto_now" json:"update_time"`
	CreateTime time.Time `orm:"type(datetime);column(create_time);auto_now_add" json:"create_time"`
}

type Response struct {
	Total int64       `json:"total"`
	Rows  interface{} `json:"rows"`
}