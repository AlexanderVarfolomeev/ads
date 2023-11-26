package main

import "os"

func allFiles(path string) []string {
	files := make([]string, 0)
	dir, _ := os.ReadDir(path)
	for _, entry := range dir {
		if entry.IsDir() {
			files = append(files, allFiles(entry.Name())...)
		} else {
			files = append(files, entry.Name())
		}
	}

	return files
}

func main() {
	l := allFiles("/Users/a.o.varfolomeev/Downloads/")
	println(l[0])
}
