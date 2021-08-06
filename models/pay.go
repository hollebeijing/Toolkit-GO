package models

type PayTagX struct {
	Id   int64  `orm:"pk;auto;unique;column(id)" json:"id"`
	Name string `orm:"unique;size(500)" json:"name"`
}

// TableName 获取对应数据库表名.
func (t *PayTagX) TableName() string {
	return "pay_tag_x"
}

// TableEngine 获取数据使用的引擎.
func (t *PayTagX) TableEngine() string {
	return "INNODB"
}
func (t *PayTagX) TableNameWithPrefix(prefix string) string {
	return prefix + t.TableName()
}

type PayTagY struct {
	Id   int64  `orm:"pk;auto;unique;column(id)" json:"id"`
	Name string `orm:"unique;size(500)" json:"name"`
}

// TableName 获取对应数据库表名.
func (t *PayTagY) TableName() string {
	return "pay_tag_y"
}

// TableEngine 获取数据使用的引擎.
func (t *PayTagY) TableEngine() string {
	return "INNODB"
}
func (t *PayTagY) TableNameWithPrefix(prefix string) string {
	return prefix + t.TableName()
}

type Pay struct {
	Base
	TagY    *PayTagY `orm:"rel(fk)" json:"tag_y"`
	TagX    *PayTagX `orm:"rel(fk)" json:"tag_x"`
	Journal string   `orm:"column(journal);type(text);null" json:"journal"`
	Pay     float64  `orm:"column(pay);digits(12);decimals(4);description(金额)" json:"pay"`
}

func (t *Pay) TableName() string {
	return "pay"
}

// TableEngine 获取数据使用的引擎.
func (t *Pay) TableEngine() string {
	return "INNODB"
}
func (t *Pay) TableNameWithPrefix(prefix string) string {
	return prefix + t.TableName()
}
