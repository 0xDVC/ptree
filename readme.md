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
PID    RSS(KB)  TOTAL(KB)  CMD
1      128      33752      init
├── 276    140      140        syslogd
├── 304    152      152        crond
├── 403    116      116        udhcpc
├── 467    124      124        getty
└── 59     10040    33092      orbstack-agent:
    └── 3944   1144     23052      sh
        └── 5322   19416    21908      go
            └── 5430   2492     2492       main
```

## requirements

- linux (uses `/proc` filesystem)
- go 1.21+

## Testing

tested on alpine 3.22 (arm64) via orbstack vm. should work on any linux distro.