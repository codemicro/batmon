# batmon

*Basic Golang program to monitor battery charge level*

---

`batmon` is a basic utility designed to be in the background of a system to monitor its battery level. Only Linux is supported.

### Installation 

The following requires at least version 1.16 of the Go SDK installed, and to have GOPATH on your PATH.

```
go install github.com/codemicro/batmon@latest
```

### Usage

Start at startup, for example with `cron`:

```
@reboot /home/akp/go/bin/batmon
```

Note that `batmon` is blocking.

### Battery information

Battery capacity is retrieved from `/sys/class/power_supply/BAT0/capacity` by default. This can be changed in the `batteryPath` constant in `main.go`.