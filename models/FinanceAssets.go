package models

// 金融资产表
type FinanceAssets struct {
	Base
	Account  *Account `orm:"rel(fk)" json:"account"`                                              //设置一对多关系
	Type     int      `orm:"column(type);description(记录类型:0:卖出,1:买入)" json:"type"` //资产类型
	Code     string   `orm:"column(code)" json:"code" `                            //代码
	CodeType int      `orm:"column(code_type);description(记录类型:0:基金，1:股票，3:黄金)" json:"code_type"`
	Count    int64    `orm:"column(count)" json:"count"` //数量
	Price    float64  `orm:"column(price);digits(12);decimals(4)" json:"price"`
	Journal  string   `orm:"column(journal);type(text);null" json:"journal"`
}

// TableName 获取对应数据库表名.
func (f *FinanceAssets) TableName() string {
	return "finance_assets"
}

// TableEngine 获取数据使用的引擎.
func (f *FinanceAssets) TableEngine() string {
	return "INNODB"
}


func NewFinanceAssets() *FinanceAssets {
	return &FinanceAssets{}
}
