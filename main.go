package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	root := "./data/"
	files, err := ioutil.ReadDir(root)
	check(err)

	var filenames []string
	for _, f := range files {
		filenames = append(filenames, root+f.Name())
	}

	file, err := os.Open(filenames[1])
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i int
	var list []string
	for scanner.Scan() {
		if i != 0 && i < 2 {
			//var name string
			s := strings.Split(scanner.Text(), ",")
			var temp string
			for j := range s {
				if s[j] != "" {
					list = append(list, s[j])
				} else if s[j] != "" && (string(s[j][0]) == "\"" || string(s[j][len(s[j])-1]) == "\"") {
					// Want to keep the name together even though it has a comma
					temp += s[j]
				} else if len(temp) > 1 {
					list = append(list, temp)
				} else {
					list = append(list, s[j])
				}
			}
		}
		i++

	}
	for _, v := range list {
		fmt.Printf("%s\n", v)
	}
}
