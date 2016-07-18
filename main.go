package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"log"
	"os"
	"path"
	"sort"
)

// TEMP_BOOKMARKER_FILE = "/tmp/bookmarker.bm"
var (
	TMP_FOLDER = "/tmp"
	TMP_FILE   = "bookmarks.bm"
)

type Bookmarks map[int]string

func NewBookmarks() Bookmarks {
	return make(Bookmarks)
}

func ReadBookmarks(file_name, file_path string) Bookmarks {
	file_path_and_name := path.Join(file_path, file_name)
	if _, err := os.Stat(file_path_and_name); os.IsNotExist(err) {
		log.Println("Path does not exists")
	} else {
		log.Println("File exists")
	}

	return NewBookmarks()
}

func PrintBookmarks(bookmarks Bookmarks) {
	if len(bookmarks) > 0 {
		var keys []int

		for key := range bookmarks {
			keys = append(keys, key)
		}

		sort.Ints(keys)

		for _, key := range keys {
			log.Printf("%d - %s\n", key, bookmarks[key])
		}
	}
}

func ToGOB64(bookmarks Bookmarks) (strBookmarks string) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(bookmarks)
	if err != nil {
		log.Println("failed gob Encode", err)
	}
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func FromGOB64(strBookmarks string) (bookmarks Bookmarks) {
	bookmarks = NewBookmarks()
	by, err := base64.StdEncoding.DecodeString(strBookmarks)
	if err != nil {
		log.Println("failed base64 Decode", err)
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&bookmarks)
	if err != nil {
		log.Println("failed gob Decode", err)
	}

	return bookmarks
}

func main() {
	log.Println("=== TEST ====")
	bookmarks := NewBookmarks()

	PrintBookmarks(bookmarks)
	bookmarks[1] = "test"
	bookmarks[2] = "mest"

	PrintBookmarks(bookmarks)
	strBookmarks := ToGOB64(bookmarks)
	log.Println(strBookmarks)
	PrintBookmarks(FromGOB64(strBookmarks))

	ReadBookmarks(TMP_FOLDER, TMP_FILE)
}
