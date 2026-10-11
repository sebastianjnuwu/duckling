package menu

import (
	"fmt"

	"duckling/content"
	"duckling/engine/i18n"
	"duckling/engine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

// lobby_rooms.go — tabelinha: X fecha, linhas com ícone
// (cadeado nas privadas), paginação de 4 em 4.

const perPage = 4

func (l *RoomState) pages() int {
	return max(1, (len(l.rooms)+perPage-1)/perPage)
}

func (l *RoomState) roomsUpdate() error {
	if l.x.Push() || l.backB.Update() {
		l.click()
		l.onBack()
		return nil
	}
	if l.create.Update() {
		l.click()
		l.tab = 1
		return nil
	}
	if l.pages() > 1 {
		if l.prevB.Update() && l.page > 0 {
			l.click()
			l.page--
		}
		if l.nextB.Update() && l.page < l.pages()-1 {
			l.click()
			l.page++
		}
	}
	for k := range l.rowB {
		i := l.page*perPage + k
		if i >= len(l.rooms) {
			break
		}
		if l.rowB[k].Update() {
			l.click()
			l.tab = 1
			return nil
		}
	}
	return nil
}

func (l *RoomState) roomsDraw(img *ebiten.Image) {
	if bg := assets.Images["images/bg"]; bg != nil {
		ui.DrawCover(img, bg)
	}
	ui.Dim(img)
	ui.Panel(img, rpX, rpY, rpW, rpH)
	ui.Title(img, ui.IconLan, i18n.T("lobby.rooms"), l.big, 135, ui.Gold)
	l.x.Draw(img, "X", l.font)
	if l.pages() > 1 {
		left(img, fmt.Sprintf("%d/%d", l.page+1, l.pages()), l.font, 920, 140)
	}
	if len(l.rooms) == 0 {
		ui.DrawCenter(img, i18n.T("lobby.empty"), l.font, 300)
	}
	for k := range l.rowB {
		i := l.page*perPage + k
		if i >= len(l.rooms) {
			break
		}
		r := l.rooms[i]
		l.rowB[k].Icon = ui.IconLan
		if r.priv {
			l.rowB[k].Icon = ui.IconLock
		}
		l.rowB[k].Draw(img, fmt.Sprintf("%s %d   %d/%d", i18n.T("lobby.room"), i+1, r.n, r.max), l.font)
	}
	if l.pages() > 1 {
		l.prevB.Draw(img, "<", l.font)
		l.nextB.Draw(img, ">", l.font)
	}
	l.create.Draw(img, i18n.T("lobby.create"), l.font)
	l.backB.Draw(img, i18n.T("shop.back"), l.font)
}

// addRoom entra na tabela (a rede lista as reais aqui).
func (l *RoomState) addRoom(n, max int, priv bool) {
	l.rooms = append(l.rooms, room{n: n, max: max, priv: priv})
}
