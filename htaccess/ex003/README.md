# Génération d'un fichier Apache htaccess - Exercice 003

## Scénario

La demande du client évolue, vous allez maintenant générer le fichier htaccess.txt à partir d’un fichier CSV.

Le fichier CSV contiendra également une colonne pour appliquer un Deny ou un Allow selon les cas.

## Fichier CSV

```
IP Address,Rule
192.168.1.10,Allow
172.100.2.11,Allow
1.1.1.1,Deny
```

## Opérations à réaliser

1. Remplacez le tableau d'adresse IP par le chargement du fichier CSV
2. Testez le fichier htaccess.txt sur http://www.htaccesscheck.com/

## Résultat attendu

Le fichier htaccess.txt doit contenir:

```
Order Allow,Deny
Allow from 192.168.1.10
Allow from 172.100.2.11
Deny from 1.1.1.1
Deny from all
```

## Les acquis

- Lecture d'un fichier CSV
- Utilisation d'un `struct`