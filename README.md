# Forum DUA LIPA FAN CLUB

## Instalation

Clonez le repo ou téléchargez le fichier `.zip`.
Lancez un serveur WAMP et importez-y la base de données `forum.sql` disponible dans le dossier `doc/database`.
Dans un terminal de commandes lancez `go run .`.
Rendez-vous ensuite dans votre navigateur à l'adresse `localhost:8080`

## Contexte du projet

Ici un projet dans le cadre de notre B1 informatique au Campus Ynov.
Le but était de créer un forum de discussion.

Le forum devait utiliser des technologies comme le Golang, HTML, CSS, JavaScript et MySQL.
Des technologies que nous avions déjà utilisé à la seule exception du MySQL.

Notre forum à nous porte sur le thème de Dua Lipa et de son fan club. Car qui ne voudrait pas un forum sur Dua Lipa ?

## Fonctionnement du Forum

À l'arrivée sur le forum vous êtes accueilli par un message de bienvenue et d'explication.
Vous pouvez vous connecter ou aller consulter les topics et leurs messages.
Pour écrire il vous faut créer un compte ou vous connecter.

## Technologies

### Fonctionnement technique

Ce forum fonction grâce à notre serveur Go qui s'occupe de distribuer les pages HTML à un client selon la route
dans l'URL sur laquelle ce client s'y trouve. Nous utilisons les packages `net/http` et `html/tempplate` pour cela.

Le serveur Go permet aussi de faire des requêtes à notre base de données, ce qui nous permet de stocker les messages, 
les utilisateurs, les topic, etc. Elle nous permet également d'identifier les utilisateurs du site.

### Vie privée

Le forum utilise des cookies pour son bon fonctionnement, mais aucunes données personnelles ne sont stockées pour des raisons
autres que techniques.

## Auteurs

Théo COLLIOT-MARTINEZ
Enzo OUDIN
