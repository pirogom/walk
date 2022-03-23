package main

import (
	"container/list"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/pirogom/walk"
	"github.com/pirogom/win"
)

const (
	walkWinName = "WALK_WRAP"
)

/**
*	tableViewHeader
**/
type tableViewHeader struct {
	Title string
	Width int
	Align string
}

/**
*	WinResMgr
**/
type WinResMgr struct {
	window     *walk.MainWindow
	parentList *list.List
}

var (
	gIcon     *walk.Icon
	iconMutex sync.Mutex
)

/**
*	LoadIcon
**/
func LoadIcon(icoBuf []byte, icoName string) {
	icoFile := filepath.Join(os.TempDir(), icoName)

	var err error

	if _, err = os.Stat(icoName); os.IsNotExist(err) {
		if err = ioutil.WriteFile(icoFile, icoBuf, 0644); err != nil {
			return
		}
	}

	iconMutex.Lock()
	gIcon, _ = walk.NewIconFromFile(icoFile)
	iconMutex.Unlock()
}

/**
*	LoadIconFromFile
**/
func LoadIconFromFile(icoPath string) {
	iconMutex.Lock()
	defer iconMutex.Unlock()
	gIcon, _ = walk.NewIconFromFile(icoPath)
}

/**
*	GetIcon
**/
func GetIcon() *walk.Icon {
	iconMutex.Lock()
	defer iconMutex.Unlock()

	return gIcon
}

/**
*	MsgBox
**/
func MsgBox(msg string, window ...*walk.MainWindow) {
	if len(window) > 0 {
		walk.MsgBox(window[0], "알림", msg, walk.MsgBoxOK|walk.MsgBoxSetForeground)
	} else {
		walk.MsgBox(nil, "알림", msg, walk.MsgBoxOK|walk.MsgBoxSetForeground)
	}
}

/**
*	Confirm
**/
func Confirm(msg string, window ...*walk.MainWindow) bool {
	if len(window) > 0 {
		if walk.MsgBox(window[0], "알림", msg, walk.MsgBoxYesNo|walk.MsgBoxSetForeground) == win.IDNO {
			return false
		} else {
			return true
		}
	} else {
		if walk.MsgBox(nil, "알림", msg, walk.MsgBoxYesNo|walk.MsgBoxSetForeground) == win.IDNO {
			return false
		} else {
			return true
		}
	}
	return false
}

/**
*	NewWindowMgr
**/
func NewWindowMgr(title string, width int, height int, icon *walk.Icon) (*WinResMgr, *walk.MainWindow) {
	rd := WinResMgr{}
	win, _ := walk.NewMainWindowWithName(walkWinName)
	rd.window = win
	rd.parentList = list.New()

	win.SetTitle(title)
	win.SetLayout(walk.NewVBoxLayout())
	win.SetWidth(width)
	win.SetHeight(height)

	if icon != nil {
		win.SetIcon(icon)
	}
	return &rd, win
}

/**
*	NewWindowMgrNoResize
**/
func NewWindowMgrNoResize(title string, width int, height int, icon *walk.Icon) (*WinResMgr, *walk.MainWindow) {
	rd := WinResMgr{}
	win, _ := walk.NewMainWindowWithName(walkWinName)
	rd.window = win
	rd.parentList = list.New()

	win.SetTitle(title)
	win.SetLayout(walk.NewVBoxLayout())
	win.SetWidth(width)
	win.SetHeight(height)

	if icon != nil {
		win.SetIcon(icon)
	}

	win.SetMinMaxSize(walk.Size{Width: width, Height: height}, walk.Size{Width: width, Height: height})
	rd.NoResize()
	return &rd, win
}

/**
*	 NewWindowMgrNoResizeNoMinMax
**/
func NewWindowMgrNoResizeNoMinMax(title string, width int, height int, icon *walk.Icon) (*WinResMgr, *walk.MainWindow) {
	rd := WinResMgr{}
	win, _ := walk.NewMainWindowWithName(walkWinName)
	rd.window = win
	rd.parentList = list.New()

	win.SetTitle(title)
	win.SetLayout(walk.NewVBoxLayout())
	win.SetWidth(width)
	win.SetHeight(height)

	if icon != nil {
		win.SetIcon(icon)
	}

	win.SetMinMaxSize(walk.Size{Width: width, Height: height}, walk.Size{Width: width, Height: height})
	rd.NoResize()
	rd.DisableMinMaxBox(true, true)

	return &rd, win
}

/**
*	 NewWindowMgrAds
**/
func NewWindowMgrAds(title string, width int, height int) (*WinResMgr, *walk.MainWindow) {
	rd := WinResMgr{}
	win, _ := walk.NewMainWindowWithName(walkWinName)
	rd.window = win
	rd.parentList = list.New()

	win.SetTitle(title)
	layout := walk.NewVBoxLayout()
	margin := walk.Margins{0, 0, 0, 0}
	layout.SetMargins(margin)
	layout.SetSpacing(0)
	win.SetLayout(layout)
	win.SetWidth(width)
	win.SetHeight(height)

	win.SetMinMaxSize(walk.Size{Width: width, Height: height}, walk.Size{Width: width, Height: height})

	rd.NoResize()
	rd.DisableTitleBar()
	rd.AdsPosition()
	return &rd, win
}

/**
*	HideAndClose
**/
func (m *WinResMgr) HideAndClose() {
	m.window.Synchronize(func() {
		m.window.SetVisible(false)
		m.window.Close()
	})
}

/**
*	Center
**/
func (m *WinResMgr) Center() {
	var x, y, width, height int32
	var rtDesk, rtWindow win.RECT
	win.GetWindowRect(win.GetDesktopWindow(), &rtDesk)
	win.GetWindowRect(m.window.Handle(), &rtWindow)

	width = rtWindow.Right - rtWindow.Left
	height = rtWindow.Bottom - rtWindow.Top
	x = (rtDesk.Right - width) / 2
	y = (rtDesk.Bottom - height) / 2

	win.MoveWindow(m.window.Handle(), x, y, width, height, true)
}

/**
*	AdsPosition
**/
func (m *WinResMgr) AdsPosition() {
	var x, y, width, height int32
	var rtDesk, rtWindow win.RECT
	win.GetWindowRect(win.GetDesktopWindow(), &rtDesk)
	win.GetWindowRect(m.window.Handle(), &rtWindow)

	width = rtWindow.Right - rtWindow.Left
	height = rtWindow.Bottom - rtWindow.Top

	x = rtDesk.Right - width
	y = rtDesk.Bottom - (height + 40)

	win.MoveWindow(m.window.Handle(), x, y, width, height, true)
}

/**
*	Foreground
**/
func (m *WinResMgr) Foreground() {
	win.SetForegroundWindow(m.window.Handle())
}

/**
*	NoResize
**/
func (m *WinResMgr) NoResize() {
	defStyle := win.GetWindowLong(m.window.Handle(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_THICKFRAME
	win.SetWindowLong(m.window.Handle(), win.GWL_STYLE, newStyle)
}

/**
*	DisableMinMaxBox
**/
func (m *WinResMgr) DisableMinMaxBox(minBox bool, maxBox bool) {
	defStyle := win.GetWindowLong(m.window.Handle(), win.GWL_STYLE)
	newStyle := defStyle
	if minBox {
		newStyle = newStyle &^ win.WS_MINIMIZEBOX
	}
	if maxBox {
		newStyle = newStyle &^ win.WS_MAXIMIZEBOX
	}
	win.SetWindowLong(m.window.Handle(), win.GWL_STYLE, newStyle)
}

/**
*	DisableCloseBox
**/
func (m *WinResMgr) DisableCloseBox() {
	defStyle := win.GetWindowLong(m.window.Handle(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_SYSMENU
	win.SetWindowLong(m.window.Handle(), win.GWL_STYLE, newStyle)
}

/**
*	DisableTitleBar
**/
func (m *WinResMgr) DisableTitleBar() {
	defStyle := win.GetWindowLong(m.window.Handle(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_CAPTION
	win.SetWindowLong(m.window.Handle(), win.GWL_STYLE, newStyle)
}

/**
*	HSplit
**/
func (m *WinResMgr) HSplit() *walk.Splitter {
	hs, _ := walk.NewHSplitter(m.GetParent())
	m.parentList.PushBack(hs)
	return hs
}

/**
*	VSplit
**/
func (m *WinResMgr) VSplit() *walk.Splitter {
	vs, _ := walk.NewVSplitter(m.GetParent())
	m.parentList.PushBack(vs)
	return vs
}

/**
*	EndSplit
**/
func (m *WinResMgr) EndSplit() {
	if m.parentList.Len() > 0 {
		popData := m.parentList.Remove(m.parentList.Back())
		parent := m.GetParent()
		parent.Children().Add(popData.(walk.Widget))
	}
}

/**
*	GroupBox
**/
func (m *WinResMgr) GroupBox(title string, isVertical bool) *walk.GroupBox {
	gb, _ := walk.NewGroupBox(m.GetParent())
	gb.SetTitle(title)
	if !isVertical {
		gb.SetLayout(walk.NewHBoxLayout())
	} else {
		gb.SetLayout(walk.NewVBoxLayout())
	}
	m.parentList.PushBack(gb)
	return gb
}

/**
*	EndGroupBox
**/
func (m *WinResMgr) EndGroupBox() {
	if m.parentList.Len() > 0 {
		popData := m.parentList.Remove(m.parentList.Back())
		parent := m.GetParent()
		parent.Children().Add(popData.(walk.Widget))
	}
}

/**
*	DefClosing
**/
func (m *WinResMgr) DefClosing() {
	m.window.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		if m.window.Visible() {
			*canceled = true
		}
	})
}

/**
*	StartWindow
**/
func (m *WinResMgr) StartWindow() {
	m.Center()
	m.window.Show()
	m.window.Run()
}

/**
*	StartForeground
**/
func (m *WinResMgr) StartForeground() {
	m.Center()
	m.Foreground()
	m.window.Show()
	m.window.Run()
}

/**
*	StartAds
**/
func (m *WinResMgr) StartAds() {
	//m.AdsPosition()
	m.Foreground()
	m.window.Show()
	m.window.Run()
}

/**
*	Hide
**/
func (m *WinResMgr) Hide() {
	m.window.Hide()
}

/**
*	ShowForeground
**/
func (m *WinResMgr) ShowForeground() {
	m.Center()
	m.Foreground()
	m.window.Show()
}

/**
*	HideStart
**/
func (m *WinResMgr) HideStart() {
	m.window.Hide()
	m.window.Run()
}

/**
*	addObj
**/
func (m *WinResMgr) addObj(item walk.Widget) {
	if m.parentList.Len() == 0 {
		m.window.Children().Add(item)
	} else {
		parent := m.parentList.Back().Value.(walk.Container)
		parent.Children().Add(item)
	}
}

/**
*	GetParent
**/
func (m *WinResMgr) GetParent() walk.Container {
	if m.parentList.Len() > 0 {
		parent := m.parentList.Back().Value.(walk.Container)
		return parent
	} else {
		return m.window
	}
}

/**
*	MultiLineLabel
**/
func (m *WinResMgr) MultiLineLabel(text string) *walk.Label {
	ne, _ := walk.NewLabelWithStyle(m.GetParent(), win.SS_EDITCONTROL) //|win.SS_CENTER)
	ne.SetText(text)
	ne.SetAlignment(walk.AlignHCenterVCenter)
	ne.SetTextAlignment(walk.AlignCenter)

	m.addObj(ne)
	return ne
}

/**
*	Label
**/
func (m *WinResMgr) Label(text string) *walk.Label {
	ne, _ := walk.NewLabel(m.GetParent())
	ne.SetText(text)
	ne.SetTextAlignment(walk.AlignDefault)

	m.addObj(ne)
	return ne
}

/**
* LabelCenter
**/
func (m *WinResMgr) LabelCenter(text string) *walk.Label {
	ne, _ := walk.NewLabel(m.GetParent())
	ne.SetText(text)
	ne.SetTextAlignment(walk.AlignCenter)

	m.addObj(ne)
	return ne
}

/**
* LabelRight
**/
func (m *WinResMgr) LabelRight(text string) *walk.Label {
	ne, _ := walk.NewLabel(m.GetParent())
	ne.SetText(text)
	ne.SetTextAlignment(walk.AlignFar)

	m.addObj(ne)
	return ne
}

/**
* LabelLeft
**/
func (m *WinResMgr) LabelLeft(text string) *walk.Label {
	ne, _ := walk.NewLabel(m.GetParent())
	ne.SetText(text)
	ne.SetTextAlignment(walk.AlignNear)

	m.addObj(ne)
	return ne
}

/**
*	CheckBox
**/
func (m *WinResMgr) CheckBox(text string, checked bool, attachFunc func()) *walk.CheckBox {
	cb, _ := walk.NewCheckBox(m.GetParent())
	cb.SetText(text)
	cb.SetChecked(checked)
	cb.CheckStateChanged().Attach(attachFunc)

	m.addObj(cb)
	return cb
}

/**
*	DropDownBox
**/
func (m *WinResMgr) DropDownBox(data []string) *walk.ComboBox {
	cb, _ := walk.NewDropDownBox(m.GetParent())
	cb.SetModel(data)
	cb.SetCurrentIndex(0)

	m.addObj(cb)
	return cb
}

/**
*	PushButton
**/
func (m *WinResMgr) PushButton(text string, clickFunc func()) *walk.PushButton {
	btn, _ := walk.NewPushButton(m.GetParent())
	btn.SetText(text)
	btn.Clicked().Attach(clickFunc)

	m.addObj(btn)
	return btn
}

/**
*	NumberEdit
**/
func (m *WinResMgr) NumberEdit() *walk.NumberEdit {
	ne, _ := walk.NewNumberEdit(m.GetParent())
	m.addObj(ne)
	return ne
}

/**
*	NumberEdit
**/
func (m *WinResMgr) NumberEditInt(val int) *walk.NumberEdit {
	ne, _ := walk.NewNumberEdit(m.GetParent())
	ne.SetValue(float64(val))
	m.addObj(ne)
	return ne
}

/**
*	LineEdit
**/
func (m *WinResMgr) LineEdit(ro bool) *walk.LineEdit {
	ne, _ := walk.NewLineEdit(m.GetParent())
	ne.SetReadOnly(ro)

	m.addObj(ne)
	return ne
}

/**
*	TextEdit
**/
func (m *WinResMgr) TextEdit(ro bool) *walk.TextEdit {
	ne, _ := walk.NewTextEdit(m.GetParent())
	ne.SetReadOnly(ro)

	m.addObj(ne)
	return ne
}

/**
*	TextEdit
**/
func (m *WinResMgr) TextEdit2(msg string, ro bool) *walk.TextEdit {
	ne, _ := walk.NewTextEdit(m.GetParent())
	ne.SetText(msg)
	ne.SetReadOnly(ro)

	m.addObj(ne)
	return ne
}

/**
*	TextArea
**/
func (m *WinResMgr) TextArea(ro bool) *walk.TextEdit {
	ne, _ := walk.NewTextEditWithStyle(m.GetParent(), win.WS_VSCROLL)
	ne.SetReadOnly(ro)

	m.addObj(ne)
	return ne
}

func (m *WinResMgr) WebView(url string) *walk.WebView {
	wv, _ := walk.NewWebView(m.GetParent())
	wv.SetURL(url)

	m.addObj(wv)
	return wv
}

/**
*	Slider
**/
func (m *WinResMgr) Slider(minVal int, maxVal int, defVal int) *walk.Slider {
	ne, _ := walk.NewSlider(m.GetParent())
	ne.SetRange(minVal, maxVal)
	ne.SetValue(defVal)

	m.addObj(ne)
	return ne
}

/**
*	TableView
**/
func (m *WinResMgr) TableView(model interface{}, header []tableViewHeader, checkBox bool, multiSelect bool) *walk.TableView {

	tv, _ := walk.NewTableView(m.GetParent())
	tv.SetCheckBoxes(checkBox)
	tv.SetMultiSelection(multiSelect)
	tv.SetModel(model)

	for i := 0; i < len(header); i++ {
		col := walk.NewTableViewColumn()
		col.SetTitle(header[i].Title)
		col.SetWidth(header[i].Width)

		switch header[i].Align {
		case "center":
			col.SetAlignment(walk.AlignCenter)
		case "right":
			col.SetAlignment(walk.AlignFar)
		case "left":
			col.SetAlignment(walk.AlignNear)
		}

		tv.Columns().Add(col)
	}

	m.addObj(tv)
	return tv
}

/**
*	ImageViewFromFile
**/
func (m *WinResMgr) ImageViewFromFile(imgFile string) *walk.ImageView {
	iv, _ := walk.NewImageView(m.GetParent())
	img, imgErr := walk.NewImageFromFile(imgFile)

	if imgErr == nil {
		iv.SetImage(img)
	}
	m.addObj(iv)
	return iv
}

/**
*	ImageViewFromFile2
**/
func (m *WinResMgr) ImageViewFromFile2(imgFile string, im walk.ImageViewMode) *walk.ImageView {
	iv, _ := walk.NewImageView(m.GetParent())
	img, imgErr := walk.NewImageFromFile(imgFile)

	if imgErr == nil {
		iv.SetMode(im)
		iv.SetImage(img)
	}
	m.addObj(iv)
	return iv
}

/**
*	ImageView
**/
func (m *WinResMgr) ImageView() *walk.ImageView {
	iv, _ := walk.NewImageView(m.GetParent())
	m.addObj(iv)
	return iv
}

/**
*	LoadImage
**/
func LoadImage(fileName string) *walk.Image {
	retImage, retImageErr := walk.NewImageFromFile(fileName)

	if retImageErr != nil {
		return nil
	}
	return &retImage
}

/**
*	SetMultiLineText
**/
func MultiLineText(text string, limit int) string {
	runeText := []rune(text)
	var res string

	for {
		if len(runeText) <= limit {
			res += string(runeText)
			break
		}

		res += string(runeText[:limit]) + "\n"
		runeText = runeText[limit:]
	}
	return res
}
