/*******
* @Author:qingmeng
* @Description:
* @File:router
* @Date2021/12/10
 */

package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()
	engine.Use(Cors())
	apiGroup := engine.Group("/api")
	{
		apiGroup.POST("/user", register) //注册
		apiGroup.GET("/user", login)     //登陆

		userGroup := apiGroup.Group("/user", jwtAuth)
		{
			userGroup.POST("/info", changeInformation) //修改信息
			userGroup.GET("/info", viewUserInfo)       //查看信息
			userGroup.POST("/money", addMoney)         //充值
			userGroup.GET("/money", viewUserMoney)     //查看个人余额
			userGroup.POST("/power", upgradePower)     //注册商家
			userGroup.PUT("/store", addStoreUser)      //商家入铺

			//密码组
			passwordGroup := userGroup.Group("/password")
			{
				passwordGroup.POST("", changePassword) //登陆后的直接修改密码
			}

			//收货地址
			userGroup.PUT("/info/address", addAddress)     //增加收货地址
			userGroup.POST("/info/address", updateAddress) //修改收货地址
			userGroup.GET("/info/address", viewAddresses)  //查看收货地址

		}

		//商品
		goodsGroup := apiGroup.Group("/goods")
		{
			goodsGroup.PUT("", jwtAuth, addGoods)     //添加商品
			goodsGroup.GET("", viewGoods)             //商品主页
			goodsGroup.POST("", jwtAuth, updateGoods) //修改商品

			//类别
			goodsGroup.PUT("/sort", jwtAuth, addSort) //添加类别
			goodsGroup.GET("/sort", viewSort)         //查看所有类别

			//购物车
			goodsGroup.PUT("/shopping", jwtAuth, addCart)       //添加到购物车
			goodsGroup.POST("/shopping", jwtAuth, updateCart)   //修改购物车
			goodsGroup.GET("/shopping", jwtAuth, viewCartByUid) //查看某用户的购物车

			//订单
			orderGroup := goodsGroup.Group("/order", jwtAuth)
			{
				orderGroup.PUT("", createOrder)              //生成订单
				orderGroup.POST("", submitOrder)             //提交订单
				orderGroup.GET("", viewOrder)                //查看所有订单
				orderGroup.GET("/details", viewOrderDetails) //查看订单详情
				orderGroup.POST("/state", updateOrder)       //修改订单状态
			}

		}

		//店铺
		apiGroup.PUT("/store", jwtAuth, createStore)          //创建店铺
		apiGroup.POST("/store", jwtAuth, updateStore)         //更新店铺
		apiGroup.GET("/store", viewStore)                     //查看店铺
		apiGroup.POST("/store/money", jwtAuth, addStoreMoney) //给店铺充值
	}

	engine.Run(":8080")
}
