package main

import (
	"github.com/pirogom/walk"
)

/**
*	TestListItem
**/
type TestListItem struct {
	Name    string
	Level   int
	Sex     int
	Class   string
	checked bool
}

/**
*	TestListModel
**/
type TestListModel struct {
	walk.TableModelBase
	items []TestListItem
}

/**
*	RowCount
**/
func (m *TestListModel) RowCount() int {
	return len(m.items)
}

/**
*	Value
**/
func (m *TestListModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Name
	case 1:
		return item.Level
	case 2:
		if item.Sex == 0 {
			return "중성"
		} else if item.Sex == 1 {
			return "남성"
		} else if item.Sex == 2 {
			return "여성"
		} else {
			return "알수없음"
		}
	case 3:
		return item.Class
	}
	panic("unexpected col")
}

/**
*	Checked
**/
func (m *TestListModel) Checked(row int) bool {
	return m.items[row].checked
}

/**
*
**/
func (m *TestListModel) CheckedCount() int {
	var cnt int
	for _, item := range m.items {
		if item.checked {
			cnt++
		}
	}

	return cnt
}

/**
*	SetChecked
**/
func (m *TestListModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

/**
*	ResetRows
**/
func (m *TestListModel) ResetRows() {
	m.items = nil
	m.PublishRowsReset()
}
