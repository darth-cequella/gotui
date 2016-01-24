package main

import (
	"fmt"
	"strconv"
)

const (
	BLACK	=0
	RED		=1
	GREEN	=2
	YELLOW	=3
	BLUE	=4
	MAGENTA	=5
	CYAN	=6
	WHITE	=7
	DEFAULT	=9
)

/*const (	
	F_BOLD 		="[1m"
	F_DIM 		="[2m"
	F_UNDER		="[4m"
	F_BLINK		="[5m"
	F_REVERSE	="[7m"
	F_INVIS		="[8m"
	F_DEFAUL	="[0m"
)*/

func Colorize(text string, foreground, background int) (out string) {
	foregroundColor := "\033[3"+strconv.Itoa(foreground)+"m"
	backgroundColor := "\033[4"+strconv.Itoa(background)+"m"
	def := "\033[39m\033[49m"
	out = foregroundColor+backgroundColor+text+def
	return
}

func main() {
	fmt.Println( Colorize("Ola", YELLOW, GREEN) )
	fmt.Println( Colorize("Ola", YELLOW, GREEN) )
	fmt.Println( Colorize("Ola", YELLOW, GREEN) )
}
