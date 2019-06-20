# Golang Rest API boilerplate

## Content
- [Installation](#installation)
- [Commands list](#commands-list)
   - [Development](#development)
   - [Production](#production)
- [Golang web server in production](#golang-web-server-in-production)
- [TODO list](#todo-list)


## Installation
- Install the lastest Golang version ([Download page](https://golang.org/dl/))
- TODO: Install dependencies
- Copy file `config.toml.dist` to `config.toml` and fill it.


## Commands list

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
