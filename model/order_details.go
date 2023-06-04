/*******
* @Author:qingmeng
* @Description:
* @File:order_details
* @Date2022/2/19
 */

package model

import "time"

type OrderDetails struct {
	OrderId        int            `json:"order_id" gorm:"primary_key"`
	Money          float32        `json:"money"` //该订单金额
	AddressId      int            `json:"address_id"`
	Date           time.Time      `json:"date"`            //下单日期
	OrderState     bool           `json:"order_state"`     //订单状态，false则取消订单
	ConfirmReceipt bool           `json:"confirm_receipt"` //是否确认收货
	ShoppingCarts  []ShoppingCart `gorm:"foreignKey:OrderId"`
}

func (OrderDetails) TableName() string {
	return "order_details"
}
