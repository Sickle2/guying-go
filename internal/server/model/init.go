package model

import (
	internal "guying-go/internal/db"
)

var Conn = internal.Conn

func init() {
	InitTable()
}

func InitTable() {
	Conn.AutoMigrate(&AdminUser{})
	Conn.AutoMigrate(&Department{})
	Conn.AutoMigrate(&JobPosition{})
	Conn.AutoMigrate(&Menu{})
	Conn.AutoMigrate(&RoleDept{})
	Conn.AutoMigrate(&Role{})
	Conn.AutoMigrate(&RoleMenu{})
	Conn.AutoMigrate(&UserPost{})
	Conn.AutoMigrate(&UserRole{})
}
