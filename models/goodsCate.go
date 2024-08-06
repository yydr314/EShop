package models

type GoodsCate struct {
	Id             int         `json:"id"`
	Title          string      `json:"title"`
	CateImg        string      `json:"cate_img"`
	Link           string      `json:"link"`
	Template       string      `json:"template"`
	Pid            int         `json:"pid"`
	SubTitle       string      `json:"sub_title"`
	Keywords       string      `json:"keywords"`
	Description    string      `json:"description"`
	Sort           int         `json:"sort"`
	Status         int         `json:"status"`
	AddTime        int         `json:"add_time"`
	GoodsCateItems []GoodsCate `gorm:"foreignKey:pid; reference:Id"`
}

func (GoodsCate) TableName() string {
	return "goods_cate"
}
