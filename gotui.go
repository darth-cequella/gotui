package gotui

import (
	"fmt"
	"log"
	"os"
    "os/exec"
    "os/signal"
    "syscall"
)

func Stop() {
	err := exec.Command("stty", "-F", "/dev/tty", "sane").Run()
	
	if err != nil{
		log.Fatal(err)
	}
}

func Start() {
	channel := make(chan os.Signal, 1)
    signal.Notify(channel, os.Interrupt)
    signal.Notify(channel, syscall.SIGTERM)
    go func() {
        <-channel
        Stop()
        os.Exit(1)
    }()
    
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run() 
	
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
        fmt.Print(string(b))
	} 

}
