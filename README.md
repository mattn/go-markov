# go-markov

markov chain

## Usage

```go
m := markov.New()
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
	m.Update(strings.TrimSpace(scanner.Text()))
}
if err := scanner.Err(); err != nil {
	log.Fatal(err)
}
fmt.Println(m.Chain(m.First()))
```

## Installation

```
go get github.com/mattn/go-markov
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
