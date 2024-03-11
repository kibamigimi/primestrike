package main

import (
	"image/color"
	"math"
)

const (
	circle_r float32 = 20
)

type ball struct {
	x      float32
	y      float32
	vx     float32
	vy     float32
	speed  float32
	number int
	num    int
	color  color.NRGBA
	shoot  bool
}

func (b *ball) first() ball {
	b.x = 370.0
	b.y = 450.0
	b.vx = 0
	b.vy = -1
	b.speed = 0.0
	b.number = 0
	b.num = 0
	b.color = color.NRGBA{0x00, 0x40, 0x80, 0xff}
	b.shoot = false
	return *b
}

func (b *ball) move() {
	b.x = b.x + (b.vx * b.speed)
	b.y = b.y + (b.vy * b.speed)
}

func (b *ball) set(line line) {
	l_x := float64(line.releas_x - line.press_x)
	l_y := float64(line.releas_y - line.press_y)
	r := math.Pow(math.Pow(l_x, 2)+math.Pow(l_y, 2), 0.5)
	n_x := -1 * l_y / r
	n_y := l_x / r
	r_x := float64(b.vx) + 2*(-float64(b.vx)*n_x-float64(b.vy)*n_y)*n_x
	r_y := float64(b.vy) + 2*(-float64(b.vx)*n_x-float64(b.vy)*n_y)*n_y
	b.vx = float32(r_x)
	b.vy = float32(r_y)
}
