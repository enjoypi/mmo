package ext

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"math/big"
	"os"
)

var (
	maxInt8   = big.NewInt(math.MaxInt8)
	maxUint8  = big.NewInt(math.MaxUint8)
	maxInt16  = big.NewInt(math.MaxInt16)
	maxUint16 = big.NewInt(math.MaxUint16)
	maxInt32  = big.NewInt(math.MaxInt32)
	maxUint32 = big.NewInt(math.MaxUint32)
	maxInt64  = big.NewInt(math.MaxInt64)
)

func RandomInt(max int) int {
	return int(RandomInt63n(int64(max)))
}

func RandomInt63n(max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	ANoError(err)
	return n.Int64()
}

func RandomBigInt() *big.Int {
	n, err := rand.Int(rand.Reader, maxInt64)
	ANoError(err)
	return n
}

func RandomUint16() uint16 {
	n, err := rand.Int(rand.Reader, maxUint16)
	ANoError(err)
	return uint16(n.Uint64())
}

func RandomUint64() uint64 {
	return RandomBigInt().Uint64()
}

func RandomInt64() int64 {
	return RandomBigInt().Int64()
}

// newUUID generates a random UUID according to RFC 4122
func NewUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	ATrue(n == len(uuid) && err == nil)
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func DeviceID() string {
	hostname, err := os.Hostname()
	ANoError(err)
	hasher := sha1.New()
	hasher.Write([]byte(hostname))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CBCEncrypt(key, plaintext []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	r := len(plaintext) % blockSize
	if r != 0 {
		a := make([]byte, blockSize-r)
		plaintext = append(plaintext, a...)
	}
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, blockSize+len(plaintext))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	return ciphertext, nil
}

func CBCDecrypt(key, ciphertext []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < blockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	b := bytes.TrimRight(ciphertext, "\x00")
	return b, nil
}
