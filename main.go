package main

import (
	"image/color"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var (
	nX        float32
	nY        float32
	vie       int
	point     int
	game      bool
	caisse    int
	withBrick int
	w         fyne.Window
	paddle    *canvas.Rectangle
	contLabel *fyne.Container
	circle    *canvas.Circle
	tbl       []*canvas.Rectangle
	life      *canvas.Text
	label     *canvas.Text
	contBrick *fyne.Container
)

func main() {
	a := app.New()
	w = a.NewWindow("Beakout")
	w.Resize(fyne.NewSize(400, 400))

	game = true
	vie = 3
	withBrick = 80

	paddle = newPaddle(100, 20, 200, 350)
	circle = newBall(15, 15, 200, 330)
	bg := newBG(400, 400)

	life = newLife()
	label = canvas.NewText("PRESS SPACE TO START", color.Black)
	label.TextSize = 30
	contLabel = container.NewMax(label)
	contLabel.Move(fyne.NewPos(15, 185))
	tapp := newTappable()

	contBrick = container.NewWithoutLayout()

	reset()

	cont := container.NewWithoutLayout(bg, circle, contBrick, life, contLabel, paddle)
	contTapp := container.NewMax(tapp, cont)
	w.SetContent(contTapp)
	w.ShowAndRun()

}

func reset() {

	caisse = 6
	tbl2 := generateBrick()
	tbl = tbl2
	for _, b := range tbl {
		b.Move(fyne.NewPos(b.Position().X, b.Position().Y+200))
		contBrick.Add(b)
	}

	label.Text = "PRESS SPACE TO START"
	label.TextSize = 30

	life.Refresh()
	circle.Move(fyne.NewPos(200, 330))
	log.Println("reset game", caisse, vie)
	contLabel.Show()
	label.Refresh()

	game = true
	startGame()
}
func playGame() {

	log.Println("play game")
	go func() {
		for game {

			checkColision(circle, tbl)
			checkColisionWalls(circle, life)
			checkColisonPaddle(circle, paddle)
			checkWin()
			circle.Move(fyne.NewPos(circle.Position1.X+nX, circle.Position1.Y+nY))
			circle.Refresh()
			time.Sleep(time.Millisecond * 40)
		}
		if !game {
			reset()
		}
	}()
}

func startGame() {
	log.Println("start game")
	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if ke.Name == fyne.KeyLeft && paddle.Position().X > 0 {
			paddle.Move(fyne.NewPos(paddle.Position().X-20, paddle.Position().Y))
			paddle.Refresh()
		}
		if ke.Name == fyne.KeyRight && paddle.Position().X < 300 {
			paddle.Move(fyne.NewPos(paddle.Position().X+20, paddle.Position().Y))
			paddle.Refresh()
		}
		if ke.Name == fyne.KeySpace {
			nX = 0
			nY = 5
			contLabel.Hide()
			playGame()
		}
	})
}

func checkColisonPaddle(circle *canvas.Circle, paddle *canvas.Rectangle) {
	if circle.Position().Y+circle.Size().Height == paddle.Position().Y && circle.Position1.X+circle.Size().Width/2 > paddle.Position().X && circle.Position1.X+circle.Size().Width/2 < paddle.Position().X+paddle.Size().Width {
		nY = nY - nY*2
		nX = calculAngle(circle.Position().X, paddle.Position().X)
	}
}

func checkWin() {
	if caisse <= 0 {
		label.Text = " YOU WIN "
		label.TextSize = 60
		contLabel.Show()
		label.Refresh()
		game = false
		time.Sleep(time.Second * 2)
		point++
		life.Text = "Vie: " + strconv.Itoa(vie) + " Point: " + strconv.Itoa(point)
		withBrick -= 10

	}
	if vie <= 0 {
		point = 0
		label.Text = " GAME OVER "
		label.TextSize = 50
		contLabel.Show()
		label.Refresh()
		time.Sleep(time.Second * 2)
		game = false
	}
}
func checkColisionWalls(circle *canvas.Circle, life *canvas.Text) {
	if circle.Position().X < 0 {
		nX = nX - nX*2
	}
	if circle.Position().X > 400-circle.Size().Width {
		nX = nX - nX*2
	}
	if circle.Position().Y < 0 {
		nY = nY - nY*2
	}
	if circle.Position().Y > 400 {
		vie--
		life.Text = "Vie: " + strconv.Itoa(vie) + " Point: " + strconv.Itoa(point)
		if vie < 2 {
			life.Color = color.NRGBA{204, 0, 0, 255}
		}
		circle.Move(fyne.NewPos(200, 330))
		nX = 0
		nY = 0
	}
}

func checkColision(c *canvas.Circle, tbl []*canvas.Rectangle) {
	for _, v := range tbl {
		colision(c, v)
	}
}

func calculAngle(c, p float32) float32 {
	var result float32
	centre := p + 50
	result = (centre - c) / 10
	return result - (result * 2)
}

func colision(c *canvas.Circle, b *canvas.Rectangle) {

	// colision dessus
	if c.Position1.X+c.Size().Width/2 > b.Position().X &&
		c.Position1.X+c.Size().Width/2 < b.Position().X+b.Size().Width &&
		c.Position1.Y+c.Size().Height == b.Position().Y {
		nY = nY - nY*2
		caisse--
		b.Hide()
		b.Move(fyne.NewPos(b.Position().X, b.Position().Y-200))
		log.Println(true)

	}
	// colision dessous
	if c.Position().X+c.Size().Width/2 > b.Position().X &&
		c.Position().X+c.Size().Width/2 < b.Position().X+b.Size().Width &&
		c.Position().Y == b.Position().Y+b.Size().Height {
		nY = nY - nY*2
		caisse--
		b.Hide()
		b.Move(fyne.NewPos(b.Position().X, b.Position().Y-200))
		log.Println(true)

	}
	// colision coté droit
	if c.Position().Y+c.Size().Height/2 > b.Position().Y &&
		c.Position().Y+c.Size().Height/2 < b.Position().Y+b.Size().Height &&
		c.Position().X == b.Position().X+b.Size().Width {
		nX = nX - nX*2
		caisse--
		b.Hide()
		b.Move(fyne.NewPos(b.Position().X, b.Position().Y-200))
		log.Println(true)
	}
	//colision coté gauche
	if c.Position().Y+c.Size().Height/2 > b.Position().Y &&
		c.Position().Y+c.Size().Height/2 < b.Position().Y+b.Size().Height &&
		c.Position().X+c.Size().Width == b.Position().X {
		nX = nX - nX*2
		caisse--
		b.Hide()
		b.Move(fyne.NewPos(b.Position().X, b.Position().Y-200))
		log.Println(true)
	}

}

func generateBrick() (tbl []*canvas.Rectangle) {

	col := 250
	row := 50
	i := caisse
	for i > 0 {
		if i == 3 {
			row = 80
			col = 250
		}
		bloc := newBloc(float32(withBrick), 20, float32(col), float32(row)-200)
		tbl = append(tbl, bloc)
		col -= 100
		i--
	}
	log.Println("TBL généré ")

	return
}
