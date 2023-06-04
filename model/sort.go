/*******
* @Author:qingmeng
* @Description:
* @File:sort
* @Date2022/2/20
 */

package model

type Sort struct {
	SortId   int    `json:"sort_id"`
	SortName string `json:"sort_name"`
}

func (Sort) TableName() string {
	return "sort"
}
