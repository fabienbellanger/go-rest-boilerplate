# Golang Rest API boilerplate

## Content
- [Installation](#installation)
- [Liste des commandes](#liste-des-commandes)
   - [Development](#development)
   - [Production](#production)
- [Architecture du projet](#architecture-du-projet)
- [Golang web server in production](#golang-web-server-in-production)
- [TODO list](#todo-list)


## Installation
- Install the lastest Golang version ([Download page](https://golang.org/dl/))
- TODO: Install dependencies
- Copy file `config.toml.dist` to `config.toml` and fill it.


## Liste des commandes

### Development

#### Launch Web server
```
make api
```

#### Launch logs rotation
```
make log
```

#### Launch database initilization
```
make dbInit
```

#### Launch database dump
```
make dbDump
```

#### Create migration
```
make make-migration
```

#### Make migrations
```
make migrate
```

#### Launch WebSockets server
```
make ws
```

### Production

Compile binary `<binaire>` with `make build` and fill configuration file `config.toml`.

#### Launch Web server
```
<binaire> api
```

#### Launch logs rotation
```
<binaire> log
```

#### Launch database initilization
```
<binaire> db --init
```

#### Launch database dump
```
<binaire> db --dump
```

#### Create migration
```
<binaire> make-migration -n <Name in CamelCase>
```

#### Make migrations
```
<binaire> migrate
```

#### Make migrations
```
<binaire> websocket
```


## Architecture du projet
```
\_ assets
   \_ js
\_ commands
\_ controllers
\_ database
\_ lib
\_ logs
\_ orm
   \_ migrations
   \_ models
\_ routes
\_ templates
   \_ example
   \_ layout
\_ websockets
```

- Le dossier `assets` contient les fichiers multimédia (images, vidéos, etc.), JavaScript ou encore CSS.
- Le dossier `commands` contient toutes les commandes que l'on peut lancer depuis un terminal.
- Le dossier `controllers` contient toutes les controleurs du serveur Web.
- Le dossier `database` contient tous les fichiers relatifs à l'utilisation de MySQL ainsi que l'initialisation et le dump de la base.
- Le dossier `lib` contient des fonctions globales à l'application.
- Le dossier `logs` contient les logs du serveur Web.
- Le dossier `orm` contient les fichiers de migrations ainsi que les modèles.
- Le dossier `routes` contient les fichiers relatifs au routing.
- Le dossier `templates` contient les templates des différentes page Web.
- Le dossier `websockets` contient les fichiers relatifs au serveur de WebSockets.


## Golang web server in production
- [Systemd](https://jonathanmh.com/deploying-go-apps-systemd-10-minutes-without-docker/)
- [ProxyPass](https://evanbyrne.com/blog/go-production-server-ubuntu-nginx)

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

To launch:
```bash
service <service name> start
```

To enable on boot:
```bash
service <service name> enable
```

To disable on boot:
```bash
service <service name> disable
```

To show status:
```bash
service <service name> status
```

To stop:
```bash
service <service name> stop
```


## TODO list
- [x] Passer aux modules introduits avec Go 1.11 :
    - https://roberto.selbach.ca/intro-to-go-modules/
    - https://www.melvinvivas.com/go-version-1-11-modules/
    - https://medium.com/@fonseka.live/getting-started-with-go-modules-b3dac652066d
- [ ] Mettre en place un système de migration avec GORM
- [ ] Utiliser [Viper](https://github.com/spf13/viper) pour gérer la config
- SQL logs
    - [x] Afficher la requête sans retour à la ligne
    - [x] Gérer la variable `limit` dans fichier de configuration
    - [x] Gérer la rotation des logs
    - [ ] Afficher les arguments directement dans la requête ou dans un tableau
    - [ ] Logger GORM
    - [ ] Faire une interface graphique pour afficher et filter les logs
- [ ] Gestion des timezones
- [ ] Facade pour les datetimes
