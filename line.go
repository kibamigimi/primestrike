package main

import (
	"image/color"
)

type line struct {
	//id       int
	press_x     float32
	press_y     float32
	releas_x    float32
	releas_y    float32
	strokeWidth float32
	color       color.NRGBA
}

func (l *line) start() line {
	//l.id = 0
	l.press_x = 0
	l.press_y = 0
	l.releas_x = 0
	l.releas_y = 0
	l.strokeWidth = 5
	l.color = color.NRGBA{0x00, 0x00, 0x00, 0xff}
	return *l
}
