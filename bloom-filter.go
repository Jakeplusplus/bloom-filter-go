package main

import (
	"errors"
	"fmt"
	"hash/fnv"
)

const FilterSize = 1000
const HashRuns = 3
const exitCommand = "/quit"

type bloomFilter struct {
	filter [FilterSize]bool
}

func (b *bloomFilter) add(word string) error {
	if b.contains(word) {
		return errors.New("this word already exists in the bloom filter")
	}
	for i := range HashRuns {
		b.filter[b.hash(word, i)%FilterSize] = true
	}
	return nil
}

func (b bloomFilter) contains(word string) (contains bool) {
	r := false
	for i := range HashRuns {
		r = r || b.filter[b.hash(word, i)%FilterSize]
	}
	return r
}

func (b bloomFilter) hash(word string, seed int) int {
	h := fnv.New32a()
	h.Write([]byte{byte(seed)})
	h.Write([]byte(word))
	return int(h.Sum32())
}

func main() {
	var bloom bloomFilter
	var word string

	fmt.Print("Welcome to the bloom filter.  To exit please type \"/quit\"\n")

	for word != exitCommand {
		fmt.Print("Enter a word:\n")

		// Get user input
		fmt.Scanln(&word)

		if word != exitCommand {
			// Attempt to add word to bloom filter
			if err := bloom.add(word); err != nil {
				fmt.Printf("Failed to add word: %v!\n\n\n", err)
			} else {
				fmt.Printf("Added \"%s\".\n\n\n", word)
			}
		}
	}
}
