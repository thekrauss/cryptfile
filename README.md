
# CryptFiles - Chiffrement et Déchiffrement de Fichiers en Go

Ce projet permet de chiffrer et de déchiffrer des fichiers à l'aide d'un mot de passe. Il utilise l'algorithme de chiffrement AES-GCM pour garantir la sécurité des fichiers.

## Fonctionnalités

- **Chiffrement** : Chiffrez n'importe quel fichier en utilisant un mot de passe sécurisé.
- **Déchiffrement** : Déchiffrez un fichier précédemment chiffré avec le mot de passe correspondant.
- **AES-GCM** : Utilisation d'AES en mode Galois/Counter Mode (GCM) pour garantir la confidentialité et l'intégrité des fichiers.
- **Interaction sécurisée** : Les mots de passe ne sont pas affichés à l'écran lors de leur saisie.

## Installation

1. Clonez le dépôt :
   ```bash
   git clone https://github.com/thekrauss/cryptfile.git
   cd cryptfiles
   ```

2. Installez les dépendances si nécessaire :
   ```bash
   go get golang.org/x/crypto/pbkdf2
   go get golang.org/x/term
   ```

3. Compilez le projet :
   ```bash
   go build -o cryptfiles
   ```

## Utilisation

Le programme prend deux modes principaux : `encrypt` pour chiffrer un fichier et `decrypt` pour le déchiffrer.

### Exemple de commande

#### Chiffrement

Pour chiffrer un fichier, utilisez la commande suivante :

```bash
go run . encrypt /chemin/du/fichier
```

Vous serez invité à entrer un mot de passe et à le confirmer. Le fichier sera ensuite chiffré en place.

#### Déchiffrement

Pour déchiffrer un fichier, utilisez la commande suivante :

```bash
go run . decrypt /chemin/du/fichier
```

Vous devrez entrer le mot de passe utilisé lors du chiffrement pour que le fichier soit déchiffré avec succès.

### Exemple de sortie

```bash
$ go run . encrypt /chemin/du/fichier
Entrez le mot de passe: ********
Confirmez le mot de passe: ********

Chiffrement en cours...
Fichier chiffré avec succès.
```

## Structure du projet

Le projet est organisé en plusieurs fichiers :

- `main.go` : Contient la logique principale de gestion des options de chiffrement/déchiffrement.
- `fcrypt/fcrypt.go` : Contient les fonctions de chiffrement et déchiffrement utilisant AES-GCM et PBKDF2.

## Fonctionnement

1. **Chiffrement** : Le programme prend le fichier à chiffrer et génère une clé dérivée du mot de passe. Il chiffre ensuite le fichier en utilisant AES-GCM et y ajoute un nonce (valeur aléatoire unique utilisée pour ce chiffrement).

2. **Déchiffrement** : Lors du déchiffrement, le programme utilise le nonce stocké dans le fichier pour dériver la clé à partir du mot de passe, puis déchiffre le fichier.
