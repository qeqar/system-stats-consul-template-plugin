package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"os"
)

func memory() {
	memory, _ := mem.VirtualMemory()
	fmt.Println(memory)
}

func cpuInfo() {
	cpuStats, _ := cpu.Info()
	fmt.Println(cpuStats)
}

func cpuCount() {
	count, _ := cpu.Counts(true)
	fmt.Println(count)
}

func main() {
	for _, statsType := range os.Args {
		if statsType == "memory" {
			memory()
		} else if statsType == "cpuInfo" {
			cpuInfo()
		} else if statsType == "cpuCount" {
			cpuCount()
		}
	}
	os.Exit(0)
}
