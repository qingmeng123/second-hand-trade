/*******
* @Author:qingmeng
* @Description:
* @File:goods
* @Date2022/2/18
 */

package service

import (
	"database/sql"
	"second-hand-trade/dao"
	"second-hand-trade/model"
)

type GoodsService struct {
}

func (s *GoodsService) AddGoods(goods model.Goods) error {
	d := dao.GoodsDao{}
	return d.AddGoods(goods)
}

func (s *GoodsService) ViewGoods() ([]model.Goods, error) {
	d := dao.GoodsDao{}
	return d.ViewGoods()
}

func (s *GoodsService) SelectGoodsesByName(name string) ([]model.Goods, error) {
	d := dao.GoodsDao{}
	return d.SelectGoodsesByName(name)
}

func (s *GoodsService) SelectGoodsByGoodsId(gid int) (model.Goods, error) {
	d := dao.GoodsDao{}
	return d.SelectGoodsByGoodsId(gid)
}

func (s *GoodsService) UpdateGoods(goods model.Goods) error {
	d := dao.GoodsDao{}
	return d.UpdateGoods(goods)
}

func (s *GoodsService) AddTurnover(goodsId int, number int) error {
	d := dao.GoodsDao{}
	return d.AddTurnover(goodsId, number)
}

func (s *GoodsService) SelectGoodsesBySortId(sortId int) ([]model.Goods, error) {
	d := dao.GoodsDao{}
	return d.SelectGoodsesBySortId(sortId)
}

func (s *GoodsService) SelectGoodsesByGoodsNameSortId(name string, sortId int) ([]model.Goods, error) {
	d := dao.GoodsDao{}
	return d.SelectGoodsesByGoodsNameSortId(name, sortId)
}

func (s *GoodsService) IsExistGoodsId(goodId int) (bool, error) {
	d := dao.GoodsDao{}
	_, err := d.SelectGoodsByGoodsId(goodId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
