package models

type IncomeTag struct {
	Id   int64  `orm:"pk;auto;unique;column(id)" json:"id"`
	Name string `orm:"unique;size(500)" json:"name"`
}

// TableName 获取对应数据库表名.
func (t *IncomeTag) TableName() string {
	return "income_tag"
}

// TableEngine 获取数据使用的引擎.
func (t *IncomeTag) TableEngine() string {
	return "INNODB"
}
func (t *IncomeTag) TableNameWithPrefix(prefix string) string {
	return prefix + t.TableName()
}

type Income struct {
	Base
	Tag     *IncomeTag `orm:"rel(fk)" json:"tag"`
	Journal string     `orm:"column(journal);type(text);null" json:"journal"`
	Income  float64    `orm:"column(income);digits(12);decimals(4);description(金额)" json:"income"`
}

func (t *Income) TableName() string {
	return "income"
}

// TableEngine 获取数据使用的引擎.
func (t *Income) TableEngine() string {
	return "INNODB"
}
func (t *Income) TableNameWithPrefix(prefix string) string {
	return prefix + t.TableName()
}
