package main

import (
	"bytes"
	"cryptfiles/fcrypt"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

var (
	colorReset = "\033[0m"
	colorBlue  = "\033[34m"
	colorRed   = "\033[31m"
)

func main() {
	Option()
}

func Usage() {
	fmt.Println("Usage :")
	fmt.Println(colorBlue, "Exemple pour lancer le mode encrypt : go run . encrypt /chemin/du/fichier/", colorReset)
	fmt.Println(colorBlue, "Exemple pour lancer le mode decrypt : go run . decrypt /chemin/du/fichier/", colorReset)
}

func Option() {
	if len(os.Args) < 3 {
		fmt.Println(colorRed, "Erreur : Vous devez spécifier le mode (encrypt/decrypt) et le chemin du fichier.", colorReset)
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
		fmt.Println(colorRed, "Erreur : Mode inconnu. Utilisez 'encrypt' ou 'decrypt'.", colorReset)
		Usage()
		os.Exit(1)
	}
}

func DecryptFunc() {
	if len(os.Args) < 3 {
		fmt.Println(colorRed, "Erreur : Vous devez spécifier le chemin du fichier.", colorReset)
		Usage()
		os.Exit(1)
	}
	file := os.Args[2]

	if !fcrypt.ValidFile(file) {
		fmt.Println(colorRed, "Erreur : Fichier non trouvé.", colorReset)
		os.Exit(1)
	}

	fmt.Printf("Entrez le mot de passe : ")
	password, _ := term.ReadPassword(0)

	if len(password) == 0 {
		fmt.Println(colorRed, "\nErreur : le mot de passe ne peut pas être vide.", colorReset)
		os.Exit(1)
	}

	fmt.Println("\nDéchiffrement en cours...")
	err := fcrypt.Decrypt(file, password)
	if err != nil {
		fmt.Println("Erreur lors du déchiffrement :", err)
		os.Exit(1)
	}

	fmt.Println("\nFichier déchiffré avec succès.")
}

func EncryptFunc() {
	if len(os.Args) < 3 {
		fmt.Println(colorRed, "Erreur : Vous devez spécifier le chemin du fichier.", colorReset)
		Usage()
		os.Exit(1)
	}
	file := os.Args[2]

	if !fcrypt.ValidFile(file) {
		fmt.Println(colorRed, "Erreur : Fichier non trouvé.", colorReset)
		os.Exit(1)
	}

	password := GetPassword()

	fmt.Println("\nChiffrement en cours...")
	err := fcrypt.Encrypt(file, password)
	if err != nil {
		fmt.Println("Erreur lors du chiffrement :", err)
		os.Exit(1)
	}

	fmt.Println("\nFichier chiffré avec succès.")
}

func GetPassword() []byte {
	fmt.Printf("Entrez le mot de passe : ")
	password, _ := term.ReadPassword(0)

	if len(password) == 0 {
		fmt.Printf(colorRed, "\nErreur : le mot de passe ne peut pas être vide.\n", colorReset)
		return GetPassword()
	}

	fmt.Printf("\nConfirmez le mot de passe : ")
	password2, _ := term.ReadPassword(0)

	if !ValidPassword(password, password2) {
		fmt.Printf(colorRed, "\nLes mots de passe ne correspondent pas. Réessayez !\n", colorReset)
		return GetPassword()
	}

	return password
}

func ValidPassword(password []byte, password2 []byte) bool {
	return bytes.Equal(password, password2)
}
