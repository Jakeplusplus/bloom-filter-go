package main

import (
	"errors"
	"fmt"
	"hash/fnv"
)

// Configuration
const FilterSize = 1000     // FilterSize defines the size of the bloom filter bit array
const HashRuns = 3          // HashRuns defines the number of hash functions to use
const ExitCommand = "/quit" // ExitCommand is a string that will cause the program to close

type bloomFilter struct {
	filter [FilterSize]bool // Bit array to store the filter state
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
	// TODO: Replace with a hash function that has an actual seeding option.  Maybe Murmur3.
	// If sticking to the Go standard library, I think maphash might be the best from the standard library.
	h := fnv.New32a()
	h.Write([]byte{byte(seed)})
	h.Write([]byte(word))
	return int(h.Sum32())
}

func main() {
	var bloom bloomFilter
	var word string

	fmt.Print("Welcome to the bloom filter.  To exit please type \"/quit\"\n")

	for word != ExitCommand {
		fmt.Print("Enter a word:\n")

		// Get user input
		fmt.Scanln(&word)

		if word != ExitCommand {
			// Attempt to add word to bloom filter
			if err := bloom.add(word); err != nil {
				fmt.Printf("Failed to add word: %v!\n\n\n", err)
			} else {
				fmt.Printf("Added \"%s\".\n\n\n", word)
			}
		}
	}
}
