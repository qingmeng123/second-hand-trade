/*******
* @Author:qingmeng
* @Description:
* @File:sort
* @Date2022/2/20
 */

package service

import (
	"database/sql"
	"second-hand-trade/dao"
	"second-hand-trade/model"
)

type SortService struct {
}

func (s *SortService) ViewSort() ([]model.Sort, error) {
	d := dao.SortDao{}
	return d.ViewSort()
}

func (s *SortService) AddSort(sortName string) error {
	d := dao.SortDao{}
	return d.AddSort(sortName)
}

func (s *SortService) IsExistSortName(sortName string) (bool, error) {
	d := dao.SortDao{}
	_, err := d.SelectSortBySortName(sortName)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s SortService) SelectSortBySortId(sortId int) (model.Sort, error) {
	d := dao.SortDao{}
	return d.SelectSortBySortId(sortId)
}

func (s *SortService) IsExistSortId(sortId int) (bool, error) {
	d := dao.SortDao{}
	_, err := d.SelectSortBySortId(sortId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
