package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)
	buckets := make(map[int]map[string]int)

	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int)
	}
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		buckets[n][word]++
	}
	for k, v := range buckets {
		fmt.Println(v, "\t", k)
	}
}

func hashBucket(word string, bucket int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % bucket
}
