/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date2021/12/10
 */

package main

import (
	"second-hand-trade/api"
	"second-hand-trade/dao"
)

func main() {
	dao.InitDB()
	api.InitEngine()
}
