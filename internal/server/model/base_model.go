package model

import "time"

/**
 * @program: guying-go
 * @description:
 * @author: sickle
 * @create: 2023-07-28 16:28
 **/
type BaseModel struct {
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	CreateBy   string
	UpdateBy   string
	DeleteBy   int8
}
