GoChannel
Objetivo: Se busca implementar un chat con salas sin incluir librerias especificas de chat

Servidor GO con framework ECHO 
Web estatica Vue.js y material (Proximamente VUEX)

### Funcionamiento
#### Backend
- Descargar dependencias

```
go test
```

- Basico

```
go run main.go
```

- Para debug F5 desde vscode

#### WEB
La pagina web se puede levantar con cualquier server HTTP.

Yo uso la extension de VSCODE

```
Name: Live Server
Id: ritwickdey.liveserver
Description: Launch a development local Server with live reload feature for static & dynamic pages
Version: 5.6.1
Publisher: Ritwick Dey
VS Marketplace Link: https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer
```

### Proximamente Docket
#### Desarrollo
- Levantar contenedores
```sh
 docker-compose -f docker-compose-dev.yml up --build --force-recreate
```

-- Remover contenedores 
```
docker-compose -f docker-compose-dev.yml rm --force
```

### Conexion a server de redis
```sh
 docker run -it --network gochat_default --rm redis redis-cli -h redis
```

### Postgres

- Ingreso a pgadmin
http://localhost:5050/browser/#

user: marcos.d.perez@gmail.com
password: password