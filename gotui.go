package main

import (
	"os"
	"syscall"
	"unsafe"
	"fmt"
	
	"gotui/widget"
)

const (
    _TIOCGWINSZ = 0x5413 // On OSX use 1074295912. Thanks zeebo
)

type WINSIZE struct {
    Row    uint16
    Col    uint16
    Xpixel uint16
    Ypixel uint16
}

//What the fuck?
func GetWinsize() (*WINSIZE, error) {
    ws := new(WINSIZE)

    r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(_TIOCGWINSZ),
        uintptr(unsafe.Pointer(ws)),
    )

    if int(r1) == -1 {
        return nil, os.NewSyscallError("GetWinsize", errno)
    }
    return ws, nil
}

func main() {
	win,_ := GetWinsize()
	
	title := "Ola mundo"
	space := (int(win.Col) - len(title))/2
	var test string
	for i:=0; i<space; i++ { test+=" " }
	test += title
	for i:=0; i<space; i++ { test+=" " }
	
	fmt.Println( widget.Label(test, widget.NORMAL, widget.WHITE, widget.CYAN) )
}
