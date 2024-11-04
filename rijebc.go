package rijebc

import (
	"bytes"
	"encoding/base64"
	"strings"
)

var Key string

func Encrypt(input []byte) string {
	encrypted := rijEncrypt(input)
	encoded := base64encode(encrypted)

	return encoded
}

func base64encode(bytes []byte) string {
	var encoded string

	encoded = base64.StdEncoding.EncodeToString(bytes)
	encoded = base64.StdEncoding.EncodeToString([]byte(encoded))
	encoded = strings.TrimSuffix(encoded, "=")

	return encoded
}
func rijEncrypt(input []byte) []byte {
	for len(input)%BlockSize > 0 {
		input = append(input, []byte{0}...)
	}

	var key [BlockSize]byte
	copy(key[:], Key) // last 2 bytes will be 0x00
	cipher := newCipher(&key)

	blocks := len(input) / BlockSize
	block := 0
	for block < blocks {
		start := block * BlockSize
		end := block*BlockSize + BlockSize

		var chunk [BlockSize]byte
		copy(chunk[:], input[start:end])

		cipher.Encrypt(&chunk, &chunk)

		copy(input[start:end], chunk[:])

		block++
	}

	return input
}
func Decrypt(input string) []byte {
	decoded := bas64decode(input)
	decrypted := rijDecrypt(decoded)

	return decrypted
}

func bas64decode(input string) []byte {
	var decoded []byte
	var err error

	if !strings.HasSuffix(input, "=") {
		input += "="
	}

	decoded, err = base64.StdEncoding.DecodeString(input)
	if err != nil {
		return []byte{}
	}

	decoded, err = base64.StdEncoding.DecodeString(string(decoded))
	if err != nil {
		return []byte{}
	}

	return decoded
}

func rijDecrypt(input []byte) []byte {
	var key [BlockSize]byte
	copy(key[:], []byte(Key)) // last 2 bytes will be 0x00
	cipher := newCipher(&key)

	blocks := len(input) / BlockSize
	block := 0
	for block < blocks {
		start := block * BlockSize
		end := block*BlockSize + BlockSize

		var chunk [BlockSize]byte
		copy(chunk[:], input[start:end])

		cipher.Decrypt(&chunk, &chunk)

		copy(input[start:end], chunk[:])

		block++
	}

	input = bytes.Trim(input, "\x00")

	return input
}
