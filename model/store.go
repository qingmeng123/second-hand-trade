/*******
* @Author:qingmeng
* @Description:
* @File:store
* @Date2022/2/19
 */

package model

//店铺
type Store struct {

	//关联必须加上primary_key?
	StoreId    int     `json:"store_id" gorm:"primary_key" ` //店铺ID
	StoreName  string  `json:"store_name"`
	Notice     string  `json:"notice,omitempty" ` //店铺公告
	StoreMoney float64 `json:"store_money"`       //店铺资金

}

func (Store) TableName() string {
	return "store"
}
