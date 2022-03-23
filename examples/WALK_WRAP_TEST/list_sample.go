package main

import "fmt"

/**
*	listTest1
**/
func listTest1() {
	mgr, window := NewWindowMgrNoResize("리스트1", 640, 500, GetIcon())

	cbModel := new(TestListModel)

	testTv := mgr.TableView(cbModel, []tableViewHeader{
		{Title: "이름", Width: 100},
		{Title: "레벨", Width: 100},
		{Title: "성별", Width: 100},
		{Title: "직업", Width: 100},
	}, false, false)

	testTv.ItemActivated().Attach(func() {
		currIdx := testTv.CurrentIndex()

		if currIdx < 0 {
			return
		}

		fmt.Println("= 더블클릭 ========================")
		fmt.Println("선택된 아이템:", currIdx)
		fmt.Println("이름:", cbModel.items[currIdx].Name)
		fmt.Println("레벨:", cbModel.items[currIdx].Level)
		fmt.Println("성별:", cbModel.items[currIdx].Sex)
		fmt.Println("직업:", cbModel.items[currIdx].Class)
	})

	window.Starting().Attach(func() {
		for i := 0; i < 10; i++ {
			od := TestListItem{}
			od.Name = fmt.Sprintf("사용자%02d", i)
			od.Level = i + 1
			od.Sex = (i % 2)
			od.Class = fmt.Sprintf("직업%d", i+1)
			cbModel.items = append(cbModel.items, od)
		}
		cbModel.PublishRowsReset()
	})

	mgr.StartForeground()
}

/**
*	listTest2
**/
func listTest2() {
	mgr, window := NewWindowMgrNoResize("리스트2(CheckBox)", 640, 500, GetIcon())

	cbModel := new(TestListModel)

	testTv := mgr.TableView(cbModel, []tableViewHeader{
		{Title: "이름", Width: 100},
		{Title: "레벨", Width: 100},
		{Title: "성별", Width: 100},
		{Title: "직업", Width: 100},
	}, true, false)

	mgr.PushButton("체크된 아이템", func() {
		for i := 0; i < len(cbModel.items); i++ {
			if cbModel.Checked(i) {
				fmt.Println(i, "체크됨", cbModel.items[i].Name)
			}
		}
	})

	mgr.PushButton("체크된 아이템 이름변경", func() {
		for i := 0; i < len(cbModel.items); i++ {
			if cbModel.Checked(i) {
				cbModel.items[i].Name = fmt.Sprintf("변경된이름%d", i)
				cbModel.PublishRowChanged(i)
			}
		}
	})

	testTv.ItemActivated().Attach(func() {
		currIdx := testTv.CurrentIndex()

		if currIdx < 0 {
			return
		}

		if cbModel.Checked(currIdx) {
			cbModel.SetChecked(currIdx, false)
		} else {
			cbModel.SetChecked(currIdx, true)
		}
		cbModel.PublishRowChanged(currIdx)
	})

	window.Starting().Attach(func() {
		for i := 0; i < 10; i++ {
			od := TestListItem{}
			od.Name = fmt.Sprintf("사용자%02d", i)
			od.Level = i + 1
			od.Sex = (i % 2)
			od.Class = fmt.Sprintf("직업%d", i+1)
			cbModel.items = append(cbModel.items, od)
		}
		cbModel.PublishRowsReset()
	})

	mgr.StartForeground()
}

/**
*	listTest3
**/
func listTest3() {
	mgr, window := NewWindowMgrNoResize("리스트2(다중선택)", 640, 500, GetIcon())

	cbModel := new(TestListModel)

	testTv := mgr.TableView(cbModel, []tableViewHeader{
		{Title: "이름", Width: 100},
		{Title: "레벨", Width: 100},
		{Title: "성별", Width: 100},
		{Title: "직업", Width: 100},
	}, false, true)

	mgr.PushButton("선택된 아이템", func() {
		idxs := testTv.SelectedIndexes()

		if len(idxs) == 0 {
			return
		}

		for _, i := range idxs {
			fmt.Println(i, "선택된", cbModel.items[i].Name)
		}

	})

	mgr.PushButton("선택된 아이템 이름변경", func() {

		idxs := testTv.SelectedIndexes()

		if len(idxs) == 0 {
			return
		}

		for _, i := range idxs {
			cbModel.items[i].Name = fmt.Sprintf("변경된이름%d", i)
			cbModel.PublishRowChanged(i)
		}
	})

	window.Starting().Attach(func() {
		for i := 0; i < 10; i++ {
			od := TestListItem{}
			od.Name = fmt.Sprintf("사용자%02d", i)
			od.Level = i + 1
			od.Sex = (i % 2)
			od.Class = fmt.Sprintf("직업%d", i+1)
			cbModel.items = append(cbModel.items, od)
		}
		cbModel.PublishRowsReset()
	})

	mgr.StartForeground()
}
