package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/mattn/go-markov"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	m := markov.New()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m.Update(strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.Chain(m.First()))
}
