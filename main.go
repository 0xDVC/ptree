package main

// idea is to create a process memory attribution tool,
// where we can see the memory attribution of a process

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd:= exec.Command("ps", "-ax", "-o", "pid")
	output, err:= cmd.Output()
	if err != nil {
		fmt.Println("error running ps command:", err)
		return
	}

	lines := strings.Split(string(output), "\n")
	for _, line:= range lines {
		line = strings.TrimSpace(line)
		if line!= "" && len(line) > 0 && line[0] >= '0' && line[0] <= '9' {
			fmt.Println(line)
		}
	}
}
