package main

import (
	"fmt"
	"image/color"
	"math"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var press_x int
var press_y int
var score int

type status int

var level = 1

var ballprime = []int{2, 3, 5, 7}
var drawlineflag = false
var ballshootflag = false

const (
	statusInit status = iota
	statusSetup
	statusPlay
	statusFin
)

type game struct {
	status   status
	ball     ball
	enemy    enemy
	timer    int
	count    int
	font     font.Face
	shootnum []*shootnum
	lines    []*line
	lineG    line
	width    int
	height   int
}

func newGame() *game {
	var g game
	g.reset()
	return &g
}

func (g *game) reset() {
	g.status = statusInit
	g.ball = g.ball.first()
	g.enemy = g.enemy.start()
	g.font = fontset()
	g.count = 0
	g.shootnum = shootnuminit()
	g.lines = []*line{}
	g.lineG = line{
		press_x:     float32(0),
		press_y:     float32(0),
		releas_x:    float32(0),
		releas_y:    float32(0),
		strokeWidth: 5,
		color:       color.NRGBA{0x00, 0x00, 0x00, 0xff},
	}
	g.width = 740
	g.height = 600
}

func (g *game) Update() error {
	switch g.status {
	case statusInit: //開始画面
		var keys []ebiten.Key
		keys = inpututil.AppendJustReleasedKeys(keys)
		if len(keys) != 1 {
			return nil
		}

		if keys[0] == ebiten.KeySpace {
			g.status = statusSetup

		}
	case statusSetup: //ゲーム画面
		if g.timer >= 5000 {
			g.status = statusFin
		}
		var keys []ebiten.Key
		keys = inpututil.AppendJustPressedKeys(keys)
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			press_x, press_y = ebiten.CursorPosition()
			if math.Pow(float64(g.ball.x)-float64(press_x), 2)+math.Pow(float64(g.ball.y)-float64(press_y), 2) <= math.Pow(float64(circle_r), 2) {
				ballshootflag = true
			} else if press_y <= 430 {
				drawlineflag = true
			}
		} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
			releas_x, releas_y := ebiten.CursorPosition()
			if drawlineflag {
				drawlineflag = false
				if math.Pow(math.Pow(float64((press_x-releas_x)), 2)+math.Pow(float64((press_y-releas_y)), 2), 0.5) > 50 {
					g.lineG = line{
						press_x:     float32(press_x),
						press_y:     float32(press_y),
						releas_x:    float32(releas_x),
						releas_y:    float32(releas_y),
						strokeWidth: 5,
						color:       color.NRGBA{0x00, 0x00, 0x00, 0xff},
					}
					g.lines = append(g.lines, &line{
						press_x:     float32(press_x),
						press_y:     float32(press_y),
						releas_x:    float32(releas_x),
						releas_y:    float32(releas_y),
						strokeWidth: 5,
						color:       color.NRGBA{0x00, 0x00, 0x00, 0xff},
					})
				}
			} else if ballshootflag {
				ballshootflag = false
				if !g.ball.shoot && g.ball.number != 0 && g.check(g.ball.num) {
					g.count += 1
					g.ball.shoot = true
					g.ball.speed = 8
					g.set(g.ball.num, false)
					g.shootnumtimer(g.ball.num)
					g.status = statusPlay
				}

			}
		} else if drawlineflag {
			now_x, now_y := ebiten.CursorPosition()
			if math.Pow(math.Pow(float64((press_x-now_x)), 2)+math.Pow(float64((press_y-now_y)), 2), 0.5) > 50 {
				g.lineG = line{
					press_x:     float32(press_x),
					press_y:     float32(press_y),
					releas_x:    float32(now_x),
					releas_y:    float32(now_y),
					strokeWidth: 5,
					color:       color.NRGBA{226, 226, 226, 0xff},
				}
			} else {
				g.lineG = line{
					press_x:     float32(press_x),
					press_y:     float32(press_y),
					releas_x:    float32(now_x),
					releas_y:    float32(now_y),
					strokeWidth: 5,
					color:       color.NRGBA{0xff, 0x00, 0x00, 0xff},
				}
			}
		} else if ballshootflag {
			now_x, now_y := ebiten.CursorPosition()
			r := math.Pow(math.Pow(float64(press_x-now_x), 2)+math.Pow(float64(press_y-now_y), 2), 0.5)
			g.ball.vx = float32(press_x-now_x) / float32(r)
			g.ball.vy = float32(press_y-now_y) / float32(r)
		}

		if len(keys) != 1 {
			return nil
		}
		switch keys[0] {
		case ebiten.KeyDigit2:
			if g.check(2) {
				g.ball.number = 2
				g.ball.num = 2
			}
		case ebiten.KeyDigit3:
			if g.check(3) {
				g.ball.number = 3
				g.ball.num = 3
			}
		case ebiten.KeyDigit5:
			if g.check(5) {
				g.ball.number = 5
				g.ball.num = 5
			}
		case ebiten.KeyDigit7:
			if g.check(7) {
				g.ball.number = 7
				g.ball.num = 7
			}
		case ebiten.KeyEnter:
			g.reset()
			g.status = statusSetup
			score = score / 2
		}

	case statusPlay:
		if g.ball.x < 0 || g.ball.y < 0 || g.ball.y > 500 || g.ball.x > 740 {
			g.status = statusSetup
			g.ball.first()
		}
		if g.ball.shoot {
			g.ball.move()
		}
		for i := range g.lines {
			line := g.lines[i]
			if bound(*line, g.ball) {
				g.ball.number = g.ball.number * g.ball.num
				g.ball.set(*line)
			}
		}
		if g.judge() {
			if g.enemy.number%g.ball.number == 0 {
				g.enemy.number = g.enemy.number / g.ball.number
			}
			g.ball.first()
			g.status = statusSetup
			if g.enemy.number == 1 {
				g.reset()
				score = g.culscore(score)
				level += 1
				g.status = statusSetup
				if g.timer < 200 {
					g.timer = 0
				} else {
					g.timer -= 200
				}

			}
		}

	case statusFin:
		var keys []ebiten.Key
		keys = inpututil.AppendJustReleasedKeys(keys)
		if len(keys) != 1 {
			return nil
		}

		if keys[0] == ebiten.KeySpace {
			g.status = statusInit
		}
	}

	return nil
}
func bound(line line, ball ball) bool {
	if ball.x+circle_r < min(line.press_x, line.releas_x) ||
		ball.x-circle_r > max(line.press_x, line.releas_x) ||
		ball.y+circle_r < min(line.press_y, line.releas_y) ||
		ball.y-circle_r > max(line.press_y, line.releas_y) {
		return false
	}
	a := (line.releas_y - line.press_y) / (line.releas_x - line.press_x)
	c := -a*line.press_x + line.press_y
	r := math.Pow(math.Pow(float64(a), 2)+1, 0.5)
	d := math.Abs(float64(a)*float64(ball.x)-float64(ball.y)+float64(c)) / r
	return circle_r >= float32(d)
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0xff, 0xff, 0xff, 0xff})
	if g.status == statusInit {
		a, err := g.setopen(screen)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}
		var opbg ebiten.DrawImageOptions
		opbg.GeoM.Translate(0, 0)
		screen.DrawImage(a, &opbg)
	}
	if g.status == statusSetup {
		if ballshootflag {
			arrowcolor := color.NRGBA{220, 220, 220, 0xff}
			arrow_x := g.ball.x + g.ball.vx*150 + 20*(g.ball.vx*float32(math.Cos(150*math.Pi/180))-g.ball.vy*float32(math.Sin(150*math.Pi/180)))
			arrow_y := g.ball.y + g.ball.vy*150 + 20*(g.ball.vx*float32(math.Sin(150*math.Pi/180))+g.ball.vy*float32(math.Cos(150*math.Pi/180)))
			arrow_x2 := g.ball.x + g.ball.vx*150 + 20*(g.ball.vx*float32(math.Cos(210*math.Pi/180))-g.ball.vy*float32(math.Sin(210*math.Pi/180)))
			arrow_y2 := g.ball.y + g.ball.vy*150 + 20*(g.ball.vx*float32(math.Sin(210*math.Pi/180))+g.ball.vy*float32(math.Cos(210*math.Pi/180)))

			vector.StrokeLine(screen, g.ball.x, g.ball.y, g.ball.x+g.ball.vx*150, g.ball.y+g.ball.vy*150, 3, arrowcolor, true)
			vector.StrokeLine(screen, g.ball.x+g.ball.vx*150, g.ball.y+g.ball.vy*150, arrow_x, arrow_y, 3, arrowcolor, true)
			vector.StrokeLine(screen, g.ball.x+g.ball.vx*150, g.ball.y+g.ball.vy*150, arrow_x2, arrow_y2, 3, arrowcolor, true)
		}
	}
	if g.status == statusSetup || g.status == statusPlay {
		g.timer += 1
		g.shootnumcount()
		g.ball.color = color.NRGBA{0xff, 0x00, 0x00, 0xff}
		under_color := color.NRGBA{254, 220, 189, 0xff}
		vector.DrawFilledRect(screen, float32(g.width), float32(g.height), float32(-g.width), -100, under_color, true)
		vector.DrawFilledCircle(screen, g.ball.x, g.ball.y, circle_r, g.ball.color, true)
		vector.DrawFilledCircle(screen, g.enemy.x, g.enemy.y, g.enemy.r, g.enemy.color, true)
		vector.StrokeLine(screen, g.lineG.press_x, g.lineG.press_y, g.lineG.releas_x, g.lineG.releas_y, g.lineG.strokeWidth, g.lineG.color, true)
		vector.StrokeRect(screen, 500, 530, 200, 40, 5, color.Black, true)

		text.Draw(screen, fmt.Sprint("制限時間"), g.font, 500, 525, color.Black)
		if g.timer < 5000 {
			vector.DrawFilledRect(screen, 500+1, 530+1, 200-float32(g.timer)/25, 38, color.NRGBA{0xff, 0x00, 0x00, 0xff}, true)
		}
		for i := range g.lines {
			line := g.lines[i]
			vector.StrokeLine(screen, line.press_x, line.press_y, line.releas_x, line.releas_y, line.strokeWidth, line.color, true)
		}
		text.Draw(screen, fmt.Sprint(g.ball.number), g.font, int(g.ball.x)-8, int(g.ball.y)+int(circle_r)+20, color.Black)
		text.Draw(screen, fmt.Sprint(g.enemy.number), g.font, int(g.enemy.x)-8, int(g.enemy.y)-int(g.enemy.r)-10, color.Black)
		for i := 0; i < len(ballprime); i++ {
			vector.StrokeRect(screen, float32(30+i*55-25), float32(570), float32(50), 20, 3, color.Black, true)
			if g.check(ballprime[i]) == false {
				vector.DrawFilledCircle(screen, float32(30+i*55), float32(540), float32(25), color.NRGBA{128, 128, 128, 0xff}, true)
				vector.DrawFilledRect(screen, float32(30+i*55-25), float32(570+1), float32(g.shootnumtime(ballprime[i])/9), 18, color.NRGBA{128, 128, 128, 0xff}, true)

			}
			if g.ball.num == ballprime[i] {
				vector.StrokeCircle(screen, float32(30+i*55), float32(540), float32(25), float32(3), g.ball.color, true)
			} else {
				vector.StrokeCircle(screen, float32(30+i*55), float32(540), float32(25), float32(3), color.Black, true)

			}
			text.Draw(screen, fmt.Sprint(ballprime[i]), g.font, int(30+i*55-8), int(540+8), color.Black)
		}
		text.Draw(screen, fmt.Sprint("score:", score), g.font, 320, 550, color.Black)
	}
	if g.status == statusFin {
		text.Draw(screen, fmt.Sprint("score:", score), g.font, 320, 300, color.Black)
	}

}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func (g *game) judge() bool {
	rsum := float64(circle_r + g.enemy.r)
	d := math.Pow(math.Pow(float64(g.ball.x-g.enemy.x), 2)+math.Pow(float64(g.ball.y-g.enemy.y), 2), 0.5)
	if rsum > d {
		return true
	}
	return false
}

func (g *game) setopen(screen *ebiten.Image) (*ebiten.Image, error) {
	fname := "open.png"
	var err error
	img, _, err := ebitenutil.NewImageFromFile(fname)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (g *game) culscore(score int) int {
	score += level + 2*100 - g.count*30
	return score
}

func fontset() font.Face {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	const dpi = 72
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		return nil
	}
	return mplusNormalFont
}
