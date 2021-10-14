package Jugo

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	if cwdPath, err := os.Getwd(); err != nil {
		log.Fatal(err)
		return
	} else {
		if _, err := os.Lstat(filepath.Join(cwdPath, "/result")); err != nil {
			// mkdir result
			if res := os.Mkdir(filepath.Join(cwdPath, "/result"), os.ModeDir); res != nil {
				log.Fatal(res)
			}
			// mkdir log
			if res := os.Mkdir(filepath.Join(cwdPath, "/log"), os.ModeDir); res != nil {
				log.Fatal(res)
			}
		}
	}
}
