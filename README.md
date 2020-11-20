# PLAN ADQUISICIONES CRUD

El API plan_adquisiciones_crud, proporciona interfaces para la manipulación(CRUD) de los datos almacenados en una base de datos no relacional MongoDB (Registro funcionamiento selección, registro inversió actividad fuente financiamiento, registro plan adquisiones).
Esta API representa la capa de datos del subsistema de plan de adquisiciones, el cual, a su vez, hacer parte de el sistema de gestión financiero KRONOS.



## Modelo de datos
![mbd](modelobd.png)


## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [MongoDB]()
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


## Instalación y Ejecución

Para instalar y correr el proyecto de debe relizar lo siguientes pasos:

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get https://github.com/udistrital/plan_adquisiciones_crud.git

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/plan_adquisiciones_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
PLAN_ADQUISICIONES_HTTP_PORT=8080 PLAN_ADQUISICIONES_DB_HOST=127.0.0.1:27017 PLAN_ADQUISICIONES_SOME_VARIABLE=some_value bee run
```

### Ejecución docker-compose

**para usar esta opcion es necesario contar con [DOCKER](https://docs.docker.com/) y [DOCKER-COMPOSE](https://docs.docker.com/compose/) en cualquier SO compatible**

```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/plan_adquisiciones_crud.git

#2. Moverse a la carpeta del repositorio
cd plan_adquisiciones_crud

#3. Crear la network **back_end** para los contenedores
docker network create back_end

#4. Ejecutar el compose del contenedor
docker-compose up --build

#5. Comprobar que los contenedores estén en ejecución
docker ps
```
- Al finalizar se podran consumir los servicios del API en los puertos definidos en **.env** y **custom.env** actualmente 8080

### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```

## Estado CI
| Develop | Relese  | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/plan_adquisiciones_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/cuentas_contables_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/plan_adquisiciones_crud/status.svg?ref=refs/heads/release)](https://hubci.portaloas.udistrital.edu.co/udistrital/cuentas_contables_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/plan_adquisiciones_crud/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/cuentas_contables_crud) |



## Licencia
This file is part of plan_adquisicion_crud

plan_adquisicion_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

cuentas_contables_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with necesidades_crud. If not, see https://www.gnu.org/licenses/.



