package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// set the split function for the scanning
	scanner.Split(bufio.ScanWords)

	// create slice to hold counts
	buckets := make([]int, 200)
	// All values by default will be 0s

	for scanner.Scan() {
		// Get the words using scanner.Text()
		n := hashBucket(scanner.Text())

		// Increments the count
		buckets[n]++
	}

	fmt.Println(buckets[65:123])

}

// Use lower letter so it won't available outside of package.
func hashBucket(word string) int {
	return int(word[0])
}
