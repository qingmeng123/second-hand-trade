/*******
* @Author:qingmeng
* @Description:
* @File:address
* @Date2022/2/18
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

// 查看收货地址
func viewAddresses(ctx *gin.Context) {
	token := ctx.PostForm("token")

	//解析token
	mc, err := service.ParseToken(token)
	if err != nil {
		fmt.Println("token err:", err.Error())
		tool.RespSuccessfulWithData(ctx, "token无效")
		ctx.Abort()
		return
	}
	as := service.AddressService{}
	addresses, err := as.SelectAddressesByUid(mc.User.Uid)
	if err != nil {
		fmt.Println("select addresses err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if addresses == nil {
		tool.RespErrorWithData(ctx, "无保存地址，请添加地址")
		return
	}
	ctx.JSON(200, gin.H{
		"status":    true,
		"data":      "查询成功",
		"addresses": addresses,
	})

}

// 修改收货地址
func updateAddress(ctx *gin.Context) {
	iAddressId := ctx.PostForm("address_id")
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	address := ctx.PostForm("address")
	as := service.AddressService{}

	//检查该addressId是否有效
	if iAddressId == "" {
		tool.RespErrorWithData(ctx, "address_id无效")
		return
	}
	addressId, err := strconv.Atoi(iAddressId)
	if err != nil {
		fmt.Println("addressId to int err:", err)
		tool.RespErrorWithData(ctx, "address_id无效")
		return
	}
	addressInfo, err := as.GetAddressInfoByAddressId(addressId)
	if err != nil {
		tool.RespErrorWithData(ctx, "address_id无效")
		return
	}

	if name != "" {
		//检查name
		flag := tool.CheckIfSensitive(name)
		if flag {
			tool.RespErrorWithData(ctx, "收货昵称不合法")
			return
		}
		addressInfo.Name = name
	}

	if phone != "" {
		//检查phone
		if len(phone) != 11 {
			tool.RespErrorWithData(ctx, "收货电话不合法")
			return
		}
		addressInfo.Phone = phone
	}

	if address != "" {
		//检查address
		flag := tool.CheckIfSensitive(address)
		if flag {
			tool.RespErrorWithData(ctx, "收货地址不合法")
			return
		}
		addressInfo.Address = address
	}

	err = as.UpdateAddress(addressInfo)
	if err != nil {
		fmt.Println("update address err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "修改成功")
}

// 添加收货地址
func addAddress(ctx *gin.Context) {
	token := ctx.PostForm("token")
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	address := ctx.PostForm("address")

	//解析token
	mc, err := service.ParseToken(token)
	if err != nil {
		fmt.Println("token err:", err.Error())
		tool.RespSuccessfulWithData(ctx, "token无效")
		ctx.Abort()
		return
	}

	flag := tool.CheckIfSensitive(name)
	if flag {
		tool.RespErrorWithData(ctx, "收货昵称不合法")
		return
	}
	//检查phone
	if len(phone) != 11 {
		tool.RespErrorWithData(ctx, "收货电话不合法")
		return
	}

	//检查address
	if address == "" {
		tool.RespErrorWithData(ctx, "收货地址不合法")
		return
	}
	flag = tool.CheckIfSensitive(address)
	if flag {
		tool.RespErrorWithData(ctx, "收货地址不合法")
		return
	}

	as := service.AddressService{}
	addressInfo := model.AddressInfo{
		Uid:     mc.User.Uid,
		Name:    name,
		Phone:   phone,
		Address: address,
	}

	//检查该地址信息是否已存在
	flag, err = as.IsExistAddress(addressInfo)
	if err != nil {
		fmt.Println("check is exist address err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag {
		tool.RespErrorWithData(ctx, "地址重复")
		return
	}

	err = as.InsertAddress(addressInfo)
	if err != nil {
		fmt.Println("insert address err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "增加成功")
}
