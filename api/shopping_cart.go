/*******
* @Author:qingmeng
* @Description:
* @File:second-hand-tradeping_cart
* @Date2022/2/19
 */

package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"second-hand-trade/model"
	"second-hand-trade/service"
	"second-hand-trade/tool"
	"strconv"
)

// 修改购物车
func updateCart(ctx *gin.Context) {
	token := ctx.PostForm("token")
	cartId := ctx.PostForm("cart_id")
	number := ctx.PostForm("number")
	remark := ctx.PostForm("remark")
	state := ctx.PostForm("state")
	isDelete := ctx.PostForm("isDelete")
	cs := service.ShoppingCartService{}

	//解析token
	mc, _ := service.ParseToken(token)

	//cart_id
	if cartId == "" {
		tool.RespErrorWithData(ctx, "cart_id无效")
		return
	}
	cid, err := strconv.Atoi(cartId)
	if err != nil {
		fmt.Println("cart id to int err:", err)
		tool.RespErrorWithData(ctx, "cart_id无效")
		return
	}
	cart, err := cs.SelectCartByCartId(cid)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("select cart by cart id err:", err)
			tool.RespInternalError(ctx)
			return
		}

		if err == sql.ErrNoRows {
			tool.RespErrorWithData(ctx, "cart_id无效")
			return
		}
	}
	if cart.Uid != mc.User.Uid {
		tool.RespErrorWithData(ctx, "cart_id无效")
		return
	}

	//number
	if number != "" {
		n, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("number to int err:", err)
			tool.RespErrorWithData(ctx, "number无效")
			return
		}
		if n <= 0 {
			tool.RespErrorWithData(ctx, "number无效")
			return
		}
		err = cs.UpdateNumber(cid, n)
		if err != nil {
			fmt.Println("update number err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	//remark
	if remark != "" {
		flag := tool.CheckIfSensitive(remark)
		if flag {
			tool.RespErrorWithData(ctx, "remark无效")
			return
		}
		err = cs.UpdateRemark(cid, remark)
		if err != nil {
			fmt.Println("update remark err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}
	//state
	if state != "" {
		if state == "true" {
			err = cs.UpdateState(cid, true)
			if err != nil {
				fmt.Println("update state err:", err)
				tool.RespInternalError(ctx)
				return
			}
		} else if state == "false" {
			err = cs.UpdateState(cid, false)
			if err != nil {
				fmt.Println("update state err:", err)
				tool.RespInternalError(ctx)
				return
			}
		} else {
			tool.RespErrorWithData(ctx, "state无效")
			return
		}
	}
	//isDelete
	if isDelete == "true" {
		err = cs.DeleteCartByCartId(cid)
		if err != nil {
			fmt.Println("delete cart by cart id err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}
	tool.RespSuccessfulWithData(ctx, "修改成功")

}

// 添加到购物车
func addCart(ctx *gin.Context) {
	token := ctx.PostForm("token")
	goodsId := ctx.PostForm("goods_id")
	number := ctx.PostForm("number")

	cart := model.ShoppingCart{}
	cs := service.ShoppingCartService{}
	gs := service.GoodsService{}
	//解析token.(已鉴权过，则no err)
	mc, _ := service.ParseToken(token)
	cart.Uid = mc.User.Uid

	//goodsId
	if goodsId == "" {
		tool.RespErrorWithData(ctx, "goods_id无效")
		return
	}
	gid, err := strconv.Atoi(goodsId)
	if err != nil {
		fmt.Println("goods id to int err:", err)
		tool.RespErrorWithData(ctx, "goods_id无效")
		return
	}
	flag, err := gs.IsExistGoodsId(gid)
	if err != nil {
		fmt.Println("judge exist goods id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "goods_id无效")
		return
	}
	cart.GoodsId = gid

	//number
	n := 1
	if number != "" {
		n, err = strconv.Atoi(number)
		if err != nil {
			fmt.Println("number to int err:", err)
			tool.RespErrorWithData(ctx, "number无效")
			return
		}
		if n <= 0 {
			tool.RespErrorWithData(ctx, "number无效")
			return
		}
	}
	cart.Number = n

	cart.State = false
	err = cs.AddCart(cart)
	if err != nil {
		fmt.Println("add cart err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "添加成功")
}

// 查看某用户的购物车
func viewCartByUid(ctx *gin.Context) {
	token := ctx.PostForm("token")
	mc, _ := service.ParseToken(token)
	cs := service.ShoppingCartService{}
	carts, err := cs.SelectCartsByUid(mc.User.Uid)
	if err != nil {
		fmt.Println("select carts by uid err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if carts == nil {
		tool.RespErrorWithData(ctx, "购物车为空")
		return
	}
	tool.RespSuccessfulWithData(ctx, carts)
}
