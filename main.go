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
func calculateMemoryTotal(pid int, procMap map[int]Process, children map[int][]Process, totals map[int]int) int {
	proc, ok := procMap[pid]
	if !ok {
		return 0
	}

	total := proc.RSS
	for _, child := range children[pid] {
		total += calculateMemoryTotal(child.PID, procMap, children, totals)
	}

	totals[pid] = total
	return total
}

func main() {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read /proc: %v\n", err)
		os.Exit(1)
	}

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
				fields := strings.Fields(line)
				if len(fields) >= 2 {
					val, err := strconv.Atoi(fields[1])
					if err == nil {
						ppid = val
					}
				}
			}
			if strings.HasPrefix(line, "VmRSS:") {
				fields := strings.Fields(line)
				if len(fields) >= 2 {
					val, err := strconv.Atoi(fields[1])
					if err == nil {
						rss = val
					}
				}
			}
			if strings.HasPrefix(line, "Name:") {
				fields := strings.Fields(line)
				if len(fields) >= 2 {
					name = fields[1]
				}
			}
		}

		processes = append(processes, Process{PID: pid, PPid: ppid, RSS: rss, Name: name})
	}

	children := make(map[int][]Process)
	for _, proc := range processes {
		children[proc.PPid] = append(children[proc.PPid], proc)
	}

	totals := make(map[int]int)

	// create a map of processes for quick lookup
	procMap := make(map[int]Process)
	for _, proc := range processes {
		procMap[proc.PID] = proc
	}

	calculateMemoryTotal(1, procMap, children, totals)

	fmt.Printf("%-10s %-8s %-10s %s\n", "PID", "RSS(KB)", "TOTAL(KB)", "CMD")

	// stack-based tree traversal to print in correct order
	stack := []StackNode{{pid: 1, prefix: "", indent: ""}}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, proc := range processes {
			if proc.PID == node.pid {
				fmt.Printf("%s%-10d %-8d %-10d %s\n", node.prefix, proc.PID, proc.RSS, totals[proc.PID], proc.Name)
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
