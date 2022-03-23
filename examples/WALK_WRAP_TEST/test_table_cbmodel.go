package main

import (
	"github.com/pirogom/walk"
)

/**
*	testTableViewListItem
**/
type testTableViewListItem struct {
	Name    string
	Desc    string
	checked bool // 기본으로 두소
}

/**
*	testTableViewListModel
**/
type testTableViewListModel struct {
	walk.TableModelBase
	items []testTableViewListItem
}

// 이 아랫것들은 그냥 기본으로 존재해야 한다 생각 하소 ..
// 귀찮으면 알아서 인터페이스 구조체 만들든가 말든가..
/**
*	RowCount
**/
func (m *testTableViewListModel) RowCount() int {
	return len(m.items)
}

/**
*	Value
**/
func (m *testTableViewListModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Name
	case 1:
		return item.Desc
	}
	panic("unexpected col")
}

/**
*	Checked
**/
func (m *testTableViewListModel) Checked(row int) bool {
	return m.items[row].checked
}

/**
*	SetChecked
**/
func (m *testTableViewListModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

/**
*	ResetRows
**/
func (m *testTableViewListModel) ResetRows() {
	m.items = nil
	m.PublishRowsReset()
}
