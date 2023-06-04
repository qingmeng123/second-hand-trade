/*******
* @Author:qingmeng
* @Description:
* @File:goods
* @Date2022/2/18
 */

package dao

import (
	"second-hand-trade/model"
)

type GoodsDao struct {
}

// AddGoods 添加商品
func (d *GoodsDao) AddGoods(goods model.Goods) error {
	_, err := DB.Exec("insert into second_hand_trade.goods_info(sort_id, store_id,goods_name, picture, price, goods_intro, style,number, shelf_date) VALUES (?,?,?,?,?,?,?,?,?) ", goods.SortId, goods.StoreId, goods.GoodsName, goods.Picture, goods.Price, goods.GoodsIntro, goods.Style, goods.Number, goods.ShelfDate)
	return err
}

// ViewGoods 查看所有商品
func (d *GoodsDao) ViewGoods() ([]model.Goods, error) {
	var goodses []model.Goods

	result := GormDB.Preload("Store").Find(&goodses)
	return goodses, result.Error
}

// SelectGoodsesByName 根据商品关键字选择查看商品
func (d *GoodsDao) SelectGoodsesByName(name string) ([]model.Goods, error) {
	var goodses []model.Goods
	name = "%" + name + "%"
	rows, err := DB.Query("select * from second_hand_trade.goods_info where goods_name like ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		goods := model.Goods{}
		err = rows.Scan(&goods.GoodsId, &goods.SortId, &goods.StoreId, &goods.GoodsName, &goods.Picture, &goods.Price,
			&goods.GoodsIntro, &goods.Turnover, &goods.Style, &goods.Number, &goods.ShelfDate)
		if err != nil {
			return nil, err
		}
		goodses = append(goodses, goods)
	}
	return goodses, err
}

// SelectGoodsByGoodsId 根据goodsId选择商品
func (d *GoodsDao) SelectGoodsByGoodsId(gid int) (model.Goods, error) {
	var goods model.Goods
	row := DB.QueryRow("select * from second_hand_trade.goods_info where goods_id=?", gid)
	if row.Err() != nil {
		return goods, row.Err()
	}
	err := row.Scan(&goods.GoodsId, &goods.SortId, &goods.StoreId, &goods.GoodsName, &goods.Picture, &goods.Price,
		&goods.GoodsIntro, &goods.Turnover, &goods.Style, &goods.Number, &goods.ShelfDate)
	if err != nil {
		return goods, err
	}
	return goods, nil
}

// UpdateGoods 修改商品
func (d *GoodsDao) UpdateGoods(goods model.Goods) error {
	_, err := DB.Exec("update second_hand_trade.goods_info set sort_id=?,goods_name=?,picture=?,price=?,goods_intro=?,style=?,number=?", goods.SortId, goods.GoodsName, goods.Picture, goods.Price, goods.GoodsIntro, goods.Number)
	return err
}

// AddTurnover 增加成交量
func (d *GoodsDao) AddTurnover(goodsId int, number int) error {
	_, err := DB.Exec("update second_hand_trade.goods_info set number=number+? where goods_id=?", number, goodsId)
	return err
}

// SelectGoodsesBySortId 根据sortId选择商品
func (d *GoodsDao) SelectGoodsesBySortId(sortId int) ([]model.Goods, error) {
	var goodses []model.Goods
	rows, err := DB.Query("select * from second_hand_trade.goods_info where sort_id=?", sortId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		goods := model.Goods{}
		err = rows.Scan(&goods.GoodsId, &goods.SortId, &goods.StoreId, &goods.GoodsName, &goods.Picture, &goods.Price,
			&goods.GoodsIntro, &goods.Turnover, &goods.Style, &goods.Number, &goods.ShelfDate)
		if err != nil {
			return nil, err
		}
		goodses = append(goodses, goods)
	}
	return goodses, err
}

// SelectGoodsesByGoodsNameSortId 根据GoodsNamesortId选择商品
func (d *GoodsDao) SelectGoodsesByGoodsNameSortId(goodsName string, sortId int) ([]model.Goods, error) {
	var goodses []model.Goods
	goodsName = "%" + goodsName + "%"
	rows, err := DB.Query("select * from second_hand_trade.goods_info where sort_id=? and goods_name like ?", sortId, goodsName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		goods := model.Goods{}
		err = rows.Scan(&goods.GoodsId, &goods.SortId, &goods.StoreId, &goods.GoodsName, &goods.Picture, &goods.Price,
			&goods.GoodsIntro, &goods.Turnover, &goods.Style, &goods.Number, &goods.ShelfDate)
		if err != nil {
			return nil, err
		}
		goodses = append(goodses, goods)
	}
	return goodses, err
}
