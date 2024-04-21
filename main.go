package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	// "github.com/pkg/profile"
	"gokomodo-test/cmd"
)

var appName = "gokomodo"
var appVersion = "v1.0.0"

// prof profiling (cpu|mem) used by ldflags
// var prof string

const banner = `
██████   ██████  ██   ██  ██████  ███    ███  ██████  ██████   ██████
██       ██    ██ ██  ██  ██    ██ ████  ████ ██    ██ ██   ██ ██    ██
██   ███ ██    ██ █████   ██    ██ ██ ████ ██ ██    ██ ██   ██ ██    ██
██    ██ ██    ██ ██  ██  ██    ██ ██  ██  ██ ██    ██ ██   ██ ██    ██
 ██████   ██████  ██   ██  ██████  ██      ██  ██████  ██████   ██████
%v : %v 
`

func main() {
	if tz := os.Getenv("TZ"); tz != "" {
		var err error
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			fmt.Printf("error loading location '%s': %v\n", tz, err)
		}
	}
	// if prof == "cpu" {
	// 	defer profile.Start(profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	// }
	// if prof == "mem" {
	// 	defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	// }

	// setup banner
	fmt.Printf(banner, appName, appVersion)
	runtime.GOMAXPROCS(runtime.NumCPU())

	cmd.Run()
}
