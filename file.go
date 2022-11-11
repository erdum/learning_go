package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("test.txt")
	defer file.Close()
	fileStat, _ := file.Stat()

	lines := make([]byte, fileStat.Size())
	file.Read(lines)
	text := string(lines)
	fmt.Println(text)
}
