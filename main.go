package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// info is all in the /proc/[pid]/status file so we are just parsing it
type Process struct {
	PID  int    
	PPid int    // parent pid
	RSS  int    // resident set size (memory used)
	Name string 
}


// stack-based tree traversal to print the process tree
// used a stack because i want to print children in reverse order
type StackNode struct {
	pid    int    
	prefix string // character to print before the pid
	indent string // character to use for children of this node
}

func main() {
	// read /proc directory 
	entries, _ := os.ReadDir("/proc")
	processes := []Process{}

	// iterate through all the numbered directories in /proc
	for _, entry := range entries {
		pidStr := entry.Name()
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			continue 
		}

		// read the status file for this process
		data, err := os.ReadFile("/proc/" + pidStr + "/status")
		if err != nil {
			continue 
		}

		// parse the status file line by line
		lines := strings.Split(string(data), "\n")
		var ppid, rss int
		var name string

		for _, line := range lines {
			if strings.HasPrefix(line, "PPid:") {
				ppid, _ = strconv.Atoi(strings.Fields(line)[1])
			}
			if strings.HasPrefix(line, "VmRSS:") {
				rss, _ = strconv.Atoi(strings.Fields(line)[1])
			}
			if strings.HasPrefix(line, "Name:") {
				name = strings.Fields(line)[1]
			}
		}

		processes = append(processes, Process{PID: pid, PPid: ppid, RSS: rss, Name: name})
	}

	// build a map of parent->children for tree traversal
	children := make(map[int][]Process)
	for _, proc := range processes {
		children[proc.PPid] = append(children[proc.PPid], proc)
	}

	fmt.Printf("PID    RSS(KB)  CMD\n")


	stack := []StackNode{{pid: 1, prefix: "", indent: ""}} //start at root

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop from end

		for _, proc := range processes {
			if proc.PID == node.pid {
				fmt.Printf("%s%-6d %-8d %s\n", node.prefix, proc.PID, proc.RSS, proc.Name)
				break
			}
		}

		// add children to stack in reverse order
		kids := children[node.pid]
		for i := len(kids) - 1; i >= 0; i-- {
			child := kids[i]
			isLast := (i == len(kids)-1) // last chld gets a different prefix

			var childPrefix, childIndent string
			if isLast {
				childPrefix = node.indent + "└── " // child is last
				childIndent = node.indent + "    " // no vertical line
			} else {
				childPrefix = node.indent + "├── " // child is not last
				childIndent = node.indent + "│   " // vertical line
			}

			stack = append(stack, StackNode{
				pid:    child.PID,
				prefix: childPrefix,
				indent: childIndent,
			})
		}
	}
}
