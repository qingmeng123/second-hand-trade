/*******
* @Author:qingmeng
* @Description:
* @File:address
* @Date2022/2/18
 */

package service

import (
	"second-hand-trade/dao"
	"second-hand-trade/model"
)

type AddressService struct {
}

// IsExistAddress 地址是否存在
func (s *AddressService) IsExistAddress(addressInfo model.AddressInfo) (bool, error) {
	d := dao.AddressDao{}
	addresses, err := d.SelectAddressesByUid(addressInfo.Uid)
	if err != nil {
		return false, err
	}
	for _, address := range addresses {
		if address.Address == addressInfo.Address && address.Name == addressInfo.Name && address.Phone == addressInfo.Phone {
			return true, nil
		}
	}
	return false, err
}

// SelectAddressesByUid 通过uid遍历address
func (s *AddressService) SelectAddressesByUid(uid int) ([]model.AddressInfo, error) {
	d := dao.AddressDao{}
	addresses, err := d.SelectAddressesByUid(uid)
	return addresses, err
}

// InsertAddress 添加地址
func (s *AddressService) InsertAddress(addressInfo model.AddressInfo) error {
	d := dao.AddressDao{}
	err := d.InsertAddress(addressInfo)
	return err
}

// GetAddressInfoByAddressId 通过addressId获取address
func (s *AddressService) GetAddressInfoByAddressId(id int) (model.AddressInfo, error) {
	d := dao.AddressDao{}
	return d.GetAddressInfoByAddressId(id)

}

func (s *AddressService) UpdateAddress(info model.AddressInfo) error {
	d := dao.AddressDao{}
	return d.UpdateAddress(info)
}
