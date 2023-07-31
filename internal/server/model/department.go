package model

/**
 * @program: guying-go
 * @description: 部门
 * @author: sickle
 * @create: 2023-07-31 10:52
 **/
type Department struct {
	DeptId      uint `gorm:"primarykey;autoIncrement"`
	ParentId    uint
	DeptName    string
	Description string
	Sort        int
	Leader      string
	LeaderPhoto string
	LeaderEmail string
	Status      int
	BaseModel
}
