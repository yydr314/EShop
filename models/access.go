package models

type Access struct {
	Id          int
	ModuleName  string //模組名稱
	ActionName  string //操作名稱
	Type        int    //節點類型： 1. 模塊 2. 菜單 3. 操作
	Url         string //路由跳轉地址
	ModuleId    int    //此module_id與當前模型_id的自關聯  module_id=0代表模塊
	Sort        int
	Description string
	Status      int
	AddTime     int
	AccessItem  []Access `gorm:"foreignKey:ModuleId;references:Id"`
	Checked     bool     `gorm:"-"`		//	用 "-" 方法代表忽略這個字段
}

func (Access) TableName() string {
	return "access"
}
