/*******
* @Author:qingmeng
* @Description:
* @File:shopping_cart
* @Date2022/2/19
 */

package model

type ShoppingCart struct {
	CartId  int  `json:"cart_id" `
	Uid     int  `json:"uid"`
	GoodsId int  `json:"goods_id"`
	Number  int  `json:"number"`
	Remark  int  `json:"remark"`                             //备注
	State   bool `json:"state"`                              //状态，是否选中购买此商品
	OrderId int  `json:"order_id" gorm:"foreignKey:OrderId"` //订单Id,未生成订单则为0
}

func (ShoppingCart) TableName() string {
	return "shopping_cart"
}
