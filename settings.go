package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func newBall(w, h, pX, pY float32) *canvas.Circle {

	b := canvas.NewCircle(color.Black)
	b.Resize(fyne.NewSize(w, h))
	b.Move(fyne.NewPos(pX, pY))

	return b
}

func newBloc(w, h, pX, pY float32) *canvas.Rectangle {

	b := canvas.NewRectangle(color.NRGBA{204, 0, 0, 150})
	b.Resize(fyne.NewSize(w, h))
	b.Move(fyne.NewPos(pX, pY))
	b.StrokeWidth = 4
	b.StrokeColor = color.NRGBA{204, 0, 0, 255}

	return b
}
func newBG(w, h float32) *canvas.Rectangle {

	b := canvas.NewRectangle(color.White)
	b.Resize(fyne.NewSize(w, h))
	b.StrokeWidth = 2
	b.StrokeColor = color.NRGBA{50, 0, 0, 255}

	return b
}
func newPaddle(w, h, pX, pY float32) *canvas.Rectangle {

	b := canvas.NewRectangle(color.Black)
	b.Resize(fyne.NewSize(w, h))
	b.Move(fyne.NewPos(pX-w/2, pY))

	return b
}

func newLife() *canvas.Text {
	life := canvas.NewText("Vie: "+strconv.Itoa(vie)+" Point: "+strconv.Itoa(point), color.Black)
	life.TextSize = 15
	life.Move(fyne.NewPos(10, 20))

	return life
}
