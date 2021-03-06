# Go Web Server boilerplate

[![Go Report Card](https://goreportcard.com/badge/github.com/fabienbellanger/go-rest-boilerplate)](https://goreportcard.com/report/github.com/fabienbellanger/go-rest-boilerplate)
[![Build Status](https://travis-ci.org/fabienbellanger/go-rest-boilerplate.svg?branch=master)](https://travis-ci.org/fabienbellanger/go-rest-boilerplate)
[![GoDoc](https://godoc.org/github.com/fabienbellanger/go-rest-boilerplate?status.svg)](https://godoc.org/github.com/fabienbellanger/go-rest-boilerplate)

## Sommaire
-  [Installation](#installation)
-  [Liste des commandes](#liste-des-commandes)
   -  [Dévelopement](#dévelopement)
   -  [Production](#production)
-  [Architecture du projet](#architecture-du-projet)
-  [Golang web server in production](#golang-web-server-in-production)
-  [Deploiement avec Docker](#Deploiement-avec-Docker)
    -  [Liens](#Liens)
    -  [Commandes](#Commandes)
    -  [TODO](#TODO)
-  [Serveur Web](#Serveur-Web)
-  [Mesure et performance](#mesure-et-performance)
    -  [pprof](#pprof)
    -  [trace](#trace)
    -  [cover](#cover)
-  [TODO list](#todo-list)
-  [Astuces et explications](#Astuces-et-explications)
    -  [Performance, Débug et profilage](#Performance,-Débug-et-profilage)
    -  [Architecture](#architecture)


## Installation
-  Installer la dernière de Go ([Download page](https://golang.org/dl/))
-  Copier le fichier `config.toml.dist` vers `config.toml` et renseigner les bonnes valeurs


## Liste des commandes

### Dévelopement
| Commande | Description |
|---|---|
| `make serve` | Launch Web server |
| `make logsRotation` | Launch logs rotation |
| `make logsExport` | Launch logs CSV export |
| `make dbInit` | Launch database initilization |
| `make dbDump` | Launch database dump |
| `make make-migration` | Create migration |
| `make migrate` | Make migrations |
| `make ws` | Launch WebSockets server |

### Production

Compiler le fichier binaire `<binaire>` avec `make build` et renseigner des bonnes valeurs de le fichier de 
configuration `config.toml`.

| Commande | Description |
|---|---|
| `./<binaire> serve` | Launch Web server |
| `./<binaire> logs-rotation` | Launch logs rotation |
| `./<binaire> logs-export` | Launch logs CSV export |
| `./<binaire> db --init` | Launch database initilization |
| `./<binaire> db --dump` | Launch database dump |
| `./<binaire> make-migration -n <Name in CamelCase>` | Create migration |
| `./<binaire> migrate` | Make migrations |
| `./<binaire> websocket` | Launch WebSockets server |


## Architecture du projet
```
\_ assets
   \_ js
\_ commands
\_ database
\_ handlers
   \_ api
   \_ web
\_ lib
\_ logs
\_ migrations
\_ models
\_ repositories
   \_ user
\_ routes
   \_ user
   \_ echo
   \_ web
\_ templates
   \_ example
   \_ layout
\_ websockets
```

-  Le dossier `assets` contient les fichiers multimédia (images, vidéos, etc.), JavaScript ou encore CSS.
-  Le dossier `commands` contient toutes les commandes que l'on peut lancer depuis un terminal.
-  Le dossier `database` contient tous les fichiers relatifs à l'utilisation de MySQL ainsi que l'initialisation 
    et le dump de la base. Il contient également l'initalisation de l'ORM.
-  Le dossier `handlers` contient tous les handlers du serveur Web. Ils sont divisés par type. 
    Par exemple, on a un dossier `api` pour gérer les API et un dossier `web` pour gérer un "site".
-  Le dossier `lib` contient des fonctions globales à l'application.
-  Le dossier `logs` contient les logs du serveur Web.
-  Le dossier `migrations` contient les fichiers de migrations.
-  Le dossier `models` contient les modèles (base de données).
-  Le dossier `repositories` contient les repositories.
    Ces fichiers permettent d'écrire les requêtes s'appliquant à un modèle.
    Les fichiers dépendant du type de base de données, ils sont suffixés par le type de base de données, 
    par exemple, `_mysql`.
-  Le dossier `routes` contient les fichiers relatifs au routing. Ils sont divisés par type. 
    Par exemple, on a un dossier `api` pour gérer les API et un dossier `web` pour gérer un "site".
-  Le dossier `templates` contient les templates des différentes page Web.
-  Le dossier `websockets` contient les fichiers relatifs au serveur de WebSockets.


## Golang web server in production
-  [Systemd](https://jonathanmh.com/deploying-go-apps-systemd-10-minutes-without-docker/)
-  [ProxyPass](https://evanbyrne.com/blog/go-production-server-ubuntu-nginx)
-  [How to Deploy App Using Docker](https://medium.com/@habibridho/docker-as-deployment-tools-5a6de294a5ff)

### Creating a Service for Systemd
```bash
touch /lib/systemd/system/<service name>.service
```

Edit file:
```
[Unit]
Description=<service description>
After=network.target

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=<path to exec with arguments>

[Install]
WantedBy=multi-user.target
```

| Commande | Description |
|---|---|
| `service <service name> start` | To launch |
| `service <service name> enable` | To enable on boot |
| `service <service name> disable` | To disable on boot |
| `service <service name> status` | To show status |
| `service <service name> stop` | To stop |


## Deploiement avec Docker

### Liens
-  [Deploying Go servers with Docker](https://blog.golang.org/docker)
-  [Building Docker Containers for Go Applications](https://www.callicoder.com/docker-golang-image-container-example/)

### Commandes
| Commande | Description |
|---|---|
| `docker image ls` | Lister les images |
| `docker image remove <image_id>` ou `docker rmi <image_id>` | Supprimer une image |
| `docker container ls` | Lister les conteneurs |
| `docker container stop <container_id>` | Stopper un conteneur |
| `docker rm <container_id>` | Supprimer un conteneur |
| `docker build -t <image_name> -f <dockerfile_name> .` | Build de l'image |
| `docker run -d -p <port_local>:<port_container> <image_name>` | Lancement de l'image |

### TODO
-  [ ] Pouvoir modifier dynamique le port et le nom du binaire dans le dockerfile

## Serveur Web
Le serveur peut contenir plusieurs sous-domaines. Leur configuration se fait dans le fichier `config.toml` via la partie
`[server]`. Le serveur possède 3 sous-domaines par défaut :
-  `apiSubDomain` : sous-domaine relatif aux API
-  `clientSubDomain` :  sous-domaine pour lancer, par exemple, une application JavaScript Vue.js
-  `webSubDomain` : sous-domaine pour créer une application côté serveur. Il contient également les routes de debug pour 
`pprof` ou `trace`.


## Mesure et performance
Go met à disposition de puissants outils pour mesurer les performances des programmes :
-  pprof (graph, flamegraph, peek)
-  trace
-  cover

=> Lien vers une vidéo intéressante [Mesure et optimisation de la performance en Go](https://www.youtube.com/watch?v=jd47gDK-yDc)

### pprof
Lancer :
```bash
curl http://localhost:8888/debug/pprof/heap?seconds=10 > <fichier à analyser>
```
Puis :
```bash
go tool pprof -http :7000 <fichier à analyser> # Interface web
go tool pprof --nodefraction=0 -http :7000 <fichier à analyser> # Interface web avec tous les noeuds
go tool pprof <fichier à analyser> # Ligne de commande
```

### trace
Lancer :
```bash
go test <package path> -trace=<fichier à analyser>
curl localhost:<port>/debug/pprof/trace?seconds=10 > <fichier à analyser>
```
Puis :
```bash
go tool trace <fichier à analyser>
```

### cover
Lancer :
```bash
go test <package path> -covermode=count -coverprofile=./<fichier à analyser>
```
Puis :
```bash
go tool cover -html=<fichier à analyser>
```


## TODO list
-  [x] Passer aux modules introduits avec Go 1.11 :
    -  https://roberto.selbach.ca/intro-to-go-modules/
    -  https://www.melvinvivas.com/go-version-1-11-modules/
    -  https://medium.com/@fonseka.live/getting-started-with-go-modules-b3dac652066d
-  [x] Mettre en place un système de migration avec GORM
-  [x] Utiliser [Viper](https://github.com/spf13/viper) pour gérer la config
-  [x] Séparer les logs d'accès des autres logs
-  [x] Ajouter une Basic Auth pour pprof
-  [x] SQL logs
    -  [x] Afficher la requête sans retour à la ligne
    -  [x] Gérer la variable `limit` dans fichier de configuration
    -  [x] Gérer la rotation des logs
    -  [] Logger GORM dans un fichier de log [GORM logger](http://gorm.io/docs/logger.html)
-  [x] Mettre les logs SQL dans un fichier à part
-  [x] Exporter les logs en CSV
-  [ ] Faire une interface graphique pour afficher et filter les logs
    -  [x] Mettre une Basic Auth
-  [ ] Gestion des timezones
-  [ ] Interface pour les datetimes


## Astuces et explications

### Performance, Débug et profilage
-  [https://medium.com/dm03514-tech-blog/sre-debugging-simple-memory-leaks-in-go-e0a9e6d63d4d](https://medium.com/dm03514-tech-blog/sre-debugging-simple-memory-leaks-in-go-e0a9e6d63d4d)
-  [pprof & Memory leaks](https://www.freecodecamp.org/news/how-i-investigated-memory-leaks-in-go-using-pprof-on-a-large-codebase-4bec4325e192/)
-  [Allocation efficiency in high-performance Go services](https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/)
-  [Astuces Slices](https://github.com/golang/go/wiki/SliceTricks)
-  [Pool](https://www.akshaydeo.com/blog/2017/12/23/How-did-I-improve-latency-by-700-percent-using-syncPool/)
-  [Understanding The Memory Model Of Golang : Part 1](https://medium.com/@edwardpie/understanding-the-memory-model-of-golang-part-1-9814f95621b4)
-  [Understanding The Memory Model Of Golang : Part 2](https://medium.com/@edwardpie/understanding-the-memory-model-of-golang-part-2-972fe74372ba)
-  [High Performance Go Workshop](https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html)

### Architecture
-  [Trying Clean Architecture on Golang](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)
-  [go-clean-arch](https://github.com/bxcodec/go-clean-arch/tree/master/models)
-  [Beautify your Golang project](https://itnext.io/beautify-your-golang-project-f795b4b453aa)
