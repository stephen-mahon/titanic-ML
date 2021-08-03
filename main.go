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
	for scanner.Scan() {
		if i != 0 && i < 2 {
			//var name string
			s := strings.Split(scanner.Text(), ",")
			for j := range s {
				// Want to keep the name together even though it has a comma
				fmt.Println(s[j])
			}
		}
		i++

	}
}
