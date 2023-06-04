/*******
* @Author:qingmeng
* @Description:
* @File:second-hand-tradeping_cart
* @Date2022/2/19
 */

package service

import (
	"database/sql"
	"second-hand-trade/dao"
	"second-hand-trade/model"
)

type ShoppingCartService struct {
}

func (s *ShoppingCartService) AddCart(cart model.ShoppingCart) error {
	d := dao.ShoppingCartService{}
	return d.AddCart(cart)
}

func (s *ShoppingCartService) SelectCartByCartId(id int) (model.ShoppingCart, error) {
	d := dao.ShoppingCartService{}
	return d.SelectCartByCartId(id)
}

func (s *ShoppingCartService) UpdateNumber(id int, n int) error {
	d := dao.ShoppingCartService{}
	return d.UpdateNumber(id, n)
}

func (s *ShoppingCartService) UpdateRemark(id int, remark string) error {
	d := dao.ShoppingCartService{}
	return d.UpdateRemark(id, remark)
}

func (s *ShoppingCartService) UpdateState(cid int, b bool) error {
	d := dao.ShoppingCartService{}
	return d.UpdateState(cid, b)
}

func (s *ShoppingCartService) DeleteCartByCartId(cid int) error {
	d := dao.ShoppingCartService{}
	return d.DeleteCartByCartId(cid)
}

func (s *ShoppingCartService) SelectCartsByUid(uid int) ([]model.ShoppingCart, error) {
	d := dao.ShoppingCartService{}
	return d.SelectCartsByUid(uid)

}

// CartsIntoOrderByUid 将购物车的商品生成到订单
func (s *ShoppingCartService) CartsIntoOrderByUid(uid int, orderId int) error {
	d := dao.ShoppingCartService{}
	return d.CartsIntoOrderByUid(uid, orderId)
}

// SelectCartsByOrderId 根据orderId返回carts
func (s *ShoppingCartService) SelectCartsByOrderId(orderId int) ([]model.ShoppingCart, error) {
	d := dao.ShoppingCartService{}
	return d.SelectCartsByOrderId(orderId)
}

func (s *ShoppingCartService) IsExistOrderByUidGoodsId(uid int, goodsId int) (bool, error) {
	d := dao.ShoppingCartService{}
	_, err := d.SelectCartsByUidGoodsId(uid, goodsId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
