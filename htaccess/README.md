# Génération d'un fichier Apache htaccess avec Golang

Un client vous demande de générer un fichier [htaccess](http://httpd.apache.org/docs/current/howto/htaccess.html) pour Apache à partir d'une liste d'adresse IP qu'il vous fournit périodiquement.

Nous allons progressivement, développer le script de traitement, en partant d'une liste d'adresse IP statiques dans le code, puis nous la remplacerons par un fichier CSV. 

D'autres évolutions sont possibles, comme partir d'un fichier XLSX ou développer une IHM pour que le client gère directement ses adresses IP.

## Exercice 001

Affichage à l'écran, du contenu qui devra être copié dans un fichier .htaccess

## Exercice 002

Création et écriture du résultat dans un fichier htaccess.txt

## Exercice 003

Chargement des adresses IP à partir d'un fichier CSV.

## Exercice 004

Validation du format des IP fournis par le CSV.

## Exercice 005

Découpage du code en fonctions spécialisés.

## Exercice 006

Utilisation d'options en ligne de commande et de variables d'environnement.

## Exercice 007

Utilisation d'un template pour générer le fichier .htaccess

Voir aussi:
- https://github.com/flosch/pongo2
- https://github.com/Masterminds/sprig

## Exercice 008

Tests qualités et bonnes pratiques Golang.

- https://golangci-lint.run/

## Exercice 009 (TODO)

Développement des tests unitaires.

- https://github.com/stretchr/testify

## Exercice 010 (TODO)

Découpage du code en modules.

## Idées d'exercices à définir et prioriser

- Appel Rest pour vérifier que sa propre IP publique soit toujours autorisé
- Gestion des IP V6 et des CIDR
- Génération automatique de la documentation.
- Automatisation CI/CD et production d'un binaire versionné.
- Utilisation d'une librairie pour le logging (logrus, zerolog, ...)
    - https://github.com/hashicorp/terraform/blob/main/internal/logging/logging.go
    - https://github.com/goharbor/harbor/blob/main/src/lib/log/logger.go
- Remplacement du module `flag`, par [Cobra](https://github.com/spf13/cobra)
    - https://github.com/ory/hydra/blob/master/cmd/clients.go
- Gestion de configuration avancée avec [Viper](https://github.com/spf13/viper)
- Création d'une API Rest (Gin ?)
- Création d'une IHM (choix du framework: https://component-party.pages.dev/)
- Version Nginx



