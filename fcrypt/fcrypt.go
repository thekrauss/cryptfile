package fcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(src string, password []byte) {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		panic(err.Error())
	}

	srcFile, err := os.Open(src)
	if err != nil {
		panic(err.Error())
	}
	defer srcFile.Close()

	textFile, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}
	key := password

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	asegm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	cipherText := asegm.Seal(nil, nonce, textFile, nil)
	cipherText = append(cipherText, nonce...)

	outFile, err := os.Create(src)
	if err != nil {
		panic(err.Error())
	}
	defer outFile.Close()

	_, err = outFile.Write(cipherText)
	if err != nil {
		panic(err.Error())
	}

}

func Decrypt(src string, password []byte) {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		panic(err.Error())
	}

	srcFile, err := os.Open(src)
	if err != nil {
		panic(err.Error())
	}
	defer srcFile.Close()

	cipherText, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}

	key := password
	salt := cipherText[len(cipherText)-12]
	str := hex.EncodeToString(salt)
}
