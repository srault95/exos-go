# Génération d'un fichier Apache htaccess - Exercice 005

## Scénario

Organisez le code avec des fonctions en découpant au mieux chaque traitement.

## Opérations à réaliser

> **Remarque**: Il y a sûrement d'autres façon d'organiser le code et l'idée de la proposition suivant est de répartir les fonctionnalités de chaque opération, de rendre réutilisable les fonctions et de prévoir la scalabilité avec les goroutines.

1. Créez la fonction loadCSV avec la signature suivante:
```golang
func loadCSV(filepath string) (addr []Ip, err error) {}
```

> La fonction loadCSV devra charger le fichier CSV et renvoyer un tableau (slice) d'Ip

2. Créez la fonction `validIp` avec la signature suivante:
```golang
func (ip *Ip) validIp() error {}
```

> La fonction validIp ne fera que l'appel à ParseIP et renverra une erreur ou nil

3. Créez la fonction `clean` avec la signature:
```golang
func clean(addr []Ip) (cleanAddr []Ip) {}
```

> La fonction clean se chargera de faire appel à `validIp` et ne renverra que les Ip valides

4. Créez la fonction `htaccess` avec la signature:
```golang
func htaccess(filepath string, addr []Ip) error {}
```

> La fonction htaccess écrira le fichier final

4. Dans la fonction `main`, appellez ces 3 fonctions dans l'ordre.

## Résultat attendu

L'erreur suivante doit s'afficher à l'écran pendant l'exécution:

```
2022/05/01 07:58:29 invalid IP Address: 2.2
```

Le fichier htaccess.txt doit contenir:

```
Order Allow,Deny
Allow from 192.168.1.10
Allow from 172.100.2.11
Deny from 1.1.1.1
Deny from all
```

## Les acquis

- Découpage du code en fonctions spécialisés
- Utilisation de variables nommées dans les retours de fonctions