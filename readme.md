# ptree

A simple process tree tool that shows memory attribution for processes and their children.


## how to run

```bash
go run main.go
```
or you could:

```bash
go build
./ptree
```

## sample output

```
PID    RSS(KB)    TOTAL(KB)  CMD
1      128        2048       init
├── 276    140        140        syslogd
├── 304    148        148        crond
└── 59     9936       2048       orbstack-agent
    └── 479    1136       1136        sh
        └── 3733   19788      19788       go
```

## requirements

- linux (uses `/proc` filesystem)
- go 1.21+

## Testing

tested on alpine 3.22 (arm64) via orbstack vm. should work on any linux distro.