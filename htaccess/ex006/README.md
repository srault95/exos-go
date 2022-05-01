# Génération d'un fichier Apache htaccess - Exercice 006

## Scénario

Dans cet exercice, vous allez ajouter le moyen de fournir le chemin du fichier CSV et celui du fichier htaccess, à l'aide d'options dans la ligne de commande.

En plus des options en ligne de commandes, nous allons ajouter la possibilité d'utiliser des variables d'environnement pour affecter des valeurs par défaut.

> Pour le moment, nous allons utiliser la librairie standard de Golang mais plus tard, nous pourrons remplacer le module `flag`, par [Cobra](https://github.com/spf13/cobra) et `os.LookupEnv()` par une gestion de configuration avancé avec [Viper](https://github.com/spf13/viper)

## Opérations à réaliser

1. A l'aide du module [flag](https://pkg.go.dev/flag), ajoutez des options à la ligne de commande pour le chemin du fichier CSV et celui du fichier htaccess.
2. Utilisez un struct pour charger les options
3. La commande `go run generate-htaccess.go --help` doit renvoyer les options disponibles
4. Les valeurs par défaut des options doivent pouvoir être surchargé par des variables d'environnement

## Résultat attendu

```shell
go run generate-htaccess.go --help
Usage of /tmp/go-build1054416755/b001/exe/generate-htaccess:
  -D    Debug mode
  -c string
        CSV filepath (default "ip-addresses.csv")
  -h string
        Htaccess filepath (default "htaccess.txt")

# En utilisant les variables d'environnement
CSV_PATH=test.csv HTACCESS_PATH=.htaccess go run generate-htaccess.go --help
Usage of /tmp/go-build94048889/b001/exe/generate-htaccess:
  -c string
        CSV filepath (default "test.csv")
  -h string
        Htaccess filepath (default ".htaccess")        
```

## Les acquis

- Ajout d'options à l'exécution avec le module `flag`
- Utilisation de variables d'environnement
