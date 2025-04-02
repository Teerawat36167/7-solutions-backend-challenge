package counter

import (
	"regexp"
	"strings"
	"sync"
)

type MeatCounter struct {
	nonWordRegex *regexp.Regexp
}

func NewMeatCounter() *MeatCounter {
	return &MeatCounter{
		nonWordRegex: regexp.MustCompile(`[^\w-]+|_`),
	}
}

func (mc *MeatCounter) CountMeats(text string) map[string]int {
	text = strings.ToLower(text)
	text = mc.nonWordRegex.ReplaceAllString(text, " ")
	words := strings.Fields(text)
	return mc.countWordsParallel(words)
}

func (mc *MeatCounter) countWordsParallel(words []string) map[string]int {
	numGoroutines := 4
	if len(words) < 1000 {
		numGoroutines = 1
	}

	wordsPerGoroutine := (len(words) + numGoroutines - 1) / numGoroutines

	var wg sync.WaitGroup
	var mutex sync.Mutex
	results := make(map[string]int)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		start := i * wordsPerGoroutine
		end := (i + 1) * wordsPerGoroutine
		if end > len(words) {
			end = len(words)
		}

		go func(start, end int) {
			defer wg.Done()

			localCounts := make(map[string]int)
			for j := start; j < end; j++ {
				word := words[j]
				if word != "" {
					localCounts[word]++
				}
			}

			mutex.Lock()
			for word, count := range localCounts {
				results[word] += count
			}
			mutex.Unlock()
		}(start, end)
	}

	wg.Wait()
	return results
}

func (mc *MeatCounter) GetBeefCounts(text string) map[string]int {
	return mc.CountMeats(text)
}
