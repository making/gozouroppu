package main

import (
	"github.com/making/gozouroppu/util"
	"log"
	"os"
)

func Clean(baseDir string) {
	log.Println("Cleaning...")
	cleanDotEmacsD()
	cleanDotProfile()
	cleanDotfiles(baseDir)
}

func cleanDotEmacsD() {
	util.Rmrf(os.Getenv("HOME") + "/" + dotEmacsD)
}

func cleanDotProfile() {
	util.Rmrf(os.Getenv("HOME") + "/" + dotProfile)
}

func cleanDotfiles(baseDir string) {
	if util.Exists(baseDir + "/.git") {
		log.Fatalln(baseDir, " cannot be removed because .git exists")
	}
	util.Rmrf(baseDir)
}
