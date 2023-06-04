/*******
* @Author:qingmeng
* @Description:
* @File:store
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
)

// 创建店铺
func createStore(ctx *gin.Context) {
	token := ctx.PostForm("token")
	storeName := ctx.PostForm("store_name")
	notice := ctx.PostForm("notice")
	us := service.UserService{}
	ss := service.StoreService{}
	//验证商家
	mc, _ := service.ParseToken(token)
	if mc.User.GroupId != 1 {
		tool.RespErrorWithData(ctx, "该用户不是商家")
		return
	}
	if mc.User.StoreId != 0 {
		tool.RespErrorWithData(ctx, "该用户已有店铺")
		return
	}

	//storeName
	if storeName == "" {
		tool.RespErrorWithData(ctx, "store_name无效")
		return
	}
	flag := tool.CheckIfSensitive(storeName)
	if flag {
		tool.RespErrorWithData(ctx, "store_name无效")
		return
	}
	flag, err := ss.IsExistStoreName(storeName)
	if err != nil {
		fmt.Println("judge exist store name err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag {
		tool.RespErrorWithData(ctx, "该店铺名已被注册")
		return
	}

	//notice
	if notice == "" {
		notice = "暂无公告"
	}
	flag = tool.CheckIfSensitive(notice)
	if flag {
		tool.RespErrorWithData(ctx, "notice无效")
		return
	}

	store := model.Store{
		StoreName:  storeName,
		Notice:     notice,
		StoreMoney: 0,
	}

	err = ss.CreateStore(store)
	if err != nil {
		fmt.Println("create store err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//加入店铺
	store, err = ss.SelectStoreByStoreName(storeName)
	if err != nil {
		fmt.Println("select store by store name err:", err)
		tool.RespInternalError(ctx)
		return
	}
	err = us.AddStoreUser(mc.User.Username, store.StoreId)
	if err != nil {
		fmt.Println("add store user err:", err)
		tool.RespInternalError(ctx)
		return
	}
	ctx.JSON(200, gin.H{
		"status":     true,
		"data":       "创建成功",
		"store_id":   store.StoreId,
		"store_name": store.StoreName,
		"notice":     store.Notice,
	})
}

// 查看店铺
func viewStore(ctx *gin.Context) {
	storeName := ctx.PostForm("store_name")
	storeId := ctx.PostForm("store_id")
	ss := service.StoreService{}

	if storeName == "" && storeId == "" {
		tool.RespErrorWithData(ctx, "参数无效")
		return
	}

	//通过storeName
	if storeName != "" {
		flag, err := ss.IsExistStoreName(storeName)
		if err != nil {
			fmt.Println("judge exist store name err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if !flag {
			tool.RespErrorWithData(ctx, "参数无效")
			return
		}

		store, err := ss.SelectStoreByStoreName(storeName)
		if err != nil {
			fmt.Println("select store by store name err:", err)
			tool.RespInternalError(ctx)
			return
		}
		ctx.JSON(200, gin.H{
			"status":     true,
			"data":       "查看成功",
			"store_id":   store.StoreId,
			"store_name": store.StoreName,
			"notice":     store.Notice,
		})
		return
	}

	//通过storeId
	if storeId != "" {
		sid, err := strconv.Atoi(storeId)
		if err != nil {
			fmt.Println("store id to int err:", err)
			tool.RespErrorWithData(ctx, "store_id无效")
			return
		}
		flag, err := ss.IsExistStoreId(sid)
		if err != nil {
			fmt.Println("judge exist store id err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if !flag {
			tool.RespErrorWithData(ctx, "参数无效")
			return
		}

		store, err := ss.SelectStoreByStoreId(sid)
		if err != nil {
			fmt.Println("select store by store id err:", err)
			tool.RespInternalError(ctx)
			return
		}
		ctx.JSON(200, gin.H{
			"status":     true,
			"data":       "查看成功",
			"store_id":   store.StoreId,
			"store_name": store.StoreName,
			"notice":     store.Notice,
			"money":      store.StoreMoney,
		})
		return
	}

}

// 给店铺充值
func addStoreMoney(ctx *gin.Context) {
	iUsername, _ := ctx.Get("iUsername")
	storeId := ctx.PostForm("store_id")
	iMoney := ctx.PostForm("money")
	username := iUsername.(string)
	us := service.UserService{}
	ss := service.StoreService{}

	//验证storeId
	if storeId == "" {
		tool.RespErrorWithData(ctx, "store_id无效")
		return
	}
	sid, err := strconv.Atoi(storeId)
	if err != nil {
		fmt.Println("store id to int err:", err)
		tool.RespErrorWithData(ctx, "store_id无效")
		return
	}
	flag, err := ss.IsExistStoreId(sid)
	if err != nil {
		fmt.Println("judge exist store id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "store_id无效")
		return
	}

	//验证user
	user, err := us.GetUserinfoByUserName(username)
	if err != nil {
		fmt.Println("get userinfo err:", err)
		tool.RespInternalError(ctx)
		return
	}
	money, err := strconv.ParseFloat(iMoney, 32)
	if err != nil {
		fmt.Println("money to float err:", err)
		tool.RespErrorWithData(ctx, "money无效")
		return
	}
	if money <= 0 {
		tool.RespErrorWithData(ctx, "money无效")
		return
	}
	if user.Money < float32(money) {
		tool.RespSuccessfulWithData(ctx, "用户余额不足")
		return
	}

	//充值
	err = us.UpdateMoney(username, user.Money-float32(money))
	if err != nil {
		fmt.Println("update user money err:", err)
		tool.RespInternalError(ctx)
		return
	}

	err = ss.AddStoreMoney(sid, money)
	if err != nil {
		fmt.Println("add store money err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "成功充值"+iMoney+"元")
}

// 修改店铺信息
func updateStore(ctx *gin.Context) {
	token := ctx.PostForm("token")
	storeName := ctx.PostForm("store_name")
	notice := ctx.PostForm("notice")
	mc, _ := service.ParseToken(token)
	ss := service.StoreService{}

	//验证商家
	if mc.User.GroupId != 1 {
		tool.RespErrorWithData(ctx, "该用户不是商家")
		return
	}

	//取出店铺
	store, err := ss.SelectStoreByStoreId(mc.User.StoreId)
	if err != nil {
		fmt.Println("select store by store name err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//storeName
	if storeName != "" {
		flag := tool.CheckIfSensitive(storeName)
		if flag {
			tool.RespErrorWithData(ctx, "store_name无效")
			return
		}
		flag, err = ss.IsExistStoreName(storeName)
		if err != nil {
			fmt.Println("judge exist store name err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if flag {
			tool.RespErrorWithData(ctx, "该店铺名已被注册")
			return
		}
		store.StoreName = storeName
	}

	//notice
	if notice != "" {
		flag := tool.CheckIfSensitive(notice)
		if flag {
			tool.RespErrorWithData(ctx, "notice无效")
			return
		}
		store.Notice = notice
	}

	err = ss.UpdateStore(store)
	if err != nil {
		fmt.Println("update store err:", err)
		tool.RespInternalError(ctx)
		return
	}

	ctx.JSON(200, gin.H{
		"status":     true,
		"data":       "修改成功",
		"store_id":   store.StoreId,
		"store_name": store.StoreName,
		"notice":     store.Notice,
	})
}
