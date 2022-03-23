package main

import (
	"fmt"
	"time"

	"github.com/pirogom/walk"
)

/**
*	FullWin
**/
func FullWin() {
	mgr, _ := NewWindowMgr("테스트 윈도/리사이즈/최대최소 가능", 1024, 768, GetIcon())

	// 최 상위로 실행
	// 창 위치 상관 없으면 mgr.StartWindow()
	mgr.StartForeground()
}

/**
*	NoResizeWin
**/
func NoResizeWin() {
	mgr, _ := NewWindowMgrNoResize("테스트 윈도/리사이즈 불가/최대최소가능", 1024, 768, GetIcon())

	mgr.StartForeground()
}

/**
*	NoResizeNoMinMxWin
**/
func NoResizeNoMinMxWin() {
	mgr, _ := NewWindowMgrNoResizeNoMinMax("테스트 윈도/리사이즈 불가/최대최소불가", 1024, 768, GetIcon())

	mgr.StartForeground()
}

/**
*	DefaultLayoutWin
**/
func DefaultLayoutWin() {
	mgr, window := NewWindowMgr("기본 레이아웃", 1025, 768, GetIcon())

	mgr.Label("이거슨 기본 라벨")
	mgr.LabelRight("이거슨 우측라벨")
	mgr.LabelCenter("이거슨 센터라벨")
	mgr.PushButton("버튼1", func() {
		MsgBox("버튼1을 누르심")
	})

	le1 := mgr.LineEdit(false)
	le1.TextChanged().Attach(func() {
		fmt.Println("라인에디트1 : " + le1.Text())
	})

	var cb1 *walk.CheckBox
	cb1 = mgr.CheckBox("이거슨 체크박스1", true, func() {
		if cb1.Checked() {
			MsgBox("체크박스1 체크됨")
		} else {
			MsgBox("체크박스1 체크안됨")
		}
	})

	// 윈도 시작됨
	window.Starting().Attach(func() {
		MsgBox("창시작")
	})
	mgr.StartForeground()
}

/**
*	CustomLayoutWin
**/
func CustomLayoutWin() {
	mgr, _ := NewWindowMgr("기본 레이아웃", 1025, 768, GetIcon())

	// HSplitter
	mgr.HSplit()
	mgr.Label("HSplitter")
	mgr.PushButton("버튼1", func() {})
	mgr.PushButton("버튼2", func() {})
	mgr.PushButton("버튼3", func() {})
	mgr.PushButton("버튼4", func() {})
	mgr.EndSplit() // HSplit, VSplit 사용후엔 EndSplit!!

	// VSplitter
	mgr.VSplit()
	mgr.Label("VSplitter")
	mgr.PushButton("버튼1", func() {})
	mgr.PushButton("버튼2", func() {})
	mgr.PushButton("버튼3", func() {})
	mgr.PushButton("버튼4", func() {})
	mgr.EndSplit()

	// HSplitter + VSplitter
	mgr.HSplit()

	mgr.VSplit()
	mgr.Label("HSplit 안에 VSplit")
	mgr.PushButton("버튼1", func() {})
	mgr.PushButton("버튼2", func() {})
	mgr.PushButton("버튼3", func() {})
	mgr.EndSplit() // End of VSplit

	mgr.HSplit()
	mgr.Label("HSplit 안에 HSplit")
	mgr.PushButton("버튼1", func() {})
	mgr.PushButton("버튼2", func() {})
	mgr.EndSplit()

	mgr.VSplit()
	mgr.Label("HSplit 안에 VSplit2")
	mgr.PushButton("버튼1", func() {})
	mgr.PushButton("버튼2", func() {})
	mgr.PushButton("버튼3", func() {})
	mgr.CheckBox("체크1", false, func() {})
	mgr.EndSplit()

	mgr.EndSplit() // End of HSplit

	// 그룹박스 Vertical
	mgr.GroupBox("그룹박스(V)", true)
	mgr.PushButton("버튼1", func() {})
	mgr.PushButton("버튼2", func() {})
	mgr.PushButton("버튼3", func() {})
	mgr.CheckBox("체크1", false, func() {})
	mgr.EndGroupBox()

	// 그룹박스 Horizen..
	mgr.GroupBox("그룹박스(H)", false)
	mgr.PushButton("버튼1", func() {})
	mgr.PushButton("버튼2", func() {})
	mgr.PushButton("버튼3", func() {})
	mgr.CheckBox("체크1", false, func() {})
	mgr.EndGroupBox()

	mgr.HSplit()

	mgr.GroupBox("HSplit으로 분리한 그룹박스1", true)
	mgr.PushButton("버튼1", func() {})
	mgr.PushButton("버튼2", func() {})
	mgr.PushButton("버튼3", func() {})
	mgr.CheckBox("체크1", false, func() {})
	mgr.EndGroupBox()

	mgr.GroupBox("HSplit으로 분리한 그룹박스2", true)
	mgr.PushButton("버튼1", func() {})
	mgr.PushButton("버튼2", func() {})
	mgr.PushButton("버튼3", func() {})
	mgr.CheckBox("체크1", false, func() {})
	mgr.EndGroupBox()

	mgr.EndSplit()

	mgr.StartForeground()
}

/**
*	WaitAndCloseWin
**/
func WaitAndCloseWin() {
	mgr, window := NewWindowMgrNoResizeNoMinMax("wait and close", 150, 100, GetIcon())

	label := mgr.Label("5초후 종료")

	ticker := time.NewTicker(time.Second)

	go func(tick *time.Ticker) {
		var count int = 5

		for {
			select {
			case <-tick.C:
				count--
				if count <= 0 {
					tick.Stop()
					mgr.HideAndClose()
					return
				} else {
					window.Synchronize(func() {
						label.SetText(fmt.Sprintf("%d초 남음", count))
					})
				}
			default:
				time.Sleep(time.Millisecond)
			}
		}
	}(ticker)

	mgr.DefClosing()
	mgr.StartForeground()
}

/**
*	TableViewWin
**/
func TableViewWin() {
	cbModel := new(testTableViewListModel)
	mgr, window := NewWindowMgrNoResizeNoMinMax("TableView 예제", 400, 300, GetIcon())

	th := []tableViewHeader{{Title: "이름", Width: 150}, {Title: "설명", Width: 150}}
	testTv := mgr.TableView(cbModel, th, true, true)

	// 마우스 클릭
	testTv.MouseDown().Attach(func(x int, y int, mouse walk.MouseButton) {
		if mouse == walk.LeftButton && x > 18 {
			nCurrItem := testTv.IndexAt(x, y)

			if nCurrItem > -1 {
				if cbModel.items[nCurrItem].checked {
					cbModel.items[nCurrItem].checked = false
				} else {
					cbModel.items[nCurrItem].checked = true
				}
				testTv.UpdateItem(nCurrItem)
			}
		}
	})

	testTv.KeyUp().Attach(func(key walk.Key) {
		if key == walk.KeyDelete {
			currIdx := testTv.CurrentIndex()

			if currIdx != -1 {
				if !Confirm(fmt.Sprintf("%s 를 삭제합니다.", cbModel.items[currIdx].Name)) {
					return
				}

				cbModel.items = append(cbModel.items[:currIdx], cbModel.items[currIdx+1:]...)
				cbModel.PublishRowsReset()
			}
		} else if key == walk.KeyReturn {
			od := testTableViewListItem{}
			od.Name = "피로곰"
			od.Desc = "만세"
			cbModel.items = append(cbModel.items, od)
			cbModel.PublishRowsReset()
		}
	})

	testTv.ItemActivated().Attach(func() {
		currIdx := testTv.CurrentIndex()

		if currIdx != -1 {
			MsgBox(fmt.Sprintf("선택된 아이템 : %d", currIdx))
		}
	})

	mgr.PushButton("선택된 아이템", func() {
		selectedIndex := testTv.SelectedIndexes()

		if len(selectedIndex) == 0 {
			MsgBox("없음")
		} else {
			MsgBox(fmt.Sprintf("%v", selectedIndex))
		}
	})

	mgr.PushButton("전체체크", func() {
		for idx, _ := range cbModel.items {
			cbModel.items[idx].checked = true
			testTv.UpdateItem(idx)
		}
	})

	mgr.PushButton("전체체크해제", func() {
		for idx, _ := range cbModel.items {
			cbModel.items[idx].checked = false
			testTv.UpdateItem(idx)
		}
	})

	window.Starting().Attach(func() {
		ad := []testTableViewListItem{
			{Name: "피로곰", Desc: "만세"},
			{Name: "Panic", Desc: "환자"},
			{Name: "용민", Desc: "환자2"},
		}

		cbModel.items = append(cbModel.items, ad...)
		cbModel.PublishRowsReset()
	})

	mgr.StartForeground()
}

/**
*	ImageViewWin
**/
func ImageViewWin() {
	var currImage *walk.Image

	defer func() {
		// 새 이미지 로드전에 기존 이미지 해제 않하면
		// 메모리릭 발생 !! 주의
		if currImage != nil {
			(*currImage).Dispose()
		}
	}()

	mgr, window := NewWindowMgrNoResizeNoMinMax("TableView 예제", 1920, 1040, GetIcon())

	mgr.HSplit()

	idealIV := mgr.ImageView()
	idealIV.SetMode(walk.ImageViewModeIdeal)

	cornerIV := mgr.ImageView()
	cornerIV.SetMode(walk.ImageViewModeCorner)

	centerIV := mgr.ImageView()
	centerIV.SetMode(walk.ImageViewModeCenter)

	mgr.EndSplit()

	mgr.HSplit()
	shrinkIV := mgr.ImageView()
	shrinkIV.SetMode(walk.ImageViewModeShrink)

	zoomIV := mgr.ImageView()
	zoomIV.SetMode(walk.ImageViewModeZoom)

	stretchIV := mgr.ImageView()
	stretchIV.SetMode(walk.ImageViewModeStretch)
	mgr.EndSplit()

	updateImg := func(fname string) {
		// 새 이미지 로드전에 기존 이미지 해제 않하면
		// 메모리릭 발생 !! 주의
		if currImage != nil {
			(*currImage).Dispose()
		}
		currImage = LoadImage(fname)
		if currImage != nil {
			idealIV.SetImage(*currImage)
			cornerIV.SetImage(*currImage)
			centerIV.SetImage(*currImage)
			shrinkIV.SetImage(*currImage)
			zoomIV.SetImage(*currImage)
			stretchIV.SetImage(*currImage)
		}
	}

	mgr.PushButton("피로곰", func() {
		updateImg("img0.png")
	})
	mgr.PushButton("환자1", func() {
		updateImg("img2.png")
	})
	mgr.PushButton("환자2", func() {
		updateImg("img1.png")
	})

	//
	window.Starting().Attach(func() {
	})

	mgr.StartForeground()
}

func test1() {
	mgr, _ := NewWindowMgrNoResize("레이아웃 테스트", 640, 1, GetIcon())

	mgr.HSplit()
	mgr.Label("1-1. 라벨")
	le1 := mgr.LineEdit(false)
	le1.SetText("1-2. 에디트 박스")
	mgr.PushButton("1-3. 버튼", func() {})
	mgr.EndSplit()

	mgr.HSplit()
	mgr.Label("2-1. 라벨")
	mgr.DropDownBox([]string{"1.하나", "2.둘", "3.셋"})
	mgr.PushButton("2-3. 버튼", func() {})
	mgr.EndSplit()

	mgr.StartForeground()
}

func test2() {
	mgr, _ := NewWindowMgrNoResize("레이아웃 테스트2", 640, 480, GetIcon())

	mgr.HSplit()

	mgr.VSplit()
	mgr.Label("HSplit 안에 Vsplit1")
	mgr.Label("HSplit 안에 Vsplit2")
	mgr.Label("HSplit 안에 Vsplit3")
	mgr.PushButton("버튼1", func() {})
	mgr.EndSplit()

	mgr.TextArea(false)

	mgr.VSplit()
	mgr.Label("HSplit 안에 Vsplit2-1")
	mgr.Label("HSplit 안에 Vsplit2-2")
	mgr.Label("HSplit 안에 Vsplit2-3")
	mgr.Label("HSplit 안에 Vsplit2-4")
	mgr.PushButton("버튼2", func() {})
	mgr.EndSplit()

	mgr.EndSplit()

	mgr.StartForeground()
}

func LabelTest() {
	mgr, _ := NewWindowMgrNoResize("라벨테스트", 640, 480, GetIcon())

	mgr.Label("Label함수")
	mgr.LabelLeft("LabelLeft함수")
	mgr.LabelCenter("LabelCenter함수")
	mgr.LabelRight("LabelRight함수")

	mgr.StartForeground()
}

func LabelTest2() {
	mgr, _ := NewWindowMgrNoResize("라벨테스트", 640, 480, GetIcon())

	label1 := mgr.Label("Label함수")
	mgr.PushButton("Label함수 라벨 변경", func() {
		label1.SetText("버튼1 클릭!!")
	})
	label2 := mgr.LabelLeft("LabelLeft함수")
	mgr.PushButton("LabelLeft함수 라벨 변경", func() {
		label2.SetText("버튼2 클릭!!")
	})
	label3 := mgr.LabelCenter("LabelCenter함수")
	mgr.PushButton("LabelCenter함수 라벨 변경", func() {
		label3.SetText("버튼3 클릭!!")
	})
	label4 := mgr.LabelRight("LabelRight함수")
	mgr.PushButton("LabelRight함수 라벨 변경", func() {
		label4.SetText("버튼4 클릭!!")
	})
	mgr.StartForeground()
}

func EditTest() {
	mgr, _ := NewWindowMgrNoResize("에디트박스 테스트", 640, 480, GetIcon())

	mgr.HSplit()
	mgr.Label("NumberEdit")
	ne := mgr.NumberEdit()
	ne.SetValue(100.0)
	mgr.EndSplit()

	mgr.HSplit()
	mgr.Label("LineEdit(read only)")
	lero := mgr.LineEdit(true)
	lero.SetText("읽기전용 LineEdit")
	mgr.EndSplit()

	mgr.HSplit()
	mgr.Label("LineEdit")
	le := mgr.LineEdit(false)
	le.SetText("수정가능 LineEdit")
	mgr.EndSplit()

	mgr.HSplit()
	mgr.Label("TextEdit(read only)")
	tero := mgr.TextEdit(true)
	tero.SetText("읽기전용 TextEdit")
	mgr.EndSplit()

	mgr.HSplit()
	mgr.Label("TextEdit")
	te := mgr.TextEdit(false)
	te.SetText("수정가능 TextEdit")
	mgr.EndSplit()

	mgr.HSplit()
	mgr.Label("TextArea(read only)")
	taro := mgr.TextArea(true)
	taro.SetText("읽기전용 TextArea")
	mgr.EndSplit()

	mgr.HSplit()
	mgr.Label("TextArea")
	ta := mgr.TextArea(false)
	ta.SetText("수정가능 TextArea")
	mgr.EndSplit()

	mgr.StartForeground()
}

func EditTest2() {
	mgr, _ := NewWindowMgrNoResize("에디트박스 테스트", 640, 100, GetIcon())

	mgr.HSplit()
	mgr.Label("NumberEdit")
	ne := mgr.NumberEdit()
	ne.SetValue(100.0)
	mgr.EndSplit()

	mgr.HSplit()
	mgr.PushButton("NumberEdit 변경", func() {
		ne.SetValue(float64(time.Now().Unix()))
	})
	mgr.PushButton("NumberEdit 값", func() {
		MsgBox(fmt.Sprintf("%d", int(ne.Value())))
	})
	mgr.EndSplit()

	mgr.HSplit()
	mgr.Label("LineEdit")
	le := mgr.LineEdit(false)
	le.SetText("수정가능 LineEdit")
	mgr.EndSplit()

	mgr.HSplit()
	mgr.PushButton("LineEdit 변경", func() {
		le.SetText(fmt.Sprintf("현재 유닉스 타임스탬프: %d", time.Now().Unix()))
	})
	mgr.PushButton("LineEdit 값", func() {
		MsgBox(le.Text())
	})
	mgr.EndSplit()

	mgr.StartForeground()
}

func EditTest3() {
	mgr, _ := NewWindowMgrNoResize("에디트박스 테스트", 640, 100, GetIcon())

	mgr.HSplit()
	mgr.Label("NumberEdit")
	ne := mgr.NumberEdit()
	ne.SetValue(100.0)
	mgr.EndSplit()

	ne.ValueChanged().Attach(func() {
		fmt.Printf("값 변경됨: %d\n", int(ne.Value()))
	})

	mgr.HSplit()
	mgr.Label("LineEdit")
	le := mgr.LineEdit(false)
	le.SetText("LineEdit")
	mgr.EndSplit()

	le.TextChanged().Attach(func() {
		fmt.Println("텍스트 변경됨:", le.Text())
	})

	le.FocusedChanged().Attach(func() {
		if le.Focused() {
			fmt.Println("LineEdit 포커스 됨")
		} else {
			fmt.Println("LineEdit 포커스 해제됨")
		}
	})

	mgr.StartForeground()
}

func checkBoxTest1() {
	mgr, _ := NewWindowMgrNoResize("체크박스 테스트1", 640, 200, GetIcon())

	mgr.CheckBox("checked false 체크박스", false, func() {})
	mgr.CheckBox("checked true 체크박스", true, func() {})

	mgr.StartForeground()
}

func checkBoxTest2() {
	mgr, _ := NewWindowMgrNoResize("체크박스 테스트2", 640, 200, GetIcon())

	var check1 *walk.CheckBox
	check1 = mgr.CheckBox("체크박스 테스트", false, func() {

		if check1.Checked() {
			MsgBox("체크됨")
		} else {
			MsgBox("해제됨")
		}
	})

	mgr.StartForeground()
}

func checkBoxTest3() {
	mgr, window := NewWindowMgrNoResize("체크박스 테스트3", 640, 200, GetIcon())

	var check1 *walk.CheckBox
	check1 = mgr.CheckBox("체크박스 테스트", false, func() {

		if check1.Checked() {
			MsgBox("체크됨", window)
		} else {
			MsgBox("해제됨", window)
		}
	})

	mgr.StartForeground()
}

func checkBoxTest4() {
	mgr, _ := NewWindowMgrNoResize("체크박스 테스트4", 640, 200, GetIcon())

	mgr.HSplit()
	mgr.CheckBox("checked false 체크박스", false, func() {})
	mgr.CheckBox("checked true 체크박스", true, func() {})
	//mgr.PushButton("테스트", func() {})
	mgr.EndSplit()

	mgr.StartForeground()
}
func checkBoxTest5() {
	mgr, _ := NewWindowMgrNoResize("체크박스 테스트5", 640, 200, GetIcon())

	// mgr.HSplit()
	// mgr.Label("GroupBox 안의 VSplit")
	// mgr.PushButton("VSpilit 버튼", func() {})

	mgr.GroupBox("체크박스", false)
	mgr.CheckBox("checked false 체크박스", false, func() {})
	mgr.CheckBox("checked true 체크박스", true, func() {})
	mgr.EndGroupBox()

	// mgr.EndSplit()

	mgr.StartForeground()
}

/**
*	comboTest1
**/
func comboTest1() {
	mgr, _ := NewWindowMgrNoResize("콤보박스 테스트", 640, 200, GetIcon())

	dd1 := mgr.DropDownBox([]string{"1.하하", "2.호호", "3.ㅋㅋ", "4.ㅎㅎ"})

	dd1.SetCurrentIndex(3)

	dd1.CurrentIndexChanged().Attach(func() {
		fmt.Printf("선택된 인덱스 : %d, 선택된 값 : %s\n", dd1.CurrentIndex(), dd1.Text())
	})

	mgr.StartForeground()
}

func webviewTest() {
	mgr, _ := NewWindowMgr("웹뷰 테스트", 1024, 680, GetIcon())

	mgr.WebView("https://modu-print.tistory.com/category/%EB%8B%A4%EC%9A%B4%EB%A1%9C%EB%93%9C/%EC%97%85%EB%8D%B0%EC%9D%B4%ED%8A%B8%20%EB%82%B4%EC%97%AD%EC%95%88%EB%82%B4")

	mgr.StartForeground()
}

func webviewTest2() {
	var wv *walk.WebView
	mgr, _ := NewWindowMgr("웹뷰 테스트2", 1024, 680, GetIcon())

	mgr.HSplit()
	mgr.Label("URL:")
	urlEdit := mgr.LineEdit(false)
	mgr.PushButton("GO!", func() {
		wv.SetURL(urlEdit.Text())
	})
	mgr.EndSplit()

	wv = mgr.WebView("https://modu-print.tistory.com/category/%EB%8B%A4%EC%9A%B4%EB%A1%9C%EB%93%9C/%EC%97%85%EB%8D%B0%EC%9D%B4%ED%8A%B8%20%EB%82%B4%EC%97%AD%EC%95%88%EB%82%B4")
	wv.SetNativeContextMenuEnabled(true)

	mgr.StartForeground()
}

/**
*	main
**/
func main() {

	// embed 된 ico 파일을 쓰고 싶으면 LoadIcon 함수 사용
	LoadIconFromFile("./test.ico")

	//ImageViewWin()
	//EditTest2()
	//webviewTest2()
	//	comboTest1()
	//LabelTest2()
	WaitAndCloseWin()
	/*	MsgBox("메시지 박스임 ㅋㅋㅋ")


		ImageViewWin()
		TableViewWin()

		FullWin()
		NoResizeWin()
		NoResizeNoMinMxWin()
		DefaultLayoutWin()
		CustomLayoutWin()
		WaitAndCloseWin()*/
}
