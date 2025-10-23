package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    entries, _ := os.ReadDir("/proc")

    for _, entry := range entries {
        pid:= entry.Name()

        // see if this is a PID directory?
        if _, err := strconv.Atoi(entry.Name()); err != nil {
                continue
        }

        // read the status file
        data,err := os.ReadFile("/proc/" +pid+ "/status")
        if err != nil {
                continue
        }

        //parse it
        lines:= strings.Split(string(data), "\n")
        var ppid, rss, name string

        for _, line:= range lines {
                if strings.HasPrefix(line, "PPid:") {
                        ppid= strings.Fields(line)[1]
                }
                if strings.HasPrefix(line, "VmRSS:") {
                        rss = strings.Fields(line)[1]
                }
                if strings.HasPrefix(line, "Name:") {
                        name = strings.Fields(line)[1]
                }
        }

        fmt.Printf("pid:  %s, ppid: %s, rss: %s KB, name: %s\n", pid, ppid, rss, name)

    }
}