// Copyright 2023 kzzfxf
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

type GUIUpdater interface {
	UpdateTunnelsTable(rows [][]string)
	UpdateBridgesTable(rows [][]string)
	Render()
}

type ui interface {
	Init(updater GUIUpdater)
	UpdateTunnelsTable(rows [][]string)
	UpdateBridgesTable(rows [][]string)
	Render()
}

type uImpl struct {
	updater GUIUpdater
}

var UI ui = &uImpl{}

// Init
func (ui *uImpl) Init(updater GUIUpdater) {
	if ui.updater == nil {
		ui.updater = updater
	}
}

// UpdateTunnelsTable
func (ui *uImpl) UpdateTunnelsTable(rows [][]string) {
	if ui.updater != nil {
		ui.updater.UpdateTunnelsTable(rows)
	}
}

// UpdateBridgeTable
func (ui *uImpl) UpdateBridgesTable(rows [][]string) {
	if ui.updater != nil {
		ui.updater.UpdateBridgesTable(rows)
	}
}

// Render
func (ui *uImpl) Render() {
	if ui.updater != nil {
		ui.updater.Render()
	}
}
