/*******
* @Author:qingmeng
* @Description:
* @File:address
* @Date2022/2/18
 */

package model

type AddressInfo struct {
	AddressId int    `json:"address_id"`
	Uid       int    `json:"uid"`
	Name      string `json:"name"`    //收货人
	Phone     string `json:"phone"`   //收货电话
	Address   string `json:"address"` //收货地址
}

func (AddressInfo) TableName() string {
	return "address_info"
}
