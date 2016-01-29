package gotui

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	
	"gotui/widget"
)

type Screen struct {
}

func (this *Screen) stop() {
	//Reabilite input/output in terminal
	exec.Command("stty", "-F", "/dev/tty", "sane").Run()
	this.clearScreen()
	//Unhide cursor
	fmt.Print("\033[?25h")
}
func (this *Screen) clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func (this *Screen) hideCursor() {
	//Ignore input/output in terminal
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run() 
	//Hide cursor
	fmt.Print("\033[?25l")
}
func (this *Screen) Start() {
	//Create a channel to listen the SIGTERMs
	channel := make(chan os.Signal, 1)
    signal.Notify(channel, os.Interrupt)
    signal.Notify(channel, syscall.SIGTERM)
    go func() {
        <-channel
        this.stop()
        os.Exit(1)
    }()
    
    this.clearScreen()
    this.hideCursor()
	
    var teste widget.Button
    teste.SetForeground(widget.WHITE)
    teste.SetBackground(widget.CYAN)
    teste.SetStyle(widget.BOLD)
    teste.SetContent("Testando")

	//Open infinite loop
	//var keyboard base.Keyboard
	//var b []byte = make([]byte, 2)
	for {
		//os.Stdin.Read(b)
        fmt.Printf( teste.Draw() )
	} 

}
