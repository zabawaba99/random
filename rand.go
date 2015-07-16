package random

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	gblRand  *rand.Rand
	randSeed int64
)

func init() {
	seed()
}

func seed() {
	randSeed = time.Now().UnixNano()
	if env := os.Getenv("GOSEED"); env != "" {
		t, err := strconv.ParseInt(env, 10, 64)
		if err == nil {
			randSeed = t
		}
	}

	gblRand = rand.New(rand.NewSource(randSeed))
}

var stringRune = []rune(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

// String returns an alphanumeric string of length n
func String(n int) string {
	parts := make([]rune, n)
	for i, _ := range parts {
		parts[i] = stringRune[Intn(n)]
	}
	return string(parts)
}

// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1). To produce a distribution with
// a different rate parameter, callers can adjust the output using:
//
//		sample = ExpFloat64() / desiredRateParameter
func ExpFloat64() float64 { return gblRand.ExpFloat64() }

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func Float32() float32 { return gblRand.Float32() }

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func Float64() float64 { return gblRand.Float64() }

// Int returns a non-negative pseudo-random int.
func Int() int { return gblRand.Int() }

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Intn(n int) int { return gblRand.Intn(n) }

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func Int31() int32 { return gblRand.Int31() }

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Int31n(n int32) int32 { return gblRand.Int31n(n) }

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func Int63() int64 { return gblRand.Int63() }

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Int63n(n int64) int64 { return gblRand.Int63n(n) }

// NormFloat64 returns a normally distributed float64 in the range
// [-math.MaxFloat64, +math.MaxFloat64] with standard normal distribution
// (mean = 0, stddev = 1). To produce a different normal distribution,
// callers can adjust the output using:
//
//		sample = NormFloat64() * desiredStdDev + desiredMean
func NormFloat64() float64 { return gblRand.NormFloat64() }

// Perm returns, as a slice of n ints, a pseudo-random
// permutation of the integers [0,n).
func Perm(n int) []int { return gblRand.Perm(n) }

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func Uint32() uint32 { return gblRand.Uint32() }
