/*******
* @Author:qingmeng
* @Description:
* @File:store
* @Date2022/2/19
 */

package dao

import "second-hand-trade/model"

type StoreDao struct {
}

// CreateStore 创建store
func (d *StoreDao) CreateStore(store model.Store) error {
	_, err := DB.Exec("insert into second_hand_trade.store ( store_name, notice) values (?,?);", store.StoreName, store.Notice)
	return err
}

// SelectStoreByStoreName 根据storeName选择store
func (d *StoreDao) SelectStoreByStoreName(name string) (model.Store, error) {
	var store model.Store
	row := DB.QueryRow("select * from second_hand_trade.store where store_name=?", name)
	if row.Err() != nil {
		return store, row.Err()
	}
	err := row.Scan(&store.StoreId, &store.StoreName, &store.Notice, &store.StoreMoney)
	if err != nil {
		return model.Store{}, err
	}
	return store, err
}

// UpdateStore 更新store信息
func (d *StoreDao) UpdateStore(store model.Store) error {
	_, err := DB.Exec("update second_hand_trade.store set store_name=?,notice=? where store_id=?", store.StoreName, store.Notice, store.StoreId)
	return err
}

// SelectStoreByStoreId 根据storeId选择store
func (d *StoreDao) SelectStoreByStoreId(sid int) (model.Store, error) {
	var store model.Store
	row := DB.QueryRow("select * from second_hand_trade.store where store_id=?", sid)
	if row.Err() != nil {
		return store, row.Err()
	}
	err := row.Scan(&store.StoreId, &store.StoreName, &store.Notice, &store.StoreMoney)
	if err != nil {
		return model.Store{}, err
	}
	return store, err
}

func (d *StoreDao) AddStoreMoney(storeId int, money float64) error {
	_, err := DB.Exec("update second_hand_trade.store set store_money =store_money+?  where store_id=?;", money, storeId)
	return err
}
