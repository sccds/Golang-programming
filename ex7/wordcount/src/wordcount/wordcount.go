package wordcount

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Pair struct {
	Key   string
	Value int
}

// PairList implement sort interface, use sort.Sort to sort
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[j].Value < p[i].Value }

// get words
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool {
		return !unicode.IsLetter(char)
	}
	return strings.FieldsFunc(s, notALetter)
}

// implement wordcount,
// methods: Merge(), Report(), SortReport(), UpdateFreq(), WordFreqCounter()

type WordCount map[string]int

// merge two WordCount
func (source WordCount) Merge(wordcont WordCount) WordCount {
	for k, v := range wordcont {
		source[k] += v
	}
	return source
}

// print word frequency
func (wordcount WordCount) Report() {
	words := make([]string, 0, len(wordcount))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range wordcount {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth, wordcount[word])
	}
}

// print word frequency in decending order
func (wordcount WordCount) SortReport() {
	p := make(PairList, len(wordcount))
	i := 0
	for k, v := range wordcount {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	wordWidth, frequencyWidth := 0, 0
	for _, pair := range p {
		word, frequency := pair.Key, pair.Value
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")

	for _, pair := range p {
		fmt.Printf("%-*s %*d\n", wordWidth, pair.Key, frequencyWidth, pair.Value)
	}

}

// read words from  file, get frequency
func (wordcount WordCount) UpdateFreq(filename string) {
	var file *os.File
	var err error

	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				wordcount[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Println("fialed to finish reading the file: ", err)
			}
			break
		}
	}
}

// concurrently get word count
func (wordcount WordCount) WordFreqCounter(files []string) {
	results := make(chan Pair, len(files))
	done := make(chan struct{}, len(files))

	for i := 0; i < len(files); i++ {
		go func(done chan<- struct{}, results chan<- Pair, filename string) {
			wordcount := make(WordCount)
			wordcount.UpdateFreq(filename)
			for k, v := range wordcount {
				pair := Pair{k, v}
				results <- pair
			}
			done <- struct{}{}
		}(done, results, files[i])
	}

	for working := len(files); working > 0; {
		select {
		case pair := <-results:
			wordcount[pair.Key] += pair.Value
		case <-done:
			working--
		}
	}

DONE:
	for {
		select {
		case pair := <-results:
			wordcount[pair.Key] += pair.Value
		default:
			break DONE
		}
	}
	close(results)
	close(done)
}
