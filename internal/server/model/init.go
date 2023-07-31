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
}
