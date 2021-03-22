package model

import (
	"encoding/base64"
	"fmt"
	"gin_blog/utils/errmsg"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;default:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

//查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username=?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//添加用户
func CreateUser(data *User) int {
	// data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //返回500
	}
	return errmsg.SUCCESS

}

//用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

//使用钩子加密
func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}

//密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)

	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("===============")
	fmt.Println(HashPw)
	Fpw := base64.StdEncoding.EncodeToString(HashPw)
	fmt.Println("-----------------")
	fmt.Println(Fpw)
	return Fpw
}

//删用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//更新用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?").Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//登入验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username=?", username).First(&user)
	fmt.Println(ScryptPw(password))
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORF_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
