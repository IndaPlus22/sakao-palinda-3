package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {
	freq := make(map[string]int)
	clean := regexp.MustCompile(`[,.]+`).ReplaceAllString(text, " ")
	words := strings.Fields(strings.ToLower(clean))
	runners := 20
	n := (len(words) + runners - 1) / (runners)

	ch := make(chan map[string]int, runners)
	wg := new(sync.WaitGroup)
	wg.Add(runners)

	for i := 0; i < runners; i++ {
		last := ((i + 1) * n)
		if last > len(words) {
			last = len(words)
		}
		sl := words[i*n : last]
		go func() {
			defer wg.Done()
			partialFreq := make(map[string]int)
			for _, word := range sl {
				partialFreq[word]++
			}
			ch <- partialFreq
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for {
		localFreq, ok := <-ch
		if !ok {
			break
		}
		for word, count := range localFreq {
			freq[word] += count
		}
	}
	return freq
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	data, err := ioutil.ReadFile("loremipsum.txt")
	if err != nil {
		log.Fatal(err)
	}

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
