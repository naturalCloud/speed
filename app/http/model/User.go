package model

import app "speed/bootstrap"

type Users struct {
	ID          int    `gorm:"primary_key;column:Id;type:int(10) unsigned;not null"` //	主键
	Openid      string `gorm:"unique;column:Openid;type:varchar(55);not null"`       //	微信openid
	NickName    string `gorm:"column:NickName;type:varchar(64);not null"`            //	昵称
	HeadImgURL  string `gorm:"column:HeadImgUrl;type:varchar(300);not null"`         //	头像地址
	IsSubscribe int    `gorm:"column:IsSubscribe;type:tinyint(4);not null"`          //	是否关注0:未,1:关注
	Sex         int    `gorm:"column:Sex;type:tinyint(4);not null"`                  //	性别,0:女,1:男
	City        string `gorm:"column:City;type:varchar(55);not null"`                //	城市
	Country     string `gorm:"column:Country;type:varchar(64);not null"`             //	国家
	Province    string `gorm:"column:Province;type:varchar(30);not null"`            //	省份
	Email       string `gorm:"column:Email;type:varchar(64);not null"`               //	电子邮箱
}

func (U Users) TableName() string {
	return "user"
}


func (Users) GetMore() []Users {
	var user []Users
	app.Db.Where("Id >= ?", 1).Find(&user)

	return user
}
