/*******
* @Author:qingmeng
* @Description:
* @File:order_details
* @Date2022/2/19
 */

package dao

import (
	"gorm.io/gorm"
	"second-hand-trade/model"
	"time"
)

type OrderDao struct {
}

// CreateOrder 传入addressId生成订单
func (d *OrderDao) CreateOrder(addressId int) (int, error) {
	rs, err := DB.Exec("insert into second_hand_trade.order_details (address_id, date) VALUES(?,?) ", addressId, time.Now())
	id, err := rs.LastInsertId()
	return int(id), err
}

// UpdateOrder 更新订单
func (d *OrderDao) UpdateOrder(order model.OrderDetails) error {
	_, err := DB.Exec("update second_hand_trade.order_details set address_id=?,order_state=?,confirm_receipt=?,money=? where order_id=?", order.AddressId, order.OrderState, order.ConfirmReceipt, order.Money, order.OrderId)
	return err
}

// SelectOrderByOrderId 根据oderId返回订单
func (d *OrderDao) SelectOrderByOrderId(orderId int) (model.OrderDetails, error) {
	var order model.OrderDetails
	row := DB.QueryRow("select * from second_hand_trade.order_details where order_id=? ", orderId)
	if row.Err() != nil {
		return order, row.Err()
	}
	err := row.Scan(&order.OrderId, &order.AddressId, &order.Date, &order.OrderState, &order.ConfirmReceipt, &order.Money)
	return order, err
}

func (d *OrderDao) SelectOrdersByUid(uid int) ([]model.OrderDetails, error) {
	var orders []model.OrderDetails
	rows, err := DB.Query("select order_details.order_id, address_id, date, order_state, confirm_receipt, money from second_hand_trade.shopping_cart,second_hand_trade.order_details where order_details.order_id=shopping_cart.order_id and uid=?", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order model.OrderDetails
		err = rows.Scan(&order.OrderId, &order.AddressId, &order.Date, &order.OrderState, &order.ConfirmReceipt, &order.Money)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, err
}

// gorm事务处理提交订单
func (d *OrderDao) SubmitOrder(user model.User, order model.OrderDetails) error {
	err := GormDB.Transaction(func(tx *gorm.DB) error {
		//更新用户余额
		if err := tx.Model(&user).Update("money", gorm.Expr("money-?", order.Money)).Error; err != nil {
			return err
		}

		//验证购物车
		//采用预加载
		if err := tx.Preload("second-hand-tradepingCart").Find(&order).Error; err != nil {
			return err
		}
		for _, cart := range order.ShoppingCarts {
			var goods model.Goods
			if err := tx.First(&goods, cart.GoodsId).Error; err != nil {
				return err
			}

			//更新货物余量
			if err := tx.Model(&goods).Update("number", gorm.Expr("number-?", cart.Number)).Error; err != nil {
				return err
			}

			//更新店铺余额
			if err := tx.Model(&model.Store{}).Where("store_id=?", goods.StoreId).Update("store_money",
				gorm.Expr("store_money+?", goods.Price*float32(cart.Number))).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
