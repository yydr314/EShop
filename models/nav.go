package models

type Nav struct {
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Url    string `json:"url,omitempty"`
	Status int    `json:"status,omitempty"`
	Sort   int    `json:"sort,omitempty"`
}

func (Nav) TableName() string {
	return "nav"
}
