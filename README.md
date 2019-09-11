# Go Web Server boilerplate

## Sommaire
-  [Installation](#installation)
-  [Liste des commandes](#liste-des-commandes)
   -  [Dévelopement](#dévelopement)
   -  [Production](#production)
-  [Architecture du projet](#architecture-du-projet)
-  [Golang web server in production](#golang-web-server-in-production)
-  [Mesure et performance](#mesure-et-performance)
    -  [pprof](#pprof)
    -  [trace](#trace)
    -  [cover](#cover)
-  [TODO list](#todo-list)
-  [Astuces et explications](#Astuces-et-explications)
    -  [Architecture](#architecture)


## Installation
-  Installer la dernière de Go ([Download page](https://golang.org/dl/))
-  Copier le fichier `config.toml.dist` vers `config.toml` et renseigner les bonnes valeurs


## Liste des commandes

### Dévelopement
| Commande | Description |
|---|---|
| `make serve` | Launch Web server |
| `make log` | Launch logs rotation |
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
| `./<binaire> log` | Launch logs rotation |
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


## Mesure et performance
Go met à disposition de puissants outils pour mesurer les performances des programmes :
-  pprof (graph, flamegraph, peek)
-  trace
-  cover

=> Lien vers une vidéo intéressante [Mesure et optimisation de la performance en Go](https://www.youtube.com/watch?v=jd47gDK-yDc)

### pprof
```bash
curl http://localhost:8888/debug/pprof/heap?seconds=10 > <fichier à analyser>
go tool pprof -http :7000 <fichier à analyser> # Interface web
go tool pprof --nodefraction=0 -http :7000 <fichier à analyser> # Interface web avec tous les noeuds
go tool pprof <fichier à analyser> # Ligne de commande
```

### trace
```bash
go test <package path> -trace=<fichier à analyser>
curl localhost:<port>/debug/pprof/trace?seconds=10 > <fichier à analyser>
go tool trace <fichier à analyser>
```

### cover
```bash
go test <package path> -covermode=count -coverprofile=./<fichier à analyser>
go tool cover -html=<fichier à analyser>
```


## TODO list
-  [ ] Utiliser et rendre paramétrable le Log Level d'Echo
-  [x] Passer aux modules introduits avec Go 1.11 :
    -  https://roberto.selbach.ca/intro-to-go-modules/
    -  https://www.melvinvivas.com/go-version-1-11-modules/
    -  https://medium.com/@fonseka.live/getting-started-with-go-modules-b3dac652066d
-  [x] Mettre en place un système de migration avec GORM
-  [x] Utiliser [Viper](https://github.com/spf13/viper) pour gérer la config
-  [x] Séparer les logs d'accès des autres logs
-  SQL logs
    -  [x] Afficher la requête sans retour à la ligne
    -  [x] Gérer la variable `limit` dans fichier de configuration
    -  [x] Gérer la rotation des logs
    -  [ ] Afficher les arguments directement dans la requête ou dans un tableau
    -  [x] Logger GORM
    -  [ ] Faire une interface graphique pour afficher et filter les logs
-  [ ] Gestion des timezones
-  [ ] Facade pour les datetimes
-  [ ] Problème de consommation mémoire
    -  [https://medium.com/dm03514-tech-blog/sre-debugging-simple-memory-leaks-in-go-e0a9e6d63d4d](https://medium.com/dm03514-tech-blog/sre-debugging-simple-memory-leaks-in-go-e0a9e6d63d4d)


## Astuces et explications
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
