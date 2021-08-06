package models

//账户表
type Account struct {
	Base
	Name           string  `orm:"column(name);unique;size(500)" json:"name"`
	Tag            string  `orm:"column(tag);size(500)" json:"tag"`
	Cost           float64 `orm:"column(cost);digits(12);decimals(4);description(资本金)" json:"cost"`
	AvaBalance     float64 `orm:"column(ava_balance);digits(12);decimals(4);description(可用资本)" json:"ava_balance"`
	// 所有当前买入资产的 数量 * 单价的总和, 买入所有资产的所有成本
	AlrBalance     float64 `orm:"column(alr_balance);digits(12);decimals(4);description(已用资本，购买资产的本金)" json:"alr_balance"`
	//计算方法：（总估值 - 资本金） / 资本金 * 100%
	RevenueRate    float64 `orm:"column(revenue_rate);digits(12);decimals(4);description(累计年化收益率)" json:"revenue_rate"`
	// 总估值 - 资本金
	Revenue        float64 `orm:"column(revenue);digits(12);decimals(4);description(累计收益)" json:"revenue"`
	//计算方法：加总（当前拥有的资本数量，* 当前价格 ） + 可用资本
	TotalValuation float64 `orm:"column(total_valuation);digits(12);decimals(4);description(总估值)" json:"total_valuation"`
	//Family        *Family                 `orm:"column(family);rel(fk)"` //设置一对多关系
}

// TableName 获取对应数据库表名.
func (t *Account) TableName() string {
	return "account"
}

// TableEngine 获取数据使用的引擎.
func (t *Account) TableEngine() string {
	return "INNODB"
}
func (t *Account) TableNameWithPrefix(prefix string) string {
	return prefix + t.TableName()
}

//// 联合唯一键
//func (a *Account) TableUnique() [][]string {
//	return [][]string{{"team_id", "member_id"}}
//}

// 多字段索引
//func (a *Account) TableIndex() [][]string {
//	return [][]string{
//		[]string{"Id", "Name"},
//	}
//}

//func (t *Account) TableNameWithPrefix() string {
//	return common.GetDatabasePrefix() + t.TableName()
//}
//func (t *Account) QueryTable() orm.QuerySeter {
//	return orm.NewOrm().QueryTable(t.TableNameWithPrefix())
//}

func NewAccount() *Account {
	return &Account{}
}
