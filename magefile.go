// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/sh"
)

// TODO: duplicate links, ..

// Find broken links
func Clean() {
	_, err := exec.LookPath("mdck")
	if err != nil {
		fmt.Println("mdck is not installed. install it https://github.com/ctm/mdck")
		os.Exit(2)
	}
	sh.RunV("mdck", ".")
}
