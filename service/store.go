/*******
* @Author:qingmeng
* @Description:
* @File:store
* @Date2022/2/19
 */

package service

import (
	"database/sql"
	"second-hand-trade/dao"
	"second-hand-trade/model"
)

type StoreService struct {
}

func (s *StoreService) CreateStore(store model.Store) error {
	d := dao.StoreDao{}
	return d.CreateStore(store)
}

func (s *StoreService) SelectStoreByStoreName(name string) (model.Store, error) {
	d := dao.StoreDao{}
	return d.SelectStoreByStoreName(name)
}

func (s *StoreService) SelectStoreByStoreId(sid int) (model.Store, error) {
	d := dao.StoreDao{}
	return d.SelectStoreByStoreId(sid)
}

func (s *StoreService) UpdateStore(store model.Store) error {
	d := dao.StoreDao{}
	return d.UpdateStore(store)
}

func (s *StoreService) IsExistStoreName(name string) (bool, error) {
	d := dao.StoreDao{}
	_, err := d.SelectStoreByStoreName(name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *StoreService) IsExistStoreId(sid int) (bool, error) {
	d := dao.StoreDao{}
	_, err := d.SelectStoreByStoreId(sid)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *StoreService) AddStoreMoney(storeId int, money float64) error {
	d := dao.StoreDao{}
	return d.AddStoreMoney(storeId, money)
}
