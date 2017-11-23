package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dhowden/tag"
)

func title(oggFileName string) string {
	f, err := os.Open(oggFileName)
	if err != nil {
		log.Fatal(err)
	}

	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}

	if t := m.Title(); len(t) > 0 {
		return t
	}

	return strings.Replace(filepath.Base(oggFileName), ".ogg", "", 1)
}
