package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Process struct {
	PID  int
	PPid int
	RSS  int
	Name string
}

type StackNode struct {
	pid    int
	prefix string
	indent string
}

// calculate total memory including children recursively
func calculateMemoryTotal(pid int, processes []Process, children map[int][]Process, totals map[int]int) int {
	var rss int
	for _, proc := range processes {
		if proc.PID == pid {
			rss = proc.RSS
			break
		}
	}

	total := rss
	for _, child := range children[pid] {
		total += calculateMemoryTotal(child.PID, processes, children, totals)
	}

	totals[pid] = total
	return total
}

func main() {
	entries, _ := os.ReadDir("/proc")
	processes := []Process{}

	for _, entry := range entries {
		pidStr := entry.Name()
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			continue
		}

		data, err := os.ReadFile("/proc/" + pidStr + "/status")
		if err != nil {
			continue
		}

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

	children := make(map[int][]Process)
	for _, proc := range processes {
		children[proc.PPid] = append(children[proc.PPid], proc)
	}

	totals := make(map[int]int)
	calculateMemoryTotal(1, processes, children, totals)

	fmt.Printf("%-6s %-8s %-10s %s\n", "PID", "RSS(KB)", "TOTAL(KB)", "CMD")

	// stack-based tree traversal to print in correct order
	stack := []StackNode{{pid:1, prefix:"", indent:""}}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, proc := range processes {
			if proc.PID == node.pid {
				fmt.Printf("%s%-6d %-8d %-10d %s\n", node.prefix, proc.PID, proc.RSS, totals[proc.PID], proc.Name)
				break
			}
		}

		// add children to stack in reverse order
		kids := children[node.pid]
		for i := len(kids) - 1; i >= 0; i-- {
			child := kids[i]
			isLast := (i == len(kids)-1)

			var childPrefix, childIndent string
			if isLast {
				childPrefix = node.indent + "└── "
				childIndent = node.indent + "    "
			} else {
				childPrefix = node.indent + "├── "
				childIndent = node.indent + "│   "
			}

			stack = append(stack, StackNode{
				pid:    child.PID,
				prefix: childPrefix,
				indent: childIndent,
			})
		}
	}
}
