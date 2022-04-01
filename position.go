// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package walk

import "sync"

var (
	PosMgr PositionManager
)

/**
*	PositionManager
**/
type PositionManager struct {
	sync.Mutex
	PosX, PosY, Width, Height int
	updated                   bool
}

/**
*	Update
**/
func (p *PositionManager) Update(x, y, w, h int) {
	p.Lock()
	p.PosX = x
	p.PosY = y
	p.Width = w
	p.Height = h
	if !p.updated {
		p.updated = true
	}
	p.Unlock()
}

/**
*	Get
**/
func (p *PositionManager) Get() (x int, y int, w int, h int) {
	return p.PosX, p.PosY, p.Width, p.Height
}

/**
*	HasPosition
**/
func (p *PositionManager) HasPosition() bool {
	return p.updated
}
