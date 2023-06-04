/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package dao

import (
	"second-hand-trade/model"
)

type UserDao struct {
}

func (dao *UserDao) InsertUser(user model.User) error {
	result := GormDB.Select("username", "password", "gender", "name", "salt").Create(&user)
	return result.Error
}

func (dao *UserDao) SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}
	result := GormDB.Where("username=?", username).First(&user)
	return user, result.Error
}

func (dao *UserDao) SelectBasicUserByUsername(username string) (model.User, error) {
	user := model.User{}
	result := GormDB.Select("uid", "password", "group_id", "store_id").Where("username=?", username).First(&user)
	return user, result.Error
}

func (dao *UserDao) UpdatePassword(username, newPassword string) error {
	result := GormDB.Model(&model.User{}).Where("username=?", username).Update("password", newPassword)
	return result.Error
}

// UpdatePhone 更新用户电话
func (dao *UserDao) UpdatePhone(username string, phone string) error {

	result := GormDB.Model(&model.User{}).Where("username=?", username).Update("phone", phone)
	return result.Error
}

// UpdateName 更新昵称
func (dao *UserDao) UpdateName(username string, name string) error {
	_, err := DB.Exec("update second_hand_trade.userinfo set name=? where username=?", name, username)
	return err
}

// UpdateGender 更新性别
func (dao *UserDao) UpdateGender(username string, gender string) error {
	_, err := DB.Exec("update second_hand_trade.userinfo set gender=? where username=?", gender, username)
	return err
}

func (dao *UserDao) UpdateAddressId(username string, id int) error {
	_, err := DB.Exec("update second_hand_trade.userinfo set address_id=? where address_id=?", username, id)
	return err
}

func (dao *UserDao) UpdateGroupId(username string, i int) error {
	_, err := DB.Exec("update second_hand_trade.userinfo set group_id=? where username=?", i, username)
	return err
}

// AddStoreUser 商家入铺
func (dao *UserDao) AddStoreUser(username string, sid int) error {
	_, err := DB.Exec("update second_hand_trade.userinfo set store_id=? where username=?", sid, username)
	return err
}

// UpdateMoney 更新余额
func (dao *UserDao) UpdateMoney(username string, money float32) error {
	_, err := DB.Exec("update second_hand_trade.userinfo set money=? where username=?", money, username)
	return err
}

// 通过电话查找用户
func (dao *UserDao) SelectUserByPhone(phone string) (model.User, error) {
	user := model.User{}
	row := DB.QueryRow("select uid,username,gender,name,money,address_id,group_id,store_id,salt from second_hand_trade.userinfo where phone=?", phone)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.Uid, &user.Username, &user.Gender, &user.Name, &user.Money, &user.AddressId, &user.GroupId, &user.StoreId, &user.Salt)
	if err != nil {
		return user, err
	}
	user.Phone = phone
	return user, err
}

// 验证码注册
func (dao *UserDao) RegisterBySms(user model.User) error {
	_, err := DB.Exec("insert into second_hand_trade.userinfo(username, phone)values(?,?)", user.Username, user.Phone)
	return err
}
