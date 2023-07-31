package model

/**
 * @program: guying-go
 * @description: 权限部门关系表
 * @author: sickle
 * @create: 2023-07-31 14:53
 **/
type RoleDept struct {
	RoleId uint `gorm:"primarykey;autoIncrement"`
	DeptId uint
}
