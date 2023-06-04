/*******
* @Author:qingmeng
* @Description:
* @File:sort
* @Date2022/2/20
 */

package dao

import "second-hand-trade/model"

type SortDao struct {
}

// ViewSort 查看所有类别
func (d *SortDao) ViewSort() ([]model.Sort, error) {
	var sorts []model.Sort
	rows, err := DB.Query("select * from second_hand_trade.sort ")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var sort model.Sort
		err = rows.Scan(&sort.SortId, &sort.SortName)
		if err != nil {
			return nil, err
		}
		sorts = append(sorts, sort)
	}
	return sorts, err
}

// AddSort 添加类别
func (d *SortDao) AddSort(sortName string) error {
	_, err := DB.Exec("insert into second_hand_trade.sort (sort_name)values (?)", sortName)
	return err
}

// SelectSortBySortName 根据sortName返回sort
func (d *SortDao) SelectSortBySortName(sortName string) (model.Sort, error) {
	var sort model.Sort
	row := DB.QueryRow("select * from second_hand_trade.sort where sort_name=? ", sortName)
	if row.Err() != nil {
		return sort, row.Err()
	}
	err := row.Scan(&sort.SortId, &sort.SortName)
	return sort, err
}

// SelectSortBySortId 根据sortId返回sort
func (d *SortDao) SelectSortBySortId(sortId int) (model.Sort, error) {
	var sort model.Sort
	row := DB.QueryRow("select * from second_hand_trade.sort where sort_id=? ", sortId)
	if row.Err() != nil {
		return sort, row.Err()
	}
	err := row.Scan(&sort.SortId, &sort.SortName)
	return sort, err
}
