package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(formatDate(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func formatDate(line string) string {
	fields := strings.Fields(line)
	d, _ := strconv.ParseInt(fields[0], 10, 64)
	t := time.Unix(d, 0)
	fields[0] = t.UTC().Format("[15:04:05]")
	return strings.Join(fields, " ")
}
