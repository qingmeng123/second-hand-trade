[用户注册]

##### 简要描述

- multipart/form-data
- 用户注册接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|username |是  |string |用户名   |
|password |是  |string | 密码    |
|name     |否  |string | 昵称    |

##### 返回示例

``` 
  {
    "status": true,
    "data": "注册成功!"
    "username": "12154545",
    "name": "清梦",
    "groupId": 0
    
  }
```

##### 返回参数说明

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|groupId |int   |用户组id，1：店家；0：普通用户  |

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"用户名不能为空"   |username 为空  |
|false |"用户名太长了"   |username 长度超过 20 个字节  |
|false |"昵称太长了"   |name 长度超过 20 个字节  |
|false |"密码不能小于6个字符"   |password 长度少于 6 个字节  |
|false |"密码不能大于20个字符"   |password 长度超过 20 个字节  |
|true |"注册成功！"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[用户登陆]

##### 简要描述

- multipart/form-data
- 用户登陆接口，检查用户名及密码

##### 请求URL

- ` http://101.43.160.254：8080/api/user`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|username |是  |string |用户名   |
|password |是  |string | 密码    |

##### 返回示例

``` 
   {
    "status": true,
    "data": "登陆成功!"
	"uid":
	"groupId":
	"token":
	"refreshToken":
  }
```

##### 返回参数说明

|参数名|类型|说明| 
|:----- |:-----|----- | 
|groupId |int|用户组id，1：店家；0：普通用户 | 
|token | string |用户token |
|refreshToken |string |用户refreshToken |
|uid | int |用户uid |

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"用户名或密码错误"   |username 与 password 不匹配 |
|true |"注册成功！"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[用户充值]

##### 简要描述

- multipart/form-data
- 用户充值接口，默认给自己充值

##### 请求URL

- ` http://101.43.160.254：8080/api/user/money`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|username |否  |String |给指定用户名的用户充值   |

##### 返回示例

``` 

{
"status": true,
"data": {
"给用户清梦成功充值1元!"
} }

```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"username无效"   |不存在该username或username格式错误 |
|false |"充值金额错误"   |参数不合法 |
|true |"给用户清梦成功值1元!"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[用户个人信息]

##### 简要描述

- multipart/form-data
- 用户个人信息接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/info`

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
        "AddressId": 0,
        "Gender": "0",
        "GroupId": 0,
        "Money": 10100,
        "Name": "匿名用户",
        "Password": "123456",
        "Phone": "0",
        "StoreId": 0,
        "Uid": 5,
        "Username": "qm"
    },

```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|true |"..."   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[用户余额]

##### 简要描述

- multipart/form-data
- 查看用户个人余额接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/money`

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
"data": "1000"

```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|true |"1000"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[登陆后修改密码]

##### 简要描述

- multipart/form-data
- 登陆后修改密码接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/password`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|old_password |是  |string |用户旧密码   |
|new_password |是  |String |用户新密码   |

##### 返回示例

``` 

{
"status":true,
"data": 修改成功！
}

```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"旧密码错误"   | |
|false |"新密码无效"   | |
|true |"修改成功！"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[用户个人信息修改]

##### 简要描述

- multipart/form-data
- 用户个人信息修改接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/info`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|address_id |否  |int |用户默认地址id修改   |
|name |否 |String |用户昵称修改   |
|gender |否 |String |用户性别修改 (0为男，1为女)  |
|phone |否  |String |用户电话修改   |

##### 返回示例

``` 

{
"status":true,
"data": 修改成功！
}

```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"address_id无效"   |address_id为空或无效 |
|false |"phone无效"   | |
|false |"地址重复"   |存在相同的所有参数 |
|false |"name格式不符合要求"   | |
|true |"修改成功！"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[用户收货地址增加]

##### 简要描述

- multipart/form-data
- 用户收货地址增加接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/info/address`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|name |是  |String |收货用户昵称   |
|phone |是  |String |收货用户电话  |
|address |是  |String |收货地址    |

##### 返回示例

``` 

{
"status":true,
"data": "添加成功！"
}

```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"收货昵称不合法"   |  |
|false |"收货电话不合法"   |  |
|false |"收货地址不合法"   |  |
|true |"增加成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[用户所有收货地址]

##### 简要描述

- multipart/form-data
- 用户所有收货地址接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/info/address`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |

##### 返回示例

``` 

{
    "satues": true
    "data": "查询成功",
    "addresses": [
        {
            "address_id": 2,
            "uid": 16,
            "name": "清梦",
            "phone": "13637833441",
            "address": "重庆市"
        },
        {
            "address_id": 3,
            "uid": 16,
            "name": "清梦",
            "phone": "13637833441",
            "address": "重庆市永川区"
        },
        {
            "address_id": 4,
            "uid": 16,
            "name": "清梦",
            "phone": "13637833441",
            "address": "重庆市永川区123"
        }
    ],

}

```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"无保存地址，请增加地址"   | |
|true |"查询成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[用户收货地址修改]

##### 简要描述

- multipart/form-data
- 用户收货地址修改接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/info/address`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|address_id |是  |int |用户收货地址id   |
|name |否  |String |收货用户昵称   |
|phone |否  |String |收货用户电话  |
|address |否  |String |收货地址    |

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
|false |"address_id无效"   |address_id为空或无效 |
|false |"收货昵称不合法"   |  |
|false |"收货电话不合法"   |  |
|false |"收货地址不合法"   |  |
|true |"修改成功"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[注册商家]

##### 简要描述

- multipart/form-data
- 注册商家接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/power`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |string |用户token   |

##### 返回示例

``` 
  {
    "status": true,
    "data": {
	"注册成功!"
      "uid": "1",
      "username": "12154545",
      "name": "清梦",
      "groupid": 1
    }
  }
```

##### 返回参数说明

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|groupid |int   |用户组id，1：店家；0：普通用户  |

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"该用户已是店家"   | |
|true |"注册成功！"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[商家入铺]

##### 简要描述

- multipart/form-data
- 商家入铺接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/store`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |string |用户token   |
|store_id |是  |int  |需要加入的店铺id   |

##### 返回示例

``` 
  {
    "status": true,
    "data": "加入x号店铺成功！"
  }
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"store_id无效"   | |
|false |"该用户不是商家"   | |
|false |"该用户已有店铺"   | |
|true |"加入x号店铺成功！"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[关注商品或店铺]

##### 简要描述

- multipart/form-data
- 关注商品或店铺接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/focus`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |string |用户token   |
|goods_id |否  |int  |需要关注的商品id   |
|store_id |否  |int  |需要关注的店铺id   |

##### 返回示例

``` 
  {
    "status": true,
    "data": "关注成功！"
  }
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"未输出参数"   | goods_id和store_id都为空|
|false |"goods_id无效"   | |
|false |"已关注过该商品"   | |
|false |"store_id无效"   | |
|false |"已关注过该店铺"   | |
|true |"关注成功！"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[查看关注]

##### 简要描述

- multipart/form-data
- 查看关注商品接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/focus`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |string |用户token   |

##### 返回示例

``` 
  {
    "status": true,
    "data": [
        {
            "uid": 16,
            "goods_id": 2,
            "store_id": 0
        }
    ],
  }
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|true |"暂无关注"   | |
|true |"..."   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

[取消关注商品或店铺]

##### 简要描述

- multipart/form-data
- 取消关注商品或店铺接口

##### 请求URL

- ` http://101.43.160.254：8080/api/user/focus`

##### 请求方式

- DELETE

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |string |用户token   |
|goods_id |否  |int  |需要关注的商品id   |
|store_id |否  |int  |需要关注的店铺id   |

##### 返回示例

``` 
  {
    "status": true,
    "data": "取消关注成功！"
  }
```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"未输出参数"   | goods_id和store_id都为空|
|false |"goods_id无效"   | |
|false |"未关注过该商品"   | |
|false |"store_id无效"   | |
|false |"未关注过该店铺"   | |
|true |"取消关注成功！"   |参数合法  |

##### 备注

- 更多返回错误代码请看首页的错误代码描述

