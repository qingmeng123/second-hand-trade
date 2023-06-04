/*******
* @Author:qingmeng
* @Description:
* @File:address
* @Date2022/2/18
 */

package dao

import "second-hand-trade/model"

type AddressDao struct {
}

func (dao *AddressDao) InsertAddress(info model.AddressInfo) error {
	_, err := DB.Exec("insert into second_hand_trade.address_info( uid, name, phone, address) VALUES(?,?,?,?) ", info.Uid, info.Name, info.Phone, info.Address)
	return err
}

// SelectAddressesByUid 通过uid遍历address
func (dao *AddressDao) SelectAddressesByUid(uid int) ([]model.AddressInfo, error) {
	var addressInfos []model.AddressInfo
	rows, err := DB.Query("select address_id,name,phone,address from second_hand_trade.address_info where uid=?", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		addressInfo := model.AddressInfo{}
		err = rows.Scan(&addressInfo.AddressId, &addressInfo.Name, &addressInfo.Phone, &addressInfo.Address)
		if err != nil {
			return nil, err
		}
		addressInfo.Uid = uid
		addressInfos = append(addressInfos, addressInfo)
	}
	return addressInfos, err
}

// GetAddressInfoByAddressId 通过addressId获取address
func (dao *AddressDao) GetAddressInfoByAddressId(id int) (model.AddressInfo, error) {
	var addressInfo model.AddressInfo
	addressInfo.AddressId = id
	row := DB.QueryRow("select uid,name,phone,address from second_hand_trade.address_info where address_id=?", id)
	if row.Err() != nil {
		return addressInfo, row.Err()
	}
	err := row.Scan(&addressInfo.Uid, &addressInfo.Name, &addressInfo.Phone, &addressInfo.Address)
	return addressInfo, err
}

// UpdateAddress 更新地址信息
func (dao *AddressDao) UpdateAddress(info model.AddressInfo) error {
	_, err := DB.Exec("update second_hand_trade.address_info set name=? , phone=? , address=? where address_id=?", info.Name, info.Phone, info.Address, info.AddressId)
	return err
}
