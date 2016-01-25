package widget

import (
	"strconv"
)

const (
	BLACK		=0
	RED			=1
	GREEN		=2
	YELLOW		=3
	BLUE		=4
	MAGENTA		=5
	CYAN		=6
	WHITE		=7
	DEFAULT		=9
	
	NORMAL 		="\033[0;"
	BOLD		="\033[1;"
	SOFT		="\033[2;"
	ITALIC		="\033[3;"
	UNDERLINE	="\033[4;"
	INVERTED	="\033[7;"
	INVISIBLE	="\033[8;"
	SOMETHING	="\033[9;"
)

func Label(text, style string, foreground, background int) (out string) {
	//BASIC STYLE
	styleType := style
	
	//FOREGROUND
	styleType += "3"+strconv.Itoa(foreground)+";"
	
	//BACKGROUND
	styleType += "4"+strconv.Itoa(background)+"m"
	
	out = styleType+text+"\033[0m"
	return
}

func FramedLabel(text, style string, foreground, background int) (out string) {
	//BASIC STYLE
	styleType := style
	
	//FOREGROUND
	styleType += "3"+strconv.Itoa(foreground)+";"
	
	//BACKGROUND
	styleType += "4"+strconv.Itoa(background)+"m"
	
	out = styleType
	out += "\u256D" //CORNER LEFT-TOP
	for i:=0; i<len(text);i++ {
		out += "\u2500" //TOP
	}
	out += "\u256E\n" //CORNER RIGHT-TOP
	out += "\u2502"+text+"\u2502\n" //LEFT TEXT RIGHT
	out += "\u2570" //CORNER LEFT-BOTTOM
	for i:=0; i<len(text);i++ {
		out += "\u2500" //BOTTOM
	}
	out += "\u256F" //CORNER RIGHT-BOTTOM
	out += "\033[0m\n"
	
	return
}
