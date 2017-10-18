package markov

import (
	"math/rand"
	"regexp"

	"github.com/ikawaha/kagome/tokenizer"
)

var (
	reIgnoreText = regexp.MustCompile(`[\[\]「」『』()]`)
)

type Markov struct {
	tbl map[string]map[string][]string
}

func New() *Markov {
	return &Markov{
		tbl: make(map[string]map[string][]string),
	}
}

func (m *Markov) Update(text string) {
	t := tokenizer.New()
	text = reIgnoreText.ReplaceAllString(text, "")
	tokens := t.Tokenize(text)

	words := []string{}
	for _, token := range tokens {
		if token.Surface == "EOS" {
			continue
		}
		words = append(words, token.Surface)
	}

	size := len(words)

	for i := 0; i < size-2; i++ {
		second, ok := m.tbl[words[i]]
		if !ok {
			second = make(map[string][]string)
			m.tbl[words[i]] = second
		}
		second[words[i+1]] = append(second[words[i+1]], words[i+2])
	}
}

func (m *Markov) First() string {
	keys := []string{}
	for k := range m.tbl {
		keys = append(keys, k)
	}
	return keys[rand.Int()%len(keys)]
}

func (m *Markov) Chain(first string) string {
	text := first

	keys := []string{}
	for k := range m.tbl[first] {
		keys = append(keys, k)
	}
	kv := rand.Int() % len(keys)
	second := keys[kv]
	text += second

	for second != "" {
		size := len(m.tbl[first][second])
		if size == 0 {
			break
		}
		idx := rand.Int() % size
		next := m.tbl[first][second][idx]
		text += next
		first = second
		second = next
	}
	return text
}
