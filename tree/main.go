package main

import (
	"fmt"
	"io"
	"os"
	"io/ioutil"
        "strconv"
)

func getCurrentDirs(path string) []os.FileInfo {
        items, _ := ioutil.ReadDir(path)
        result := make([]os.FileInfo, 0)
        for _, singleItem := range items {
                if singleItem.IsDir() {
                        result = append(result, singleItem)
                }
        }
        return result
}

func getFileSize(path string) string {
        file, _ := os.Stat(path)
        if file.Size() == 0 {
                return " (empty)"
        }
        size := strconv.Itoa(int(file.Size()))
        return " (" + size + "b)"
}

func dirAnother(out io.Writer, path string, withFiles bool, tabsString string) error {
        container := getCurrentDirs(path)
        if withFiles {
                container, _ = ioutil.ReadDir(path)
        }
        for index, singleItem := range container {
                prefixChar := ""
                tabsStringAnother := tabsString
                if index == len(container) - 1 {
                        prefixChar = "└───"
                        tabsStringAnother += "\t"
                } else {
                        prefixChar = "├───"
                        tabsStringAnother += "│\t"
                }
                resultString := singleItem.Name()
                if !singleItem.IsDir() {
                        resultString += getFileSize(path + "/" + singleItem.Name())
                }
                fmt.Fprintln(out, tabsString + prefixChar + resultString)
                dirAnother(out, path + "/" + singleItem.Name(), withFiles, tabsStringAnother)
        }
        return nil
}

func dirTree(out io.Writer, path string, withFiles bool) error {
        dirAnother(out, path, withFiles, "")
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
