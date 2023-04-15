package markov

import (
	"math/rand"
	"regexp"
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

func (m *Markov) UpdateLine(text string) error {
	second, ok := m.tbl[text]
	if !ok {
		second = make(map[string][]string)
		m.tbl[text] = second
	}
	return nil
}

func (m *Markov) First() string {
	keys := []string{}
	for k := range m.tbl {
		keys = append(keys, k)
	}
	if len(keys) == 0 {
		return ""
	}
	return keys[rand.Int()%len(keys)]
}

func (m *Markov) Chain(first string) string {
	text := first

	keys := []string{}
	for k := range m.tbl[first] {
		keys = append(keys, k)
	}
	if len(keys) == 0 {
		return ""
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
