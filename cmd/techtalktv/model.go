package main

import (
	"log"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"os"
	"sort"
	"errors"
)

// Returns the location of the data folder
func GetRootFolder() (string, error) {
	root,err := os.Getwd()
	if (err == nil) {
		return filepath.Join(root, "events"), nil
	}
	return "", err
}

// Load the events model from a given location
func LoadModel(root string) (EventsIndex, error) {
	dirnames, _ := readSubdirNames(root)

	list := make(Events, 0)
	for _, dirname := range dirnames {
		fn := filepath.Join(root, dirname, "contents.json")
		log.Printf("file: [%s] ", fn)
		if data, err := ioutil.ReadFile(fn); err == nil {
			var arr Events
			e := json.Unmarshal(data, &arr)
			if (e != nil) {
				return nil, e
			}
			log.Printf("  evts: [%d] ", len(arr))
			for idx, _ := range arr {
				arr[idx].Base = dirname
			}
			list = append(list, arr...)
		} else {
			// skip
			log.Println(err)
		}
	}

	var index = make(map[string]Event)
	for _, event := range list {
		index[event.Id] = event
	}

	return index, nil
}

// Read the data file for a given Event, assumed to be accessible
// in a specified data repository
func ReadEventFile(root string, event Event) ([]byte, error) {
	fp := filepath.Join(root, event.Base, event.Index)
	if data, err := ioutil.ReadFile(fp); err == nil {
		//var s Events
		//e := json.Unmarshal(data, &s)
		return data, nil
	}
	return nil, errors.New("Cannot read file ")
}

// ----------------------------------------------------------------------------
// Internal

func readSubdirNames(parentDir string) ([]string, error) {
	f, err := os.Open(parentDir)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}
