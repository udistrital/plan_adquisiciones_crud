# plan_adquisiciones_crud

El API plan_adquisiciones_crud, proporciona interfaces para la manipulación(CRUD) de los datos almacenados en una base de datos no relacional MongoDB (Registro funcionamiento selección, registro inversió actividad fuente financiamiento, registro plan adquisiones).
Esta API representa la capa de datos del subsistema de plan de adquisiciones, el cual, a su vez, hacer parte de el sistema de gestión financiero KRONOS.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones

* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [MongoDB](https://docs.mongodb.com/manual/)
* [Postgres](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/postgres.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)


### Variables de Entorno

```shell
# Ejemplo que se debe actualizar acorde al proyecto
PLAN_ADQUISICIONES_CRUD_DB_USER = [descripción]
PLAN_ADQUISICIONES_CRUD_DB_PASS = [descripción]
PLAN_ADQUISICIONES_CRUD_DB_HOST = [descripción]
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con PLAN_ADQUISICIONES_CRUD_...

### Ejecución del Proyecto

```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/plan_adquisiciones_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/plan_adquisiciones_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
SOLICITUD_CRUD_HTTP_PORT=8080 SOLICITUD_CRUD_PGURL=127.0.0.1 SOLICITUD_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile

```shell
# Implementado para despliegue del Sistema de integración continua CI.
```

### Ejecución docker-compose

```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/entradas_crud

#2. Moverse a la carpeta del repositorio
cd solicitudes_crud

#3. Crear un fichero con el nombre **custom.env**
touch .env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias

```shell
# Not Data
```

## Modelo de datos

[Mode Relacional plan_adquisicion_crud](modelobd.png)

## Estado CI

| Develop | Release | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/plan_adquisiciones_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/plan_adquisiciones_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/plan_adquisiciones_crud/status.svg?ref=refs/heads/release)](https://hubci.portaloas.udistrital.edu.co/udistrital/plan_adquisiciones_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/plan_adquisiciones_crud/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/plan_adquisiciones_crud) |

## Licencia

This file is part of plan_adquisiciones_crud

plan_adquisiciones_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

plan_adquisiciones_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with plan_adquisiciones_crud. If not, see https://www.gnu.org/licenses/.
