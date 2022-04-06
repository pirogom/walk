// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package walk

import (
	"sync"
)

var (
	PosMgr PositionManager
)

type SavedPosition struct {
	PosX       int
	PosY       int
	Width      int
	Height     int
	DeskWidth  int
	DeskHeight int
}

/**
*	PositionManager
**/
type PositionManager struct {
	sync.Mutex

	PosX       int
	PosY       int
	Width      int
	Height     int
	DeskWidth  int
	DeskHeight int

	updated bool

	SavedPos SavedPosition
}

/**
*	Update
**/
func (p *PositionManager) Update(x, y, w, h, dw, dh int) {
	p.Lock()
	p.PosX = x
	p.PosY = y
	p.Width = w
	p.Height = h
	p.DeskWidth = dw
	p.DeskHeight = dh

	if !p.updated {
		p.updated = true
	}
	p.Unlock()
}

/**
* Clear
**/
func (p *PositionManager) Clear() {
	p.Lock()
	p.updated = false
	p.Unlock()
}

/**
*	Get
**/
func (p *PositionManager) Get() (x int, y int, w int, h int, dw int, dh int) {
	return p.PosX, p.PosY, p.Width, p.Height, p.DeskWidth, p.DeskHeight
}

/**
*	HasPosition
**/
func (p *PositionManager) HasPosition() bool {
	return p.updated
}

/**
*	Save
**/
func (p *PositionManager) Save() {
	p.SavedPos.PosX = p.PosX
	p.SavedPos.PosY = p.PosY
	p.SavedPos.Width = p.Width
	p.SavedPos.Height = p.Height
	p.SavedPos.DeskWidth = p.DeskHeight
	p.SavedPos.DeskHeight = p.DeskHeight
}

/**
*	Restore
**/
func (p *PositionManager) Restore() {
	p.PosX = p.SavedPos.PosX
	p.PosY = p.SavedPos.PosY
	p.Width = p.SavedPos.Width
	p.Height = p.SavedPos.Height
	p.DeskHeight = p.SavedPos.DeskWidth
	p.DeskHeight = p.SavedPos.DeskHeight
}
