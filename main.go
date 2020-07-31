package main

import (
	"os"
	"path/filepath"
	"fmt"
	"io/ioutil"
)

func findExecutable(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return nil
	}
	return os.ErrPermission
}

func main() {
	path := os.Getenv("PATH")
	for _, dir := range filepath.SplitList(path) {
		if dir == "" {
			// Unix shell semantics: path element "" means "."
			dir = "."
		}
		fmt.Println("SEARCHING", dir)
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			continue;
		}
		for _, f := range files {
			fmt.Printf(f.Name()+", ")
		}
		fmt.Printf("\n\n")
		// path := filepath.Join(dir, file)
		// if err := findExecutable(path); err == nil {
		// 	return path, nil
		// }
	}
}