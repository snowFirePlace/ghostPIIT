package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	dictionary map[string]map[string]map[string]string
)

func transiteFile(path, file string) {
	var i int
	defer fmt.Printf("%d:%s\n",i,file)

	mainSection := file

	if strings.ContainsAny(file, string(os.PathSeparator)) {
		sep := strings.LastIndex(file, string(os.PathSeparator))
		path = path + string(os.PathSeparator) + file[:sep]
		file = file[sep+1:]
		// sep := strings.Split(file, string(os.PathSeparator))
		// path = path + string(os.PathSeparator) + sep[0]
		// file = sep[1]
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("Directory %s not find", path)
		return
		log.Fatal(err)
	}
	for _, f := range files {
		// fmt.Println(file, f.Name())
		// fmt.Println(regexp.MatchString(file, f.Name()))
		name, err := regexp.MatchString(file, f.Name())
		if name {
			file = f.Name()
			i+=1
			break
		}
		if err != nil {
			log.Fatal(err)
		}

	}
	fmt.Println("File: ", path+string(os.PathSeparator)+file)
	data, err := os.ReadFile(path + string(os.PathSeparator) + file)
	if err != nil {
		log.Fatal(err)
	}
	out := string(data)
	for o, m := range dictionary[mainSection] {
		fmt.Printf("\tSection: %s\n", o)
		for key, word := range m {
			if strings.Count(out, fmt.Sprintf(o, key)) == 0 {
				fmt.Printf("\t\t\"%s\"\n", key)
				// delete(dictionary[mainSection], key)
				continue
			}
			out = strings.Replace(out, fmt.Sprintf(o, key), fmt.Sprintf(o, word), -1)
		}
		err = ioutil.WriteFile(path+string(os.PathSeparator)+file, []byte(out), 0644)
		if err != nil {
			log.Println("Error creating", path+string(os.PathSeparator)+file)
			return
		}
	}
}
func readJson(f string) {
	fileGhostJSON, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fileGhostJSON, &dictionary)
	if err != nil {
		log.Fatal(err)
	}
}

func writeJson(f string) {
	os.Remove(f)
	file, _ := json.MarshalIndent(dictionary, "", " ")
	_ = ioutil.WriteFile(f, file, 0644)
}

func main() {
	dictionary = make(map[string]map[string]map[string]string)
	// * For windows
	// readJson("C:\\docker\\kw\\golang\\src\\ru.json")
	// * For docker
	readJson("/app/ru.json")

	for file, _ := range dictionary {
		// * For windows
		// transiteFile("C:\\docker\\kw\\ghost\\src", file)
		// * For docker
		transiteFile("/temp", file)
	}
	// * For windows
	// writeJson("C:\\docker\\kw\\golang\\src\\ru.json")
	// * For docker
	// writeJson("./ru.json")
}
