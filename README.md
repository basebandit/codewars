# Codewars

[![CodeWars](https://www.codewars.com/users/devmars/badges/large)](https://www.codewars.com/users/devmars)

> Practice make Perfect

- :tada: Unit test with `testing` package
- :zap: Simple code
- :art: Readable code

### Getting Started

```Go
go get https://github.com/basebandit/codewars.git
cd codewars
```

#### Running tests

```Go
go test -v ./kata1
go test -v ./kata2
```

#### Running benchmarks

```Go
 go test -run=^$ -cpu=1,2,4 -bench=. ./kata1
 go test -run=^$ -cpu=1,2,4 -bench=. ./kata2
```
