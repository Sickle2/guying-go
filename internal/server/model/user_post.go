package model

/**
 * @program: guying-go
 * @description: 用户职位关系表
 * @author: sickle
 * @create: 2023-07-31 14:18
 **/
type UserPost struct {
	UserId uint `gorm:"primarykey;autoIncrement"`
	PostId uint
	BaseModel
}
