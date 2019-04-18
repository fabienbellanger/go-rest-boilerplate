# go-rest-boilerplate

Golang Rest API boilerplate 

## TODO list
- Passer aux modules introduits avec Go 1.11 :
    - https://roberto.selbach.ca/intro-to-go-modules/
    - https://www.melvinvivas.com/go-version-1-11-modules/
    - https://medium.com/@fonseka.live/getting-started-with-go-modules-b3dac652066d
- Faire recover personnalisé
- SQL logs
    - afficher la requête sans retour à la ligne
    - gérer la variable `limit` dans fichier de configuration

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
