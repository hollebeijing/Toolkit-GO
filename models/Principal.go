package models

//金鹅转入表
type Principal struct {
	Base
	Account *Account `orm:"rel(fk)" json:"account"` //设置一对多关系
	Type    int      `orm:"column(type);" json:"type"`
	Goose   float64  `orm:"column(goose);digits(12);decimals(4)" json:"goose"`
	Journal string   `orm:"column(journal);type(text);null" json:"journal"`
}

// TableName 获取对应数据库表名.
func (p *Principal) TableName() string {
	return "principal"
}

// TableEngine 获取数据使用的引擎.
func (p *Principal) TableEngine() string {
	return "INNODB"
}

func NewPrincipalIncomet() *Principal {
	return &Principal{}
}

func (p *Principal) GetABSGoose() float64 {
	if p.Type == 2 {
		return 0 - p.Goose
	} else {
		return p.Goose
	}
}
