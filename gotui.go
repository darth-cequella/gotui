package main

import (
	"fmt"
	"os/exec"
	"bytes"
	"log"
	
	//"gotui/widget"
)

func main() {
	cmd := exec.Command ("echo", "-e lines\ncols | tput -S")
	
	var out bytes.Buffer
	cmd.Stdout = &out
	
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out.String())
	
	//fmt.Println( widget.FramedLabel("Ola mundo", widget.NORMAL, widget.WHITE, widget.CYAN) )
}
