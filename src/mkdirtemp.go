package main

import (
	"fmt"
	"os"
	//"path/filepath"
)

func main() {
	dirname, err := os.MkdirTemp("", "testdir")
	defer os.RemoveAll(dirname)
	if err != nil {
		panic(err)
	}
	fmt.Printf("dirname: %s\n", dirname)


}
