package main

import (
	"fmt"
	"os"

	"github.com/landlock-lsm/go-landlock/landlock"
)

func main() {
	dir := "subdir"
	os.Mkdir(dir, 0700)

	err := landlock.V1.RestrictPaths(landlock.RODirs(dir))
	if err != nil {
		panic(err)
	}

	d, err := os.ReadDir(dir)
	fmt.Printf("os.ReadDir(%q) = %v, %v\n", dir, d, err)
}
