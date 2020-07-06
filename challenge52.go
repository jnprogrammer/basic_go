package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS:\t ", runtime.GOOS)
	fmt.Println("ARCH:\t ", runtime.GOARCH)
	fmt.Println("CPUs:\t ", runtime.NumCPU())
	fmt.Println("Go routines:\t ", runtime.NumGoroutine())
}

// displayed the OS, ARCH, CPUS and routines. then i used build and installed the application.
