package random

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSeedNoEnv(t *testing.T) {
	tmp := os.Getenv("GOSEED")
	os.Setenv("GOSEED", "")
	gblRand = nil

	seed()
	require.NotNil(t, gblRand)

	sample := rand.New(rand.NewSource(randSeed))
	assert.Equal(t, sample.ExpFloat64(), ExpFloat64())
	assert.Equal(t, sample.Float32(), Float32())
	assert.Equal(t, sample.Float64(), Float64())
	assert.Equal(t, sample.Int(), Int())
	assert.Equal(t, sample.Intn(42), Intn(42))
	assert.Equal(t, sample.Int31(), Int31())
	assert.Equal(t, sample.Int31n(42), Int31n(42))
	assert.Equal(t, sample.Int63(), Int63())
	assert.Equal(t, sample.Int63n(42), Int63n(42))
	assert.Equal(t, sample.NormFloat64(), NormFloat64())
	assert.Equal(t, sample.Perm(42), Perm(42))
	assert.Equal(t, sample.Uint32(), Uint32())

	os.Setenv("GOSEED", tmp)
}

func TestSeedEnv(t *testing.T) {
	seedVal := int64(42)

	tmp := os.Getenv("GOSEED")
	os.Setenv("GOSEED", strconv.FormatInt(seedVal, 10))
	gblRand = nil

	seed()
	require.NotNil(t, gblRand)

	sample := rand.New(rand.NewSource(seedVal))
	assert.Equal(t, sample.ExpFloat64(), ExpFloat64())
	assert.Equal(t, sample.Float32(), Float32())
	assert.Equal(t, sample.Float64(), Float64())
	assert.Equal(t, sample.Int(), Int())
	assert.Equal(t, sample.Intn(42), Intn(42))
	assert.Equal(t, sample.Int31(), Int31())
	assert.Equal(t, sample.Int31n(42), Int31n(42))
	assert.Equal(t, sample.Int63(), Int63())
	assert.Equal(t, sample.Int63n(42), Int63n(42))
	assert.Equal(t, sample.NormFloat64(), NormFloat64())
	assert.Equal(t, sample.Perm(42), Perm(42))
	assert.Equal(t, sample.Uint32(), Uint32())

	os.Setenv("GOSEED", tmp)
}

func TestString(t *testing.T) {
	seedVal := int64(42)

	tmp := os.Getenv("GOSEED")
	os.Setenv("GOSEED", strconv.FormatInt(seedVal, 10))
	gblRand = nil

	seed()
	require.NotNil(t, gblRand)

	size := 42
	assert.Len(t, String(size), size)

	os.Setenv("GOSEED", tmp)
}
