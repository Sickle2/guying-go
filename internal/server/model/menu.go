package model

/**
 * @program: guying-go
 * @description: 菜单
 * @author: sickle
 * @create: 2023-07-31 10:55
 **/
type Menu struct {
	MenuId     uint `gorm:"primarykey;autoIncrement"`
	MenuName   string
	ParentName string
	ParentId   uint
	Path       string
	sort       int
	Component  string
	Query      string
	IsFrame    int
	IsCache    int
	MenuType   int
	Visible    int
	Status     int
	Perms      string
	Icon       string
}
