/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package model

type User struct {
	Uid         int     `json:"uid"`
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	Salt        string  `json:"salt"`
	Gender      bool    `json:"gender"`
	Name        string  `json:"name"`         //昵称
	Phone       string  `json:"phone"`        //账号电话
	Money       float32 `json:"money"`        //余额
	AddressId   int     `json:"address_id"`   //默认收货地址id
	GroupId     int     `json:"group_id"`     //成员组id,1为超级管理员，0为普通用户
	StoreId     int     `json:"store_id"`     //商家的店铺id
	GithubLogin string  `json:"github_login"` //github登陆账号
}

func (User) TableName() string {
	return "userinfo"
}
