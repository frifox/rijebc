package rijebc

import (
	"bytes"
	"testing"
)

func init() {
	Key = "this is secret"
}

func TestEncrypt(t *testing.T) {
	plain := `{"foo":1}`
	expected := `MFdZbWt5WkNvM1pMOEc5TkdWdXp4TEtYT1J5cDR2TDBER1R4NHdWUGdjMD0`

	encrypted := Encrypt([]byte(plain))
	if encrypted != expected {
		t.Errorf("failed to encrypt. Encrypted=%s", encrypted)
	}
}

func TestDecrypt(t *testing.T) {
	encrypted := `MFdZbWt5WkNvM1pMOEc5TkdWdXp4TEtYT1J5cDR2TDBER1R4NHdWUGdjMD0`
	expected := `{"foo":1}`

	decrypted := Decrypt(encrypted)
	if !bytes.Equal(decrypted, []byte(expected)) {
		t.Errorf("failed to decrypt. Decrypted=%s", decrypted)
	}
}
