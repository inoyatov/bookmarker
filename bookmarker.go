package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type Bookmarks map[string]string

func NewBookmarks() (bm Bookmarks) {
	return make(Bookmarks)
}

type ActionsForBookmarks interface {
	Add(key, path string) (ok bool)
	Modify(key, path string) (ok bool)
	Delete(key string) (ok bool)
	DleteAll()
	Get(key string) (path string, err error)
	IsExists(key string) (exists bool)
	Print()
}

func (bm *Bookmarks) Modify(key, path string) (ok bool) {
	if !bm.IsExists(key) {
		(*bm)[key] = path
		return true
	}
	return false
}

func (bm Bookmarks) Print() {
	var keys []string
	var max_len_of_key int
	max_len_of_key = 0

	for key := range bm {
		keys = append(keys, key)

		len_of_key := len(key)
		if max_len_of_key < len_of_key {
			max_len_of_key = len_of_key
		}
	}

	sort.Strings(keys)

	str_max_len_of_key := strconv.FormatInt(int64(max_len_of_key), 10)
	if str_max_len_of_key == "" {
		str_max_len_of_key = "10"
	}

	format := "%-" + str_max_len_of_key + "s - %s\n"
	for _, key := range keys {
		fmt.Printf(format, key, bm[key])
	}
}

func (bm *Bookmarks) Add(key, path string) (ok bool) {

	re, err := regexp.Compile(`^[0-9A-Za-z]+$`)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if !re.MatchString(key) {
		return false
	}

	if !bm.IsExists(key) {
		(*bm)[key] = path
		return true
	}
	return false
}

func (bm Bookmarks) IsExists(key string) (exists bool) {
	if _, ok := bm[key]; ok {
		return true
	}

	return false
}

func (bm *Bookmarks) Delete(key string) (ok bool) {
	if bm.IsExists(key) {
		delete(map[string]string(*bm), key)
		return true
	}

	return false
}

func (bm *Bookmarks) DeleteAll() {
	*bm = NewBookmarks()
}

func (bm Bookmarks) Get(key string) (path string, err error) {
	if value, ok := bm[key]; ok {
		err = nil
		path = value
	} else {
		err = errors.New("No such bookmark path.")
		path = ""
	}

	return path, err
}

func main() {

	bm := NewBookmarks()

	fmt.Println("Step 1 --------------------------")
	bm.Add("home", "/home/phoenix/")
	bm.Add("projects", "/home/phoenix/Documents/Projects/")
	bm.Print()

	fmt.Println("Step 2 --------------------------")
	bm.DeleteAll()
	bm.Print()

	fmt.Println("Step 3 --------------------------")
	bm.Add("zoo", "/home/phoenix/Zoo/")
	bm.Add("projects*", "/home/phoenix/Documents/Projects/")
	bm.Add("home", "/home/phoenix/")
	bm.Add("projects-test-hello* asd", "/home/phoenix/Documents/Projects/")
	bm.Print()

	fmt.Println("Step 4 --------------------------")
	bm.Delete("home")
	bm.Print()

	fmt.Println("Step 5 --------------------------")
	bm.Add("projects-test-hello asd", "/home/phoenix/Documents/Projects/")
	bm.Print()
}
