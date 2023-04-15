package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ikawaha/kagome-dict/uni"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"github.com/mattn/go-markov"
)

func contains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	var length int
	flag.IntVar(&length, "n", -1, "letters")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	m := markov.New()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m.Update(strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	t, err := tokenizer.New(uni.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		log.Fatal(err)
	}

	bad := []string{
		"助詞",
		"補助記号",
	}
	var result string
	for {
		var first string
		for {
			first = m.First()
			tokens := t.Tokenize(first)
			if !contains(bad, tokens[0].Features()[0]) {
				break
			}
		}

		result = strings.TrimSpace(m.Chain(first))
		if result != "" && (length == -1 || len([]rune(result)) <= length) {
			break
		}
	}
	fmt.Println(result)
}
