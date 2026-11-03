package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	path := "."
	showAll, longFormat := false, false

	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "-") {
			path = arg
			continue
		}
		for _, char := range arg[1:] {
			switch char {
			case 'a':
				showAll = true
			case 'l':
				longFormat = true
			default:
				fmt.Fprintf(os.Stderr, "usage: %s [-al] [path]\n", os.Args[0])
				os.Exit(1)
			}
		}
	}

	var names []string

	if showAll {
		names = append(names, ".", "..")
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") && !showAll {
			continue
		}
		names = append(names, entry.Name())
	}
	for _, name := range names {
		if longFormat {
			printLong(path, name)
		} else {
			fmt.Println(name)
		}
	}
}

func printLong(dir, name string) {
	fileInfo, err := os.Lstat(filepath.Join(dir))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
	}
	fileStats := fileInfo.Sys().(*syscall.Stat_t)
	fmt.Println(
		fileStats.Mode,
		fileStats.Nlink,
		fileStats.Uid,
		fileStats.Gid,
		fileStats.Size,
		name,
	)
}
