/*******
* @Author:qingmeng
* @Description:
* @File:sort
* @Date2022/2/19
 */

package tool

import "second-hand-trade/model"

// 通过商品，商品属性，排序方式对商品排序 desc:true为降序，false为升序,其他不排序
func SortGoodsByNature(goodses []model.Goods, nature string, desc string) []model.Goods {
	//降序
	if desc == "true" {
		if nature == "price" {
			for i := 0; i < len(goodses); i++ {
				for j := 0; j < len(goodses)-1; j++ {
					if goodses[j].Price < goodses[j+1].Price {
						goodses[j], goodses[i] = goodses[i], goodses[j]
					}
				}
			}
		}
		if nature == "turnover" {
			for i := 0; i < len(goodses); i++ {
				for j := 0; j < len(goodses)-1; j++ {
					if goodses[j].Turnover < goodses[j+1].Turnover {
						goodses[j], goodses[i] = goodses[i], goodses[j]
					}
				}
			}
		}

	} else if desc == "false" {
		if nature == "price" {
			for i := 0; i < len(goodses); i++ {
				for j := 0; j < len(goodses)-1; j++ {
					if goodses[j].Price > goodses[j+1].Price {
						goodses[j], goodses[i] = goodses[i], goodses[j]
					}
				}
			}
		}
		if nature == "turnover" {
			for i := 0; i < len(goodses); i++ {
				for j := 0; j < len(goodses)-1; j++ {
					if goodses[j].Turnover < goodses[j+1].Turnover {
						goodses[j], goodses[i] = goodses[i], goodses[j]
					}
				}
			}
		}
	} else {
		return goodses
	}
	return goodses
}
