/*******
* @Author:qingmeng
* @Description:
* @File:goods
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
	"time"
)

// 商品主页
func viewGoods(ctx *gin.Context) {
	goodsName := ctx.PostForm("goods_name")            //商品关键字
	iSortId := ctx.PostForm("sort_id")                 //商品类别
	isPriceDesc := ctx.PostForm("is_price_desc")       //按商品价格排序，true降序，false升序,其他乱序
	isTurnoverDesc := ctx.PostForm("is_turnover_desc") //按商品销量排序，true降序，false升序。其他乱序
	gs := service.GoodsService{}
	ss := service.SortService{}
	var goodses []model.Goods
	var err error

	//goodsName
	if goodsName != "" && iSortId == "" {
		goodses, err = gs.SelectGoodsesByName(goodsName)
		if err != nil {
			fmt.Println("select goodses by name err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	//sortId
	if iSortId != "" {
		sortId, err := strconv.Atoi(iSortId)
		if err != nil {
			fmt.Println("sort id to int err:", err)
			tool.RespErrorWithData(ctx, "sort_id无效")
			return
		}
		flag, err := ss.IsExistSortId(sortId)
		if err != nil {
			fmt.Println("judge exist sort id err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if !flag {
			tool.RespErrorWithData(ctx, "暂无该分类")
			return
		}
		if goodsName == "" {
			goodses, err = gs.SelectGoodsesBySortId(sortId)
			if err != nil {
				fmt.Println("select goodses by sort id err:", err)
				tool.RespInternalError(ctx)
				return
			}
		}
		if goodsName != "" {
			goodses, err = gs.SelectGoodsesByGoodsNameSortId(goodsName, sortId)
			if err != nil {
				fmt.Println("select goodses by sort id err:", err)
				tool.RespInternalError(ctx)
				return
			}
		}
	} else {
		//goodsName和sortName都为空
		goodses, err = gs.ViewGoods()
		if err != nil {
			fmt.Println("view goods err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}
	if goodses == nil {
		tool.RespSuccessfulWithData(ctx, "暂无该商品")
		return
	}

	//排序
	goodses = tool.SortGoodsByNature(goodses, "price", isPriceDesc)
	goodses = tool.SortGoodsByNature(goodses, "turnover", isTurnoverDesc)

	tool.RespSuccessfulWithData(ctx, goodses)

}

// 更新商品
func updateGoods(ctx *gin.Context) {
	token := ctx.PostForm("token")
	goodsId := ctx.PostForm("goods_id")
	goodsName := ctx.PostForm("goods_name")
	iSortId := ctx.PostForm("sort_id")
	picture := ctx.PostForm("picture")
	price := ctx.PostForm("price")
	goodsIntro := ctx.PostForm("goods_intro")
	style := ctx.PostForm("type")
	number := ctx.PostForm("number")

	goods := model.Goods{}
	gs := service.GoodsService{}

	//取出商品
	gid, err := strconv.Atoi(goodsId)
	if err != nil {
		fmt.Println("goods id to int err:", err)
		tool.RespInternalError(ctx)
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
	goods, err = gs.SelectGoodsByGoodsId(gid)

	//验证身份
	mc, err := service.ParseToken(token)
	if mc.User.GroupId != 1 || mc.User.StoreId != goods.StoreId {
		tool.RespErrorWithData(ctx, "该用户不支持修改商品")
		return
	}

	//验证sortId
	if iSortId != "" {
		goods.SortId, err = strconv.Atoi(iSortId)
		if err != nil {
			fmt.Println("sort_id to int err:", err)
			tool.RespErrorWithData(ctx, "sort_id无效")
			return
		}
	}

	//验证goodsName
	if goodsName != "" {
		flag := tool.CheckIfSensitive(goodsName)
		if flag {
			tool.RespErrorWithData(ctx, "goods_name无效")
			return
		}
		goods.GoodsName = goodsName
	}

	//验证picture
	if picture != "" {
		goods.Picture = picture
	}

	//验证goodsIntro
	if goodsIntro != "" {
		flag := tool.CheckIfSensitive(goodsIntro)
		if flag {
			tool.RespErrorWithData(ctx, "goods_intro无效")
			return
		}
		goods.GoodsIntro = goodsIntro
	}

	//验证style
	if style != "" {
		flag := tool.CheckIfSensitive(style)
		if flag {
			tool.RespErrorWithData(ctx, "style无效")
			return
		}
		goods.Style = style
	}

	//number
	if number != "" {
		n, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("number to int err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if n <= 0 {
			tool.RespErrorWithData(ctx, "number无效")
			return
		}
	}

	//price
	if price != "" {
		p, err := strconv.ParseFloat(price, 32)
		if err != nil {
			fmt.Println("price to int err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if p < 0 {
			tool.RespErrorWithData(ctx, "price无效")
			return
		}
		goods.Price = float32(p)
	}

	err = gs.UpdateGoods(goods)
	if err != nil {
		fmt.Println("update goods err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, "修改成功")
}

// 添加商品
func addGoods(ctx *gin.Context) {
	token := ctx.PostForm("token")
	goodsName := ctx.PostForm("goods_name")
	iSortId := ctx.PostForm("sort_id")
	picture := ctx.PostForm("picture")
	price := ctx.PostForm("price")
	goodsIntro := ctx.PostForm("goods_intro")
	style := ctx.PostForm("type")
	number := ctx.PostForm("number")

	if goodsName == "" || picture == "" || price == "" || number == "" {
		tool.RespSuccessfulWithData(ctx, "参数不完整")
		return
	}
	//验证身份
	mc, err := service.ParseToken(token)
	if mc.User.GroupId != 1 {
		tool.RespErrorWithData(ctx, "该用户不支持添加商品")
		return
	}

	//验证sortId
	sortId := 0
	if iSortId != "" {
		sortId, err = strconv.Atoi(iSortId)
		if err != nil {
			fmt.Println("sort_id to int err:", err)
			tool.RespErrorWithData(ctx, "sort_id无效")
			return
		}
	}

	//验证goodsName
	flag := tool.CheckIfSensitive(goodsName)
	if flag {
		tool.RespErrorWithData(ctx, "goods_name无效")
		return
	}

	//验证goodsIntro
	flag = tool.CheckIfSensitive(goodsIntro)
	if flag {
		tool.RespErrorWithData(ctx, "goods_intro无效")
		return
	}
	if goodsIntro == "" {
		goodsIntro = "0"
	}
	//验证style
	flag = tool.CheckIfSensitive(style)
	if flag {
		tool.RespErrorWithData(ctx, "style无效")
		return
	}
	if style == "" {
		style = "0"
	}

	//number
	n, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("number to int err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if n <= 0 {
		tool.RespErrorWithData(ctx, "number无效")
		return
	}

	//price
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Println("price to int err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if p < 0 {
		tool.RespErrorWithData(ctx, "price无效")
		return
	}

	goods := model.Goods{
		SortId:     sortId,
		GoodsName:  goodsName,
		Picture:    picture,
		Price:      float32(p),
		GoodsIntro: goodsIntro,
		Style:      style,
		Number:     n,
		StoreId:    mc.User.StoreId,
		ShelfDate:  time.Now(),
	}
	gs := service.GoodsService{}
	err = gs.AddGoods(goods)
	if err != nil {
		fmt.Println("add goods err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "添加成功!")
}
