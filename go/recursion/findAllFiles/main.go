package main

import "os"

func AllFiles(path string) []string {
	dir, _ := os.ReadDir(path)
	return allFiles(dir, make([]string, 0))
}

func allFiles(dir []os.DirEntry, files []string) []string {
	for _, entry := range dir {
		if entry.IsDir() {
			subDir, _ := os.ReadDir(entry.Name())
			files = append(files, allFiles(subDir, files)...)
		} else {
			files = append(files, entry.Name())
		}
	}

	return files
}

func main() {
}
