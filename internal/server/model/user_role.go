package model

/**
 * @program: guying-go
 * @description: 用户权限表
 * @author: sickle
 * @create: 2023-07-31 14:17
 **/
type UserRole struct {
	UserId uint `gorm:"primarykey;autoIncrement"`
	RoleId uint
	BaseModel
}
