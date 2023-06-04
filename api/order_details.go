/*******
* @Author:qingmeng
* @Description:
* @File:order_details
* @Date2022/2/19
 */

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"second-hand-trade/model"
	"second-hand-trade/service"
	"second-hand-trade/tool"
	"strconv"
	"time"
)

// 生成订单
func createOrder(ctx *gin.Context) {
	iAddressId := ctx.PostForm("address_id")
	iUsername, _ := ctx.Get("iUsername")
	username := iUsername.(string)

	cs := service.ShoppingCartService{}
	os := service.OrderService{}
	us := service.UserService{}
	gs := service.GoodsService{}

	user, err := us.GetUserinfoByUserName(username)
	if err != nil {
		fmt.Println("get userinfo err:", err)
		tool.RespInternalError(ctx)
		return
	}
	//预生成订单
	if user.AddressId == 0 && iAddressId == "" {
		tool.RespErrorWithData(ctx, "地址无效")
		return
	}
	var addressId int
	if iAddressId == "" {
		addressId = user.AddressId
	} else {
		addressId, err = strconv.Atoi(iAddressId)
		if err != nil {
			fmt.Println("addressId to int err:", err)
			tool.RespErrorWithData(ctx, "address_id无效")
			return
		}
	}
	orderId, err := os.CreateOrder(addressId)
	if err != nil {
		fmt.Println("create order err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//将购物车加入订单
	err = cs.CartsIntoOrderByUid(user.Uid, orderId)
	if err != nil {
		fmt.Println("carts into order err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//取出加入订单的购物车
	carts, err := cs.SelectCartsByOrderId(orderId)
	if err != nil {
		fmt.Println("select carts by order id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if carts == nil {
		tool.RespErrorWithData(ctx, "该用户没商品可加入订单")
		return
	}

	//计算金额
	var sum float32
	for _, cart := range carts {
		goods, err := gs.SelectGoodsByGoodsId(cart.GoodsId)
		if err != nil {
			fmt.Println("select goods by goods id err:", err)
			tool.RespInternalError(ctx)
			return
		}
		s := float32(cart.Number) * goods.Price
		sum += s
	}
	if sum < 0 {
		fmt.Println("sum err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//生成订单
	order := model.OrderDetails{
		OrderId:        orderId,
		Money:          sum,
		AddressId:      addressId,
		Date:           time.Now(),
		OrderState:     false,
		ConfirmReceipt: false,
	}

	err = os.UpdateOrder(order)
	if err != nil {
		fmt.Println("update order err:", err)
		tool.RespInternalError(ctx)
		return
	}

	outPutMap := tool.ObjToMap(order)
	tool.RespSuccessfulWithData(ctx, outPutMap)
}

// 查看所有订单
func viewOrder(ctx *gin.Context) {
	token := ctx.PostForm("token")
	os := service.OrderService{}
	mc, _ := service.ParseToken(token)
	orders, err := os.SelectOrdersByUid(mc.User.Uid)
	if err != nil {
		fmt.Println("select order by uid err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, orders)
}

// 查看订单详情
func viewOrderDetails(ctx *gin.Context) {
	iOrderId := ctx.PostForm("order_id")

	os := service.OrderService{}
	ss := service.ShoppingCartService{}

	//取出订单,验证orderId
	if iOrderId == "" {
		tool.RespErrorWithData(ctx, "order_id无效")
		return
	}
	orderId, err := strconv.Atoi(iOrderId)
	if err != nil {
		fmt.Println("order_id to int err:", err)
		tool.RespErrorWithData(ctx, "order_id无效")
		return
	}
	order, err := os.SelectOrderByOrderId(orderId)
	if err != nil {
		fmt.Println("select order by order id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if order.OrderId == 0 {
		tool.RespErrorWithData(ctx, "order_id无效")
		return
	}

	//取出购物车
	carts, err := ss.SelectCartsByOrderId(orderId)
	if err != nil {
		fmt.Println("select carts by orderId err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if carts == nil {
		tool.RespErrorWithData(ctx, "该订单无记录")
		return
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"order":  order,
		"carts":  carts,
	})

}

// 提交订单
func submitOrder(ctx *gin.Context) {
	iOrderId := ctx.PostForm("order_id")
	iUsername, _ := ctx.Get("iUsername")
	os := service.OrderService{}
	us := service.UserService{}

	//取出订单,验证orderId
	if iOrderId == "" {
		tool.RespErrorWithData(ctx, "order_id无效")
		return
	}
	orderId, err := strconv.Atoi(iOrderId)
	if err != nil {
		fmt.Println("order_id to int err:", err)
		tool.RespErrorWithData(ctx, "order_id无效")
		return
	}
	order, err := os.SelectOrderByOrderId(orderId)
	if err != nil {
		fmt.Println("select order by order id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if order.OrderId == 0 {
		tool.RespErrorWithData(ctx, "order_id无效")
		return
	}

	//验证订单
	if order.OrderState == true {
		tool.RespErrorWithData(ctx, "该订单已支付")
		return
	}

	//验证余额
	username := iUsername.(string)
	user, err := us.GetUserinfoByUserName(username)
	if err != nil {
		fmt.Println("get userinfo err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if user.Money < order.Money {
		tool.RespErrorWithData(ctx, "支付失败，余额不足")
		return
	}

	//通过事务处理订单
	err = os.SubmitOrder(user, order)
	if err != nil {
		fmt.Println("submit order err:", err)
		tool.RespErrorWithData(ctx, "订单提交未成功")
		return
	}

	//支付
	err = us.UpdateMoney(user.Username, user.Money-order.Money)
	if err != nil {
		fmt.Println("update money err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//店铺收钱
	ss := service.StoreService{}
	cs := service.ShoppingCartService{}
	gs := service.GoodsService{}

	//取出加入订单的购物车
	carts, err := cs.SelectCartsByOrderId(orderId)
	if err != nil {
		fmt.Println("select carts by order id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	for _, cart := range carts {
		goods, err := gs.SelectGoodsByGoodsId(cart.GoodsId)
		if err != nil {
			fmt.Println("select goods by goods id err:", err)
			tool.RespInternalError(ctx)
			return
		}
		err = ss.AddStoreMoney(goods.StoreId, float64(float32(cart.Number)*goods.Price))
		if err != nil {
			fmt.Println("add store money err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	//更新订单
	order.OrderState = true
	err = os.UpdateOrder(order)
	if err != nil {
		fmt.Println("update order err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "提交成功")
}

// 修改订单状态
func updateOrder(ctx *gin.Context) {
	iOrderId := ctx.PostForm("order_id")
	iUsername, _ := ctx.Get("iUsername")
	orderState := ctx.PostForm("order_state")
	confirmReceipt := ctx.PostForm("confirm_receipt")
	os := service.OrderService{}
	us := service.UserService{}

	//取出订单,验证orderId
	if iOrderId == "" {
		tool.RespErrorWithData(ctx, "order_id无效")
		return
	}
	orderId, err := strconv.Atoi(iOrderId)
	if err != nil {
		fmt.Println("order_id to int err:", err)
		tool.RespErrorWithData(ctx, "order_id无效")
		return
	}
	order, err := os.SelectOrderByOrderId(orderId)
	if err != nil {
		fmt.Println("select order by order id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if order.OrderId == 0 {
		tool.RespErrorWithData(ctx, "order_id无效")
		return
	}

	//确认收货
	if confirmReceipt == "true" {
		if order.OrderState == false {
			tool.RespErrorWithData(ctx, "确认收货失败")
			return
		}
		order.ConfirmReceipt = true
		err = os.UpdateOrder(order)
		if err != nil {
			fmt.Println("update order err:", err)
			tool.RespInternalError(ctx)
			return
		}

		//添加商品成交量
		gs := service.GoodsService{}
		cs := service.ShoppingCartService{}
		carts, err := cs.SelectCartsByOrderId(orderId)
		if err != nil {
			fmt.Println("select carts by order id err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if carts == nil {
			tool.RespErrorWithData(ctx, "确认收货失败")
			return
		}
		for _, cart := range carts {
			err = gs.AddTurnover(cart.GoodsId, cart.Number)
			if err != nil {
				fmt.Println("add turnover err:", err)
				tool.RespInternalError(ctx)
				return
			}
		}

		tool.RespSuccessfulWithData(ctx, "确认收货成功")
		return
	}

	//取消订单
	if orderState == "false" {
		if order.OrderState == false || order.ConfirmReceipt == true {
			tool.RespErrorWithData(ctx, "取消订单失败")
			return
		}

		//退款
		username := iUsername.(string)
		user, err := us.GetUserinfoByUserName(username)
		if err != nil {
			fmt.Println("get userinfo err:", err)
			tool.RespInternalError(ctx)
			return
		}
		err = us.UpdateMoney(user.Username, user.Money+order.Money)
		if err != nil {
			fmt.Println("update money err:", err)
			tool.RespInternalError(ctx)
			return
		}

		//更新订单
		order.OrderState = false
		err = os.UpdateOrder(order)
		if err != nil {
			fmt.Println("update order err:", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessfulWithData(ctx, "退款成功")
		return
	}
	tool.RespSuccessfulWithData(ctx, "未作任何修改")
}
