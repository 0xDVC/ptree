package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Process struct {
        PID int
        PPid int
        RSS int
        Name string
}

func main() {
        entries, _ := os.ReadDir("/proc")
        processes:= []Process{} //read processes into a slice

        for _, entry := range entries {
                pidStr:= entry.Name()

                pid, err := strconv.Atoi(pidStr)
                if err != nil {
                        continue
                }

                data, err := os.ReadFile("/proc/" + pidStr + "/status")
                if err != nil {
                        continue
                }


                lines:= strings.Split(string(data), "\n")
                var ppid, rss int
                var name string

                for _, line:= range lines {
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

                processes = append(processes, Process{PID: pid,PPid: ppid,RSS: rss,Name: name})
        }

        children := make(map[int][]Process)
        for _,p := range processes {
                children[p.PPid] = append(children[p.PPid], p)
        }

        //basic tree traversal
        //in unix systems, it's a hierchical traversal for the process tree
        //TODO: implement recursion next to internally walk through
        type Item struct {
                PID int
                Indent string
        }

        stack:= []Item{{PID:1, Indent:""}}
        for len(stack) >0 {
                current := stack[len(stack)-1]
                stack = stack[:len(stack)-1]


                for _,p := range processes{
                        if p.PID == current.PID {
                                fmt.Printf("%s%d %d KB %s\n", current.Indent, p.PID, p.RSS, p.Name)
                                break
                        }
                }

                kids:= children[current.PID]
                for _,child:= range kids {
                        stack = append(stack, Item{PID:child.PID, Indent:current.Indent + " "})
                }
        }
}
