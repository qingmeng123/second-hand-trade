/*******
* @Author:qingmeng
* @Description:
* @File:order_details
* @Date2022/2/19
 */

package service

import (
	"database/sql"
	"second-hand-trade/dao"
	"second-hand-trade/model"
)

type OrderService struct {
}

func (s *OrderService) CreateOrder(addressId int) (int, error) {
	d := dao.OrderDao{}
	return d.CreateOrder(addressId)
}

func (s *OrderService) UpdateOrder(order model.OrderDetails) error {
	d := dao.OrderDao{}
	return d.UpdateOrder(order)
}

func (s *OrderService) SelectOrderByOrderId(orderId int) (model.OrderDetails, error) {
	d := dao.OrderDao{}
	order, err := d.SelectOrderByOrderId(orderId)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.OrderDetails{}, nil
		}
		return order, err
	}

	return order, nil
}

func (s *OrderService) SelectOrdersByUid(uid int) ([]model.OrderDetails, error) {
	d := dao.OrderDao{}
	return d.SelectOrdersByUid(uid)

}

// 提交订单
func (s *OrderService) SubmitOrder(u model.User, order model.OrderDetails) error {
	d := dao.OrderDao{}
	return d.SubmitOrder(u, order)
}
