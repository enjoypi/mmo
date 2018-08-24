package ext

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func BenchmarkRandomUint64(b *testing.B) {
	m := make(map[uint64]bool)
	for i := 0; i < b.N; i++ {
		m[RandomUint64()] = true
	}
	AssertB(b, b.N == len(m))
}

func BenchmarkRandomInt64(b *testing.B) {
	m := make(map[int64]bool)
	for i := 0; i < b.N; i++ {
		m[RandomInt64()] = true
	}
	AssertB(b, b.N == len(m))
}

func TestEncrypt(t *testing.T) {
	key := []byte("1234567890123456")
	plain := []byte("abcdefghijklmn")

	cipher, err := CBCEncrypt(key, plain)
	assert.NoError(t, err)

	p, err := CBCDecrypt(key, cipher)
	assert.NoError(t, err)

	assert.EqualValues(t, plain, p)
}

func TestRandomIntn(t *testing.T) {
	m := make(map[int]int)
	max := 10
	n := 1000000
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		m[rand.Intn(max)] += 1
	}

	assert.Equal(t, max, len(m))
	minAmount := int(math.MaxInt64)
	maxAmount := int(math.MinInt64)
	total := int(0)
	for _, amount := range m {
		total += amount
		if amount < minAmount {
			minAmount = amount
		}
		if amount > maxAmount {
			maxAmount = amount
		}
	}
	fmt.Println(t.Name(), maxAmount, minAmount, float64((maxAmount-minAmount)*max)/float64(n))
}
