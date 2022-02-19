package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type tappable struct {
	widget.Icon
}

func newTappable() *tappable {

	icon := &tappable{}
	icon.ExtendBaseWidget(icon)
	return icon
}

func (*tappable) Dragged(ev *fyne.DragEvent) {
	if ev.Position.X >= 0+paddle.Size().Width/2 && ev.Position.X <= 400-paddle.Size().Width/2 {
		paddle.Move(fyne.NewPos(ev.Position.X-paddle.Size().Width/2, paddle.Position().Y))
		paddle.Refresh()
	}
}
func (*tappable) DragEnd() {
	log.Println("STOP")
}
