[查看商品]

##### 简要描述

- multipart/form-data
- 查看商品接口

##### 请求URL

- ` http://101.43.160.254：8080/api/goods`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|goods_name |否  |String |商品名字关键字   |
|sort_id |否  |int |商品类别id   |
|is_price_desc |否  |String |按商品价格排序，true降序，false升序,其他则乱序   |
|is_turnover_desc |否  |String |按商品销量排序，true降序，false升序,其他则乱序   |
|is_feedback_desc |否  |String |按商品好评率排序，true降序，false升序 ,其他则乱序 |

##### 返回示例

``` 
{
    "status": true
    "data": [
        {
            "goods_id": 2,
            "sort_id": 0,
            "store_id": 2,
            "goods_name": "可乐",
            "picture": "home/picture/part-00356-3963.jpg",
            "price": 4,
            "goods_intro": "0",
            "number": 10,
            "turnover": "0",
            "comment_number": 0,
            "feed_back": 0,
            "type": "0",
            "shelf_date": "2022-02-21T06:46:04Z"
        }
    ],
}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----
|false |"goods_name无效"   |goods_name无效 |
|false |"sort_id无效"   |goods_name无效 |
|true |"..."   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[添加商品]

##### 简要描述

- multipart/form-data
- 添加商品接口

##### 请求URL

- ` http://101.43.160.254：8080/api/goods`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|goods_name |是  |String |商品名字   |
|picture |是  |String |商品图片   |
|price |是  |float |商品价格   |
|number |是  |int |商品数量   |
|goods_intro |否  |String |商品介绍   |
|sort_id |否  |int |商品类别   |
|style |否  |String |商品款式   |

##### 返回示例

``` 
 {
	"status":true,
	 "data": {
	"添加成功！"
    }

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |该用户不支持添加商品 ||
|false |"goods_name无效"   |goods_name无效 |
|false |"picture无效"   | |
|false |"price无效"   | |
|false |"goods_intro无效"   | |
|false |"style无效"   | |
|false |"sort_id无效"   | |
|false |"number无效"   | |
|false |"参数不完整"   |未将所有必须参数填充 |
|true |"添加成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[修改商品]

##### 简要描述

- multipart/form-data
- 修改商品接口

##### 请求URL

- ` http://101.43.160.254：8080/api/goods`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|goods_id |是  |int |商品id   |
|goods_name |否  |String |商品名字   |
|picture |否  |String |商品图片   |
|price |否  |float |商品价格   |
|goods_intro |否  |String |商品介绍   |
|sort_id |否  |int |商品类别   |

##### 返回示例

``` 
 {
	"status":true,
	 "data": {
	"修改成功！"
    }

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"goods_id无效"   |该商品修改无效 |
|false |"goods_name无效"   |goods_name无效 |
|false |"picture无效"   | |
|false |"price无效"   | |
|false |"goods_intro无效"   | |
|false |"sort_id无效"   | |
|true |"修改成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[添加商品类别]

##### 简要描述

- multipart/form-data
- 添加商品类别接口

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/sort`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|sort_name |是  |string |类别名称   |

##### 返回示例

``` 
 {
	"status":true,
	 "data": {
	"添加成功！"
    }

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"该用户不支持添加分类"   | |
|false |"goods_name无效"   | |
|false |"已存在该类别"   | |
|true |"添加成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[查看所有类别]

##### 简要描述

- multipart/form-data
- 查看所有类别接口

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/sort`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |

##### 返回示例

``` 
 {
	"status":true,
	 "data": {
	"sort_id":1,
	"sort_name":饮料
    }

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"暂无分类"   | |
|true |"..."   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[加入购物车]

##### 简要描述

- multipart/form-data
- 加入购物车接口

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/shopping`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|goods_id |是  |int |商品id   |
|number |否  |int |商品数量  |

##### 返回示例

``` 
 {
	"status":true,
	 "data": {
	"添加成功！"
    }

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"goods_id无效"   | |
|false |"number无效"   |number为0或达到最大购物限制 |
|true |"添加成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[修改购物车]

##### 简要描述

- multipart/form-data
- 修改购物车接口

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/shopping`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|cart_id |是  |int |购物车id   |
|type |否  |string |商品款式   |
|number |否  |int |商品数量   |
|remark |否  |string |备注  |
|state |否  |bool|状态，是否勾选购买该商品，true为勾选，默认false  |
|isDelete |否  |bool |是否删除该商品，true为删除  |

##### 返回示例

``` 
 {
	"status":true,
	 "data": {
	"修改成功！"
    }

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"goods_id无效"   | |
|false |"type无效"   | |
|false |"remark无效"   | |
|false |"number无效"   |number为0或达到最大购物限制 |
|false |"state无效"   | |
|true |"修改成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[查看购物车]

##### 简要描述

- multipart/form-data
- 查看购物车接口

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/shopping`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |

##### 返回示例

``` 
 {
	"status":true,
	 "data":  [
        {
            "cart_id": 1,
            "uid": 17,
            "goods_id": 2,
            "number": 1,
            "remark": 0,
            "state": false
        },
        {
            "cart_id": 2,
            "uid": 17,
            "goods_id": 2,
            "number": 1,
            "remark": 0,
            "state": false
        },
        {
            "cart_id": 3,
            "uid": 17,
            "goods_id": 2,
            "number": 2,
            "remark": 0,
            "state": true
        }
    ],
}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|true |"购物车为空"   |参数合法  |
|true |"..."   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[生成订单]

##### 简要描述

- multipart/form-data
- 生成订单接口

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/order`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|address_id |否  |int |运送地址信息，不选则用默认信息   |

##### 返回示例

``` 
 {
	"status":true,
	"data": {
        "AddressId": 1,
        "ConfirmReceipt": false,
        "Date": "2022-02-20T16:00:01.3969136+08:00",
        "Money": 310,
        "OrderId": 3,
        "OrderState": false
    },

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"address_id无效"   | |
|true |"该用户没商品可加入订单"   | |
|true |"..."   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[查看订单详情]

##### 简要描述

- multipart/form-data
- 查看订单详情接口

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/order/details`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|order_id |是  |int  |订单id   |

##### 返回示例

``` 
 {
	"status":true,
	  "order": {
        "order_id": 4,
        "money": 310,
        "address_id": 6,
        "date": "2022-02-20T10:20:30Z",
        "order_state": false,
        "confirm_receipt": false
    }
    "carts": [
        {
            "cart_id": 4,
            "uid": 16,
            "goods_id": 2,
            "number": 2,
            "remark": 0,
            "state": true,
            "order_id": 4
        }
    ],

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"order_id无效"   | |
|false |"该单号暂无记录"   | |
|true |"..."   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[查看用户所有订单]

##### 简要描述

- multipart/form-data
- 查看所有订单接口

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/order`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |

##### 返回示例

``` 
 {
	"status":true,
	"data": {
        "AddressId": 6,
        "ConfirmReceipt": false,
        "Date": "2022-02-20T18:20:29.9774043+08:00",
        "Money": 310,
        "OrderId": 4,
        "OrderState": false
    },

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|true |"..."   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[提交订单]

##### 简要描述

- multipart/form-data
- 提交订单接口,用户提交成功时，用户消费扣钱，店铺收益

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/order`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|order_id |是  |int |订单id   |

##### 返回示例

``` 
 {
	"status":true,
	 "data": {
	"提交成功"
    }

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"order_id无效"   | |
|false |"该订单已支付"   | |
|true |"支付失败,余额不足 "   |  |
|true |"提交成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[修改订单状态]

##### 简要描述

- multipart/form-data
- 修改订单状态接口,收货成功添加成交量

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/order/state`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|order_id |是  |int |订单id   |
|order_state |否  |bool |订单状态，false则取消订单   |
|confirm_receipt |否  |bool |收货状态，true则确认收货 |

##### 返回示例

``` 
 {
	"status":true,
	 "data": {
	"修改成功"
    }

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"order_id无效"   | |
|true |"确认收货失败"   |该订单未生效  |
|true |"取消订单失败"   |已收到货或该订单未生效  |
|true |"确认收货成功"   |  |
|true |"未作任何修改"   |  |
|true |"取消订单成功"   |参数合法，退款成功  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[添加评论]

##### 简要描述

- multipart/form-data
- 添加评论接口,收货成功才可评论

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/comment`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|goods_id |是  |int |商品id   |
|text |是  |string |评论内容   |
|star |是  |int |评星（1~5） |

##### 返回示例

``` 
 {
	"status":true,
	 "data": {
	"感谢您的评价"
    }

}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"参数无效"   |参数不完全  |
|false |"goods_id无效"   |  |
|false |"star无效"   |  |
|false |"text无效"   |  |
|false |"该用户对该商品未确认收货"   |  |
|true |"感谢您的评价"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[查看评论]

##### 简要描述

- multipart/form-data
- 查看评论接口,token和good_id至少填一个

##### 请求URL

- ` http://101.43.160.254：8080/api/goods/comment`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |否  |String |用户token   |
|goods_id |否  |int |商品id   |

##### 返回示例

``` 
 {
	"status":true,
	 "data": [
	{
            "comment_id": 1,
            "goods_id": 2,
            "uid": 16,
            "text": "好耶",
            "star": 5,
            "date": "2022-02-20T18:58:16Z"
        }
    ]
}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"goods_id无效"   | |
|true |"暂无评论"   |参数合法  |
|true |"..."   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述






