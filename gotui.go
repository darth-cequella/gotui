package main

import (
	"fmt"
	"log"
	"os"
    "os/exec"
    "os/signal"
    "syscall"
	
	"gotui/widget"
	"gotui/base"
)

func cleanup() {
	fmt.Println("Ola mundo!")
	//exec.Command("stty", "sane").Run()
}

func main() {
	screen,err := base.GetScreenCanvas()
	if err != nil { log.Fatal(err) }
	
	var label widget.Label
	
	test := base.Canvas{0,0, screen.W/2,screen.H/2}
	
	label.SetContent("Testando")
	label.SetCanvas( test )
	label.SetForeground(widget.WHITE)
	label.SetBackground(widget.BLUE)
	label.SetStyle(widget.BOLD)
	
	fmt.Println( label.Draw() )
	
	//------------------------------------------------------------------
	channel := make(chan os.Signal, 1)
    signal.Notify(channel, os.Interrupt)
    signal.Notify(channel, syscall.SIGTERM)
    go func() {
        <-channel
        cleanup()
        os.Exit(1)
    }()
    
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
        fmt.Println(string(b))
	}
}
