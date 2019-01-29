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

func cpuStats() {
	cpuStats := cpu.InfoStat{}
	fmt.Println(cpuStats)
}

func main() {
	for _, statsType := range os.Args {
		if statsType == "memory" {
			memory()
		} else if statsType == "cpu" {
			cpuStats()
		}
	}
	os.Exit(0)
}
