package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
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

func modeString(mode uint32) string {
	var s [10]byte
	switch mode & syscall.S_IFMT {
	case syscall.S_IFDIR:
		s[0] = 'd'
	case syscall.S_IFLNK:
		s[0] = 'l'
	case syscall.S_IFBLK:
		s[0] = 'b'
	case syscall.S_IFCHR:
		s[0] = 'c'
	case syscall.S_IFIFO:
		s[0] = 'p'
	case syscall.S_IFSOCK:
		s[0] = 's'
	default:
		s[0] = '-'
	}
	return string(s[:1]) + os.FileMode(mode).Perm().String()[1:]
}

func printLong(dir, name string) {
	fileInfo, err := os.Lstat(filepath.Join(dir, name))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
	}

	fileStats := fileInfo.Sys().(*syscall.Stat_t)

	owner, group := "?", "?"
	if u, err := user.LookupId(strconv.Itoa(int(fileStats.Uid))); err == nil {
		owner = u.Username
	}
	if g, err := user.LookupGroupId(strconv.Itoa(int(fileStats.Gid))); err == nil {
		group = g.Name
	}

	fmt.Printf("%s %d %s %s %d %s %s\n",
		modeString(fileStats.Mode),
		fileStats.Nlink,
		owner,
		group,
		fileStats.Size,
		fileInfo.ModTime().Format("Jan _2 15:04"),
		name,
	)
}
