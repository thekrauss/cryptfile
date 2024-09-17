package fcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

// prend en entrée le chemin du fichier et un mot de passe sous forme de tableau d'octets (password).
func Encrypt(src string, password []byte) error {

	if !ValidFile(src) {
		panic("fichier non trouver")
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	textFile, err := io.ReadAll(srcFile)
	if err != nil {
		return err
	}

	// nonce de 12 octets pour l'algorithme AES-GCM
	// nonce est une valeur aléatoire utilisée pour garantir que chaque chiffrement est unique.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// PBKDF2 pour dériver une clé à partir du mot de passe et du nonce
	dk := pbkdf2.Key(password, nonce, 4096, 32, sha1.New)

	// nouveau bloc de chiffrement AES en utilisant la clé dérivée
	block, err := aes.NewCipher(dk)
	if err != nil {
		return err
	}

	// objet GCM (Galois/Counter Mode), un mode de chiffrement sécurisé pour AES
	asegm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// chiffre le texte du fichier en utilisant le nonce et AES-GCM
	cipherText := asegm.Seal(nil, nonce, textFile, nil)

	// Ajoute le nonce à la fin du texte chiffré
	cipherText = append(cipherText, nonce...)

	outFile, err := os.Create(src)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// ecrit le texte chiffré dans le fichier
	_, err = outFile.Write(cipherText)
	return err
}

func Decrypt(src string, password []byte) error {

	if !ValidFile(src) {
		panic("fichier non trouver")
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	cipherText, err := io.ReadAll(srcFile)
	if err != nil {
		return err
	}

	// recupère le nonce (les 12 derniers octets du fichier chiffré)
	nonce := cipherText[len(cipherText)-12:]

	// Sépare le texte chiffré du nonce
	cipherText = cipherText[:len(cipherText)-12]

	// Utilise PBKDF2 pour dériver la clé à partir du mot de passe et du nonce
	dk := pbkdf2.Key(password, nonce, 4096, 32, sha1.New)

	// crée un nouveau bloc de chiffrement AES en utilisant la clé dérivée
	block, err := aes.NewCipher(dk)
	if err != nil {
		return err
	}

	asegm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// dechiffre le texte chiffré en utilisant le nonce et AES-GCM
	textFile, err := asegm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return err
	}

	outFile, err := os.Create(src)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = outFile.Write(textFile)
	return err
}

func ValidFile(file string) bool {
	// Vérifie l'existence du fichier
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
