package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	rs, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(rs.Body)
	rs.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", data)
}
