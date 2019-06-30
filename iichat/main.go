package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	serverAndChannel := os.Args[1]
	home, _ := os.UserHomeDir()
	filename := filepath.Join(home, "irc", serverAndChannel, "in")
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s> ", serverAndChannel)
		s, _ := reader.ReadString('\n')
		f.WriteString(s)
	}
}
