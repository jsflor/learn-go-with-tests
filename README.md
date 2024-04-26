# Learn go with tests

[link](https://quii.gitbook.io/learn-go-with-tests)

## Running tests

Run tests

```bash
    go test
```

Run benchmarks

```bash
    go test -bench=. 
```

## Godoc

Install go doc util

```bash
    go install golang.org/x/tools/cmd/godoc@latest
```

Run it

```bash
    godoc -http=:6060
```

Check documentation and examples in this [url](http://localhost:6060/pkg/example/)