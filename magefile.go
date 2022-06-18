// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/sh"
)

// TODO: duplicate links, broken links (https://github.com/tmcw/linkrot has way to check if link is bad)
// TODO: scan through all links in # Links or all links in wiki as whole
// TODO: find .md files that are not in SUMMARY.md
// TODO: run https://github.com/smallhadroncollider/brok on whole wiki
// TODO: remove all ?utm extra useless stuff added to links
// TODO: find broken links. Of this kind: `[]Link Title(https://..)` or similar
// TODO: jump over all imgur links & check if the URL is still up

// Find broken links
func Clean() {
	_, err := exec.LookPath("mdck")
	if err != nil {
		fmt.Println("mdck is not installed. install it https://github.com/ctm/mdck")
		os.Exit(2)
	}
	sh.RunV("mdck", ".")
}
