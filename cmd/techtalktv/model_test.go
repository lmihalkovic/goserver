package main

import (

	"testing"

	"log"
	"fmt"
	"os"
)

func init() {
	path := os.Args[2]
	if path != "" {
		os.Chdir(path)
	}

	fmt.Printf("Root: [%s]", path)
}

func TestFolders(t *testing.T) {
	root,err := GetRootFolder()
	if (err != nil) {
		log.Fatal("Cannot find working directory")
	}

	fmt.Print(root)
	dirnames, _ := readSubdirNames(root)
	for index, fn := range dirnames {
		t.Logf("%d : %s", index, fn)
	}

	//// Just for fun
	//StringArray(dirnames).ForEach(func(index int, val string) bool {
	//	t.Logf("%d : %s", index, val)
	//	return true
	//})
}

func TestLoad(t *testing.T) {

	root,err := GetRootFolder()
	if (err != nil) {
		log.Fatal("Cannot find working directory")
	}

	index, err := LoadModel(root)
	if(err != nil) {
		panic(err)
	}

	for id, event := range index {
		log.Printf("%s -> %s [%s]", id, event, event.Base)
	}

}
