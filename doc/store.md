[创建店铺]

##### 简要描述

- multipart/form-data
- 创建店铺接口,成功则该用户自动加入该店铺

##### 请求URL

- ` http://101.43.160.254：8080/api/store`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|store_name |是  |String |店铺名字   |
|notice |是  |String |店铺公告   |

##### 返回示例

``` 
{
		"status": true,
		"data": "创建成功",
		"store_id":  1,
		"store_name":饮品店,
		"notice":物美价廉
	}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"该用户不是商家"   | |
|false |"该用户已有店铺"   | |
|false |"notice无效"   | |
|false |"store_name无效"   | |
|false |"该店铺名已被注册"   | |
|true |"创建成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[更新店铺]

##### 简要描述

- multipart/form-data
- 更新店铺接口

##### 请求URL

- ` http://101.43.160.254：8080/api/store`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|store_name |否  |String |店铺名字   |
|notice |否  |String |店铺公告   |

##### 返回示例

``` 
{
		"status": true,
		"data": "修改成功",
		"store_id":
		"store_name":
		"notice":
	}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"该用户不是商家"   | |
|false |"notice无效"   | |
|false |"该店铺名已被注册"   | |
|false |"store_name无效"   | |
|true |"修改成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[查看店铺]

##### 简要描述

- multipart/form-data
- 查看店铺接口

##### 请求URL

- ` http://101.43.160.254：8080/api/store`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|store_name |否  |String |店铺名字  (不可查看店铺资金) |
|store_id |否  |String |店铺id  (可查看店铺资金) |

##### 返回示例

``` 
{
    "status": true,
    "data": "查看成功",
    "notice": "暂无公告",
    "store_id": 1,
    "store_name": "汽水"
    "money": 500,
}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"参数无效"   | store_id或store_name无效|
|true |"查看成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[店铺充值]

##### 简要描述

- multipart/form-data
- 店铺充值接口

##### 请求URL

- ` http://101.43.160.254：8080/api/store/money`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|store_id |是  |String |店铺id   |
|money |是  |String |充值金额   |

##### 返回示例

``` 
{
		"status": true,
		"data": "成功充值100元",
	}
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"store_id无效"   | |
|false |"money无效"   | |
|true |"用户余额不足"   | |
|true |"成功充值x元"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

