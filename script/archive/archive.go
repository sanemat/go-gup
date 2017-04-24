package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/mholt/archiver"
	"github.com/sanemat/go-gup/script/gitdescribetags"
)

func main() {
	tag, err := gitdescribetags.Get()
	if err != nil {
		log.Fatal(err)
	}
	dirPath := filepath.Join("pkg", tag)
	fileinfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range fileinfos {
		if !info.IsDir() {
			continue
		}
		if err := archiver.Zip.Make(filepath.Join(dirPath, info.Name()+".zip"), []string{filepath.Join(dirPath, info.Name())}); err != nil {
			log.Fatal(err)
		}
	}
}
