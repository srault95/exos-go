# Génération d'un fichier Apache htaccess - Exercice 004

## Scénario

À la suite d’erreurs dans le format d’adresse IP, vous allez ajouter une validation du format dans le générateur de fichier .htaccess à l’aide du module [net](https://pkg.go.dev/net) 

Vous en profiterez pour metter en oeuvre le Logging.

## Opérations à réaliser

1. Ajoutez une ligne dans le fichier CSV avec une adresse IP mal formatée. ex: `2.2`
2. Pour chaque adresse ip trouvée dans le fichier CSV, validez le format avec la fonction [net.ParseIP](https://pkg.go.dev/net#ParseIP)
3. Gérez et affichez l'erreur ainsi que l'adresse IP concernée

## Fichier CSV

```
IP Address,Rule
192.168.1.10,Allow
172.100.2.11,Allow
1.1.1.1,Deny
2.2,Deny
```

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

- Utilisation des instructions `if` et `continue`
- Utilisation de la fonction net.ParseIP
- Ajout d'une méthode (pointer receiver) à un struct
- Génération d'une erreur avec `errors.New()`
- Gestion des erreurs
- Utilisation du module log
