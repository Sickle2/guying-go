package model

/**
 * @program: guying-go
 * @description:
 * @author: sickle
 * @create: 2023-07-28 16:32
 **/
type AdminUser struct {
	UserID      uint `gorm:"primarykey;autoIncrement"`
	UserName    string
	NickName    string
	UserType    int
	Email       string
	PhoneNumber string
	Sex         string
	Avatar      string
	Password    string
	Status      int
	DeptId      string
	PostId      string
	OpenId      string
	UnionId     string
	BaseModel
}
