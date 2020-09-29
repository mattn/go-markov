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

	"github.com/mattn/go-markov"
)

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

	var result string
	for {
		result = strings.TrimSpace(m.Chain(m.First()))
		if result != "" && (length == -1 || len([]rune(result)) <= length) {
			break
		}
	}
	fmt.Println(result)
}
