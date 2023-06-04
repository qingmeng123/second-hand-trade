/*******
* @Author:qingmeng
* @Description:
* @File:sort
* @Date2022/2/20
 */

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"second-hand-trade/service"
	"second-hand-trade/tool"
)

// 添加分类
func addSort(ctx *gin.Context) {
	token := ctx.PostForm("token")
	sortName := ctx.PostForm("sort_name")
	ss := service.SortService{}

	//验证身份
	mc, _ := service.ParseToken(token)
	if mc.User.GroupId != 1 {
		tool.RespErrorWithData(ctx, "该用户不支持添加分类")
		return
	}

	//验证sortName
	if sortName == "" {
		tool.RespErrorWithData(ctx, "sort_name无效")
		return
	}
	flag := tool.CheckIfSensitive(sortName)
	if flag {
		tool.RespErrorWithData(ctx, "sort_name无效")
		return
	}
	flag, err := ss.IsExistSortName(sortName)
	if err != nil {
		fmt.Println("judge exist sortName err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag {
		tool.RespErrorWithData(ctx, "已存在该类别")
		return
	}

	//添加分类
	err = ss.AddSort(sortName)
	if err != nil {
		fmt.Println("add sort err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "添加成功")
}

// 查看所有分类
func viewSort(ctx *gin.Context) {
	ss := service.SortService{}
	sorts, err := ss.ViewSort()
	if err != nil {
		fmt.Println("view sort err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if sorts == nil {
		tool.RespSuccessfulWithData(ctx, "暂无分类")
		return
	}

	tool.RespSuccessfulWithData(ctx, sorts)
}
