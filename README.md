# Random [![Build Status](https://travis-ci.org/zabawaba99/random.svg?branch=master)](https://travis-ci.org/zabawaba99/random) [![Coverage Status](https://coveralls.io/repos/zabawaba99/random/badge.svg?branch=master)](https://coveralls.io/r/zabawaba99/random?branch=master)

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