/*******
* @Author:qingmeng
* @Description:
* @File:shopping_cart
* @Date2022/2/19
 */

package dao

import (
	"second-hand-trade/model"
)

type ShoppingCartService struct {
}

// AddCart 添加到购物车
func (s *ShoppingCartService) AddCart(cart model.ShoppingCart) error {
	_, err := DB.Exec("insert into second_hand_trade.shopping_cart( uid, goods_id, number, state) values (?,?,?,?)", cart.Uid, cart.GoodsId, cart.Number, cart.State)
	return err
}

// SelectCartByCartId 通过CartId查看购物车中的商品
func (s *ShoppingCartService) SelectCartByCartId(id int) (model.ShoppingCart, error) {
	var cart model.ShoppingCart
	row := DB.QueryRow("select * from second_hand_trade.shopping_cart where cart_id=? and order_id=0", id)
	if row.Err() != nil {
		return cart, row.Err()
	}
	err := row.Scan(&cart.CartId, &cart.Uid, &cart.GoodsId, &cart.Number, &cart.Remark, &cart.State, &cart.OrderId)
	return cart, err
}

// UpdateNumber 修改购物车商品数量
func (s *ShoppingCartService) UpdateNumber(id int, n int) error {
	_, err := DB.Exec("update second_hand_trade.shopping_cart set number=? where cart_id=?", n, id)
	return err
}

// UpdateState 修改状态
func (s *ShoppingCartService) UpdateState(cid int, b bool) error {
	_, err := DB.Exec("update second_hand_trade.shopping_cart set state=? where cart_id=?", b, cid)
	return err
}

// UpdateRemark 修改备注
func (s *ShoppingCartService) UpdateRemark(id int, remark string) error {
	_, err := DB.Exec("update second_hand_trade.shopping_cart set remark=? where cart_id =?", remark, id)
	return err
}

// DeleteCartByCartId 删除购物车
func (s *ShoppingCartService) DeleteCartByCartId(cid int) error {
	_, err := DB.Exec("delete from second_hand_trade.shopping_cart where cart_id=?", cid)
	return err
}

// SelectCartsByUid 根据uid返回carts
func (s *ShoppingCartService) SelectCartsByUid(uid int) ([]model.ShoppingCart, error) {
	var carts []model.ShoppingCart
	rows, err := DB.Query("select * from second_hand_trade.shopping_cart where uid=? and order_id=0", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var cart model.ShoppingCart
		err = rows.Scan(&cart.CartId, &cart.Uid, &cart.GoodsId, &cart.Number, &cart.Remark, &cart.State, &cart.OrderId)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

// CartsIntoOrderByUid 将购物车的商品生成到订单
func (s *ShoppingCartService) CartsIntoOrderByUid(uid int, orderId int) error {
	_, err := DB.Exec("update second_hand_trade.shopping_cart set order_id=? where uid=? and state=1", orderId, uid)
	return err
}

// SelectCartsByOrderId 根据orderId返回购物车
func (s *ShoppingCartService) SelectCartsByOrderId(orderId int) ([]model.ShoppingCart, error) {
	var carts []model.ShoppingCart
	rows, err := DB.Query("select * from second_hand_trade.shopping_cart where order_id=? and state=1", orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var cart model.ShoppingCart
		err = rows.Scan(&cart.CartId, &cart.Uid, &cart.GoodsId, &cart.Number, &cart.Remark, &cart.State, &cart.OrderId)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

// SelectCartsByUidGoodsId 根据uid和goodsId返回已收货的cart
func (s *ShoppingCartService) SelectCartsByUidGoodsId(uid int, goodsId int) (model.ShoppingCart, error) {
	var cart model.ShoppingCart
	row := DB.QueryRow("select shopping_cart.cart_id, shopping_cart.uid, shopping_cart.goods_id, shopping_cart.number, shopping_cart.remark, shopping_cart.state, shopping_cart.order_id from second_hand_trade.shopping_cart,second_hand_trade.order_details where uid=? and goods_id=? and shopping_cart.order_id<>0 and shopping_cart.order_id=order_details.order_id and confirm_receipt=true", uid, goodsId)
	if row.Err() != nil {
		return cart, row.Err()
	}
	err := row.Scan(&cart.CartId, &cart.Uid, &cart.GoodsId, &cart.Number, &cart.Remark, &cart.State, &cart.OrderId)
	return cart, err
}
