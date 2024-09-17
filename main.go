package main

import (
	"bytes"
	"cryptfiles/fcrypt"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	Option()
}

func Usage() {
	colorReset := "\033[0m"
	colorBlue := "\033[34m"

	fmt.Println("Usage :")
	fmt.Println(colorBlue, "Exemple pour lancer le mode encrypt : go run . encrypt /chemin/du/fichier/", colorReset)
	fmt.Println(colorBlue, "Exemple pour lancer le mode decrypt : go run . decrypt /chemin/du/fichier/", colorReset)
}

func Option() {
	if len(os.Args) < 3 {
		Usage()
		os.Exit(1)
	}

	mode := strings.ToLower(os.Args[1])

	switch mode {
	case "encrypt":
		EncryptFunc()
	case "decrypt":
		DecryptFunc()
	default:
		Usage()
		os.Exit(1)
	}
}

func DecryptFunc() {
	if len(os.Args) < 3 {
		fmt.Println("Erreur : chemin de fichier manquant.")
		Usage()
		os.Exit(1)
	}
	file := os.Args[2]

	if !ValidFile(file) {
		panic("Fichier non trouvé")
	}

	fmt.Printf("Entrez le mot de passe: ")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDéchiffrement en cours...")

	fcrypt.(file, password)
	fmt.Println("\nFichier déchiffré avec succès")
}

func EncryptFunc() {
	if len(os.Args) < 3 {
		fmt.Println("Erreur : chemin de fichier manquant.")
		Usage()
		os.Exit(1)
	}
	file := os.Args[2]

	if !ValidFile(file) {
		panic("Fichier non trouvé")
	}

	password := GetPassword()
	fmt.Println("\nChiffrement en cours...")

	fcrypt.Encrypt(file, password)
	fmt.Println("\nFichier chiffré avec succès")
}

func GetPassword() []byte {
	fmt.Printf("Entrez le mot de passe: ")
	password, _ := term.ReadPassword(0)

	fmt.Printf("\nConfirmez le mot de passe: ")
	password2, _ := term.ReadPassword(0)

	if !ValidPassword(password, password2) {
		fmt.Printf("\nLes mots de passe ne correspondent pas. Réessayez !\n")
		return GetPassword()
	}
	return password
}

func ValidPassword(password []byte, password2 []byte) bool {
	return bytes.Equal(password, password2)
}

func ValidFile(file string) bool {
	// Vérifie l'existence du fichier
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
