package task

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func TestABCBCB() {
	fmt.Println("异步任务执行开始")
	printMemoryInfo()
}

func printMemoryInfo() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("=== 内存占用情况 ===")
	fmt.Printf("Alloc: %v MiB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc: %v MiB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys: %v MiB\n", bToMb(m.Sys))
	fmt.Printf("NumGC: %v\n", m.NumGC)
	fmt.Println()
}

func printCPUInfo() {
	cmd := exec.Command("ps", "aux")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("无法获取 CPU 信息:", err)
		return
	}

	lines := strings.Split(string(output), "\n")
	fmt.Println("=== CPU 使用情况 ===")
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println()
}

//func printDiskInfo() {
//	var stat syscall.Statfs_t
//	err := syscall.Statfs("/", &stat)
//	if err != nil {
//		fmt.Println("无法获取磁盘信息:", err)
//		return
//	}
//
//	availableSpace := stat.Bavail * uint64(stat.Bsize)
//	fmt.Println("=== 磁盘可用空间 ===")
//	fmt.Printf("Available Space: %v bytes\n", availableSpace)
//	fmt.Println()
//}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
