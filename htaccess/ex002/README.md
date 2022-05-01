# Génération d'un fichier Apache htaccess - Exercice 002

## Scénario

Précédement, nous avons appris à générer le htaccess à l'écran. Nous allons maintenant l'écrire 
dans un fichier.

> Le fichier à générer se nomme habituellement `.htaccess` mais pour simplifier son utilisation, nous allons le nommer htaccess.txt

## Opérations à réaliser

1. Remplacez l'affichage à l'écran par l'écriture dans un fichier htaccess.txt
2. Testez le fichier htaccess.txt sur http://www.htaccesscheck.com/

## Résultat attendu

Contenu du fichier:

```
Order Allow,Deny
Allow from 192.168.1.10
Allow from 172.100.2.11
Deny from all
```

## Les acquis

- Ecrire dans un fichier texte
- Utilisation de l'instruction `defer`
- Traitement des erreurs
