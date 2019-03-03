package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	parser "github.com/nikitavoloboev/markdown-parser"
)

func getMdFiles(path string) []string {
	var fileNames []string
	files, err := filepath.Glob(path + "/*.md")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fileNames = append(fileNames, f)
	}
	dir, err := ioutil.ReadDir(path + "/")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range dir {
		if f.IsDir() && f.Name() != ".git" {
			fileNames = append(fileNames, getMdFiles(path+"/"+f.Name())...)
		}
	}
	return fileNames
}

func main() {
	// Get all of the existing MD files
	fileNames := getMdFiles(".")
	// Get all of the image links and download them from all of the MD files
	for _, f := range fileNames {
		fileHandle, _ := os.Open("./" + f)
		defer fileHandle.Close()
		scanner := bufio.NewScanner(fileHandle)
		for scanner.Scan() {
			link := parser.ParseImageLink(scanner.Text())
			if link != "" {
				fmt.Println(link)
				cmd := exec.Command("wget", link, "-P", "./imgs")
				cmd.Stdout = os.Stdout
				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
