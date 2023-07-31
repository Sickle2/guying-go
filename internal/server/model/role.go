package model

/**
 * @program: guying-go
 * @description:
 * @author: sickle
 * @create: 2023-07-31 14:14
 **/
type Role struct {
	RoleId   uint `gorm:"primarykey;autoIncrement"`
	RoleName string
	roleKey  string
	sort     int
	status   int
}
