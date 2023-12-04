# microbloggin-pltf

## Instrucciones para levantar el proyecto

1. Clona el repositorio en tu máquina local.
2. Dentro del directorio raíz del proyecto `microblogging-pltf` navega hasta la carpeta `cmd`  y ejecuta el archivo main.go utilizando el comando ```go run main.go``` .

## Documentación

En la carpeta docs se encuentra un png que ilustra de arquitectura general de la plataforma de microblogging como solución escalable y optimizada para lecturas a implementar.

## Consideraciones 

Para la estructura del proyecto utilicé arquitectura hexagonal.
La aplicación posee los servicios básicos del back-end que le permitirian a un usuario publicar tweets, seguir a otros usuarios y ver el timeline de tweets.

Para simplificar se implemento una data base in memory sin embargo en el documento de arquitectura general de una aplicación escalable se especifica el tipo de base de datos que usaria. 