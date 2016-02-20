package util

import (
	"log"
	"os"
)

func Exists(path string) bool {
	info, _ := os.Stat(path)
	// log.Println("Stat", path, info, err)
	return (info != nil)
}

func Mkdirp(path string) {
	if !Exists(path) {
		log.Println("mkdir -p", path)
		if err := os.MkdirAll(path, 0755); err != nil {
			log.Fatalln("OMG!", err)
		}
	}
}

func Rmrf(path string) {
	if path == "/" || path == "~/" || path == "~" || path == os.Getenv("HOME") {
		log.Fatalln(path, " cannot be removed by this tool")
	}

	if Exists(path) {
		log.Println("rm -rf", path)
		if err := os.RemoveAll(path); err != nil {
			log.Fatalln("OMG!", err)
		}
	}
}

func Lns(source, target string) {
	log.Println("ln -s", source, target)
	if !Exists(source) {
		log.Println(source, "not exist!")
		return
	}
	if err := os.Symlink(source, target); err != nil {
		log.Fatalln("OMG!", err)
	}
}

func Mv(source, target string) {
	log.Println("mv", source, target)
	if !Exists(source) {
		log.Fatalln(source, "not exist!")
	}
	if err := os.Rename(source, target); err != nil {
		log.Fatalln("OMG!", err)
	}
}
