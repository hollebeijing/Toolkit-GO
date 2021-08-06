package relational

import (
	"github.com/beego/beego/v2/client/orm"
)

type Query struct {
	Page       int          `json:"page"`
	PageSize   int          `json:"page_size"`
	Conditions []*Condition `json:"conditions"`
	Orderby    []string     `json:"orders"`
	Fields     []string     `json:"fields"`
}

func NewQuery() *Query {
	return &Query{}
}

func (q *Query) SetCondition(c *Condition) {
	q.Conditions = append(q.Conditions, c)
}

func (q *Query) GetCondition() *orm.Condition {
	cond := orm.NewCondition()
	for _, c := range q.Conditions {
		switch c.Operator {
		case "$AND":
			cond = cond.And(c.Field, c.GetValues())
		case "$OR":
			cond = cond.Or(c.Field, c.GetValues())
		case "$AND_NOT":
			cond = cond.AndNot(c.Field, c.GetValues())
		case "$OR_NOT":
			cond = cond.OrNot(c.Field, c.GetValues())
		}
	}
	return cond
}
