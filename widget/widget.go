package widget

import (
	"strconv"
	
	"gotui/base"
)

const (
	BLACK			=1
	RED				=2
	GREEN			=3
	YELLOW			=4
	BLUE			=5
	MAGENTA			=6
	CYAN			=7
	WHITE			=8
	DEFAULT			=10
)

const (
	NORMAL 			=0
	BOLD			=1
	SOFT			=2
	ITALIC			=3
	UNDERLINE		=4
	INVERTED		=7
	INVISIBLE		=8
	SOMETHING		=9
)

type Widget struct {
	canvas base.Canvas
	style			int
	foreground		int
	background		int
	content string
}

func (this *Widget) SetCanvas(canvas base.Canvas) {
	this.canvas = canvas
}
func (this *Widget) SetStyle(style int) {
	this.style = style
}
func (this *Widget) SetForeground(color int) {
	this.foreground = color
}
func (this *Widget) SetBackground(color int) {
	this.background = color
}
func (this *Widget) Draw() string {
	var out string
	
	//DEFAULT COLOR
	if this.foreground == 0 { this.foreground = DEFAULT }
	if this.background == 0 { this.background = DEFAULT }
	
	//DEFINE POSITION
	out = "\033["+strconv.Itoa(int(this.canvas.X))+";"+strconv.Itoa(int(this.canvas.Y))+"H"
	
	//DEFINE STYLE
	out += "\033["
	out += strconv.Itoa(this.style)+";"
	out += "3"+strconv.Itoa(this.foreground-1)+";"
	out += "4"+strconv.Itoa(this.background-1)+"m"
	
	//SET CONTENT
	out += this.content
	
	//UNDEFINE STYLE
	out += "\033[0m"
	
	return out
}