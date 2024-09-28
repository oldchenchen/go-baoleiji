package Model

import (
	. "ZgoMast/application/databases"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `validate:"required,gte=5" gorm:"unique_index;size:255;comment:'账户名'" json:"username"`
	HashPassword string `gorm:"not null;size:255;comment:'密码'" json:"-"`
	Password     string `validate:"required" gorm:"-" json:"password,omitempty"`
	Nickname     string `gorm:"size:255;comment:'昵称' " json:"nickname" sql:"index"`
	Mobile       string `gorm:"index;size:255;comment:'手机';" json:"mobile"`
	Email        string `validate:"omitempty,email" gorm:"index;size:255;comment:'邮箱'" json:"email"`
	Avatar       string `gorm:"size:255;comment:'头像'" json:"avatar"`
	Sex          bool   `gorm:"type:boolean;default:true;comment:'性别'" json:"sex"`
	IP           string `validate:"omitempty,ipv4" gorm:"size:255;comment:'ip地址'" json:"ip"`
	Status       bool   `gorm:"type:boolean;default:true;comment:'状态'" json:"status"`
}

func (User) TableName() string {
	return "user_info"
}
func (user User) Insert() (id uint, err error) {
	result := Orm.Create(&user)
	return user.ID, result.Error
}
func (user User) GetOneByid(id int) {
	Orm.First(&user, id)
}
func (user User) GetOneByAccount(account string) {
	Orm.First(&user, "username = ? or mobile = ? or email", account, account, account)
}
func (user User) GetALL() []User {
	var users []User
	Orm.Find(&users)
	return users
}
func (user User) Chpasswd(HashPasswd string) {
	user.HashPassword = HashPasswd
	Orm.Save(&user)
}

func init() {

}
