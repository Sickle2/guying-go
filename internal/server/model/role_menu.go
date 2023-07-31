package model

/**
 * @program: guying-go
 * @description: 权限菜单关系表
 * @author: sickle
 * @create: 2023-07-31 14:50
 **/
type RoleMenu struct {
	RoleId uint `gorm:"primarykey;autoIncrement"`
	MenuId uint
	BaseModel
}
