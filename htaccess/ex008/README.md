# Génération d'un fichier Apache htaccess - Exercice 008

## Scénario

Nous allons vérifier et améliorer la qualité de notre code.

## Opérations à réaliser

1. Installation de golangci-lint

```shell
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2

golangci-lint --version
```

2. Exécutez le linter que le code actuel

```shell
golangci-lint run generate-htaccess.go

generate-htaccess.go:94:18: Error return value of `template.Execute` is not checked (errcheck)
        template.Execute(f, addr)
                        ^
generate-htaccess.go:30:10: S1028: should use fmt.Errorf(...) instead of errors.New(fmt.Sprintf(...)) (gosimple)
                return errors.New(fmt.Sprintf("invalid IP Address: %s", ip.Address))
                       ^
generate-htaccess.go:133:11: ineffectual assignment to err (ineffassign)
        options, err := parseOptions()
                 ^
```

3. Corrigez le code en suivant les recommandations

## Résultat attendu

La commande `golangci-lint run generate-htaccess.go` ne doit plus rien afficher.

## Les acquis

- Amélioration de la qualité du code avec golangci-lint