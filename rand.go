package random

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

var seeder *rand.Rand

func init() {
	seed()
}

func seed() {
	seed := time.Now.Unix()
	if env := os.Getenv("GOSEED"); env != "" {
		t, err := strconv.ParseInt(env, 10, 64)
		if err != nil {
			seed = t
		}
	}

	seeder = rand.New(rand.NewSource(seed))
}

// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1). To produce a distribution with
// a different rate parameter, callers can adjust the output using:
//
//		sample = ExpFloat64() / desiredRateParameter
func ExpFloat64() float64 { return seeder.ExpFloat64() }

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func Float32() float32 { return seeder.Float32() }

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func Float64() float64 { return seeder.Float64() }

// Int returns a non-negative pseudo-random int.
func Int() int { return seeder.Int() }

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Intn(n int) int { return seeder.Intn(n) }

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func Int31() int32 { return seeder.Int31() }

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Int31n(n int32) int32 { return seeder.Int31n(n) }

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func Int63() int64 { return seeder.Int63() }

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Int63n(n int64) int64 { return seeder.Int63n(n) }

// NormFloat64 returns a normally distributed float64 in the range
// [-math.MaxFloat64, +math.MaxFloat64] with standard normal distribution
// (mean = 0, stddev = 1). To produce a different normal distribution,
// callers can adjust the output using:
//
//		sample = NormFloat64() * desiredStdDev + desiredMean
func NormFloat64() float64 { return seeder.NormFloat64() }

// Perm returns, as a slice of n ints, a pseudo-random
// permutation of the integers [0,n).
func Perm(n []int) []int { return seeder.Perm(n) }

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func Uint32() uint32 { return seeder.Uint32() }
