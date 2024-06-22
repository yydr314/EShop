//	檔案名稱要和資料表相同
package models

//	定義資料表的結構
//	名稱要和資料表中的內容相同，如果資料表中長add_time，這裡要改成AddTime
type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

//表示配置操作資料庫的表名稱
func (User) TableName() string {
	return "user"
}