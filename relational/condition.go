package relational

import (
	"strings"
)

type Condition struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
	Related  bool        `json:"related"`
}

func NewCondition() *Condition {
	return &Condition{}
}

func (c Condition) GetValues() []interface{} {
	var in []interface{}
	switch c.Value.(type) {
	case nil:
	case float64:
		in = append(in, c.Value.(float64))
	case float32:
		in = append(in, c.Value.(float32))
	case int:
		in = append(in, c.Value.(int))
	case int64:
		in = append(in, c.Value.(int64))
	default:
		for _, v := range strings.Split(c.Value.(string), ",") {
			in = append(in, v)
		}
	}
	return in
}
