package base

import (
	"os"
	"syscall"
	"unsafe"
)

const (
    _TIOCGWINSZ = 0x5413 // On OSX use 1074295912. Thanks zeebo
)

type Canvas struct {
    W		uint16
    H		uint16
    X		uint16
    Y		uint16
}

//----------------------------------------------------------------------

func GetScreenCanvas() (Canvas, error) {
	ws := new(Canvas)

    r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(_TIOCGWINSZ),
        uintptr(unsafe.Pointer(ws)),
    )

    if int(r1) == -1 {
		var emptyCanvas Canvas
        return emptyCanvas, os.NewSyscallError("GetWinsize", errno)
    }
    return *ws, nil
}
