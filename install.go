package main

import (
	"github.com/making/gozouroppu/util"
	"log"
	"os"
)

func Install(baseDir string) {
	log.Println("Installing...")
	installDotEmacsD(baseDir)
	installDotProfile(baseDir)
}

func installDotEmacsD(baseDir string) {
	util.Mkdirp(os.Getenv("HOME") + "/" + dotEmacsD)
	util.Rmrf(os.Getenv("HOME") + "/" + initEl)
	util.Lns(baseDir + "/" + initEl, os.Getenv("HOME") + "/" + initEl)
}

func installDotProfile(baseDir string) {
	util.Rmrf(os.Getenv("HOME") + "/" + dotProfile)
	util.Lns(baseDir + "/" + dotProfile, os.Getenv("HOME") + "/" + dotProfile)
}
