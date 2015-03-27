# Random

---

A Rand utility that wraps `math/rand`

## Installation

```bash
go get -u github.com/zabawaba99/random
```

## Usage

Import the package

```go
import "github.com/zabawaba99/random"
```

By importing the package, the package will be
seeded with `time.Now().Unix()`. You can also specify a seed
by setting the `GOSEED` environment variable.

```go
GOSEED=12 go test .
```
---

This package exports all the functions that [math/rand#Rand](http://godoc.org/math/rand#Rand)
exports.


Check the [GoDocs](http://godoc.org/github.com/zabawaba99/random) for more details

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b new-feature`)
3. Commit your changes (`git commit -am 'Some cool reflection'`)
4. Push to the branch (`git push origin new-feature`)
5. Create new Pull Request