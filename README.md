Voici un exemple de fichier README pour votre projet **ASCII Art Color** :

# ASCII Art Color

## Description

**ASCII Art Color** est un programme en Go qui permet de convertir des chaînes de caractères en art ASCII et de colorer des lettres spécifiques dans le terminal. Vous pouvez spécifier les lettres à colorer et la couleur souhaitée via des options de ligne de commande.

## Installation

1. Clonez le dépôt ou téléchargez les fichiers source.
2. Assurez-vous d'avoir Go installé sur votre machine.
3. Compilez ou exécutez directement le programme avec `go run`.

## Utilisation

### Commande de base

```sh
go run . --color=<color> <letters to be colored> "<string>"
```

### Options

- `--color=<color>` : Spécifiez la couleur des lettres. Les couleurs disponibles sont : `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`, `default`.

### Exemple

Pour colorer les lettres `a`, `b`, et `c` en rouge dans le texte "hello abc" :

```sh
go run . --color=red abc "hello abc"
```

Cela affichera le texte "hello abc" avec les lettres `a`, `b`, et `c` en rouge.

## Fichier de police

Par défaut, le programme utilise un fichier de police nommé `standard.txt`. Vous pouvez spécifier un autre fichier de police en ajoutant un argument supplémentaire :

```sh
go run . --color=red abc "hello abc" customFont
```

Assurez-vous que le fichier de police se trouve dans le même répertoire que le programme et qu'il est formaté correctement.

## Exemple de fichier de police

Un exemple de fichier de police (`standard.txt`)

## Auteur

Développé par Baran.
