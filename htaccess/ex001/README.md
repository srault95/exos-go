# Génération d'un fichier Apache htaccess - Exercice 001

## Scénario

Dans ce scénario, la liste des adresses IP sera statique et écrite en dur dans le code.

Pour le moment, nous allons uniquement afficher le résultat à l'écran, nous verrons dans un prochain exercices, comment l'écrire dans un fichier.

## Adresses fournies par le client

```
192.168.1.10
172.100.2.11
```

## Opérations à réaliser

1. Créez un fichier generate-htaccess.go
2. Placez les adresses IP dans un tableau dynamique (slice)
3. Bouclez sur le tableau pour afficher chaque adresse IP à l'écran avec le "Allow from"

## Résultat attendu

Affichage à l'écran:

```
Order Allow,Deny
Allow from 192.168.1.10
Allow from 172.100.2.11
Deny from all
```

## Les acquis

- Utilisation d'un slice (tableau dynamique)
- Boucler sur un slice avec l'instruction range
