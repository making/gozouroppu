package main

import (
	"flag"
	"github.com/making/gozouroppu/util"
	"os"
	"strings"
)

const (
	dotfilesZip = "dotfiles.zip"
	dotEmacsD   = ".emacs.d"
	initEl      = dotEmacsD + "/init.el"
	dotProfile  = ".profile"
)

func main() {
	install := flag.Bool("i", false, "Install")
	clean := flag.Bool("c", false, "Clean")
	baseDir := flag.String("d", os.Getenv("PWD")+"/dotfiles", "Dotfiles directory")
	remoteZip := flag.String("z", "https://github.com/making/dotfiles/archive/master.zip", "Remote zip")
	flag.Parse()

	if *clean {
		Clean(*baseDir)
	} else if !util.Exists(*baseDir) {
		downloadDotfiles(*remoteZip, *baseDir)
	}

	if *install {
		Install(*baseDir)
	}
}

func downloadDotfiles(remoteZip, baseDir string) {
	tokens := strings.Split(remoteZip, "/")
	unzippedName := tokens[len(tokens)-3] + "-" + strings.Replace(tokens[len(tokens)-1], ".zip", "", 1)
	util.Wget(remoteZip, dotfilesZip)
	util.Unzip(dotfilesZip, ".")
	util.Mv(unzippedName, baseDir)
	util.Rmrf(dotfilesZip)
}
