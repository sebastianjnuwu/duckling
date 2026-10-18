package menu

import (
	"duckling/engine/audio"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// menu_room_update.go — roteia o modal da sala por aba:
// 0 browse (LAN), 1 criar, 2 entrar.

func (l *RoomState) click() { _ = audio.Play("click") }

func (l *RoomState) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		l.nameFocus, l.jFocus = false, 0
		if l.tab == 0 {
			l.onBack()
		} else {
			l.tab = 0
		}
		return nil
	}
	switch l.tab {
	case 1:
		l.scanStop()
		return l.configUpdate()
	case 2:
		return l.joinUpdate()
	default:
		l.scanStart()
		return l.roomsUpdate()
	}
}
