package model

/**
 * @program: guying-go
 * @description: 职位
 * @author: sickle
 * @create: 2023-07-31 10:47
 **/
type JobPosition struct {
	PostId      uint `gorm:"primarykey;autoIncrement"`
	Title       string
	Description string
	Salary      int
	Location    string
	sort        int
	status      int
	BaseModel
}
