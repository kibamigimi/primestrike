package main

import (
	"image/color"
	"math/rand"
)

var randprime = []int{2, 3, 5, 7}

type enemy struct {
	x      float32
	y      float32
	r      float32
	color  color.NRGBA
	number int
	//vx float32
	//vy float32
}

func (e *enemy) start() enemy {
	e.x = e.setX()
	e.y = 50
	e.r = 20
	e.color = color.NRGBA{0x00, 0x00, 0x00, 0xff}
	e.number = e.setnumber()
	return *e
}
func (e *enemy) setX() float32 {
	a := float32(rand.Intn(690))
	if a < 50 {
		a = e.setX()
	}
	return a
}

func (e *enemy) setnumber() int {
	rand.Shuffle(len(randprime), func(i, j int) {
		randprime[i], randprime[j] = randprime[j], randprime[i]
	})

	randprime = randprime[0:3]
	a := 1
	for i := 0; i < level+2; i++ {
		b := randprime[rand.Intn(len(randprime))]
		a = a * b
	}
	return a
}
