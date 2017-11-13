# self-esteem

Debes tener [PostgreSQL](https://www.postgresql.org/download/) instalado y los comandos asociados `psql`, `dropdb`, `dropuser`, `createdb`, `createuser` disponibles desde tu shell favorito.

Debes tener [Go 1.9](https://golang.org/dl/) (ver **featured downloads**) instalado así como [dep](https://github.com/golang/dep) (corriendo `go get -u github.com/golang/dep/cmd/dep` ya que tengas [configurado](https://golang.org/doc/install) el `GOPATH`).

1. Corre `make init` para:
    - Crear el usuario `esteemate` en postgres
    - Crear la base de datos `self_esteem_dev`
    - Asegurarse que las bibliotecas de terceros estén presentes y listas para usarse
2. Corre `make run` para:
    - Compilar el código en el binario `self-esteem`
    - Ejecutar el binario con las variables de entorno de `.env`

El Makefile es bastante legible, quizá lo que saque de onda es que el programa no funcione si se ejecuta directamente el comando
`go run main.go`. Esto es porque las variables de entorno no se sourcean automaticamente desde Go, pero puedes sourcearlas
manualmente y luego correr el comando en cuestión sin broncas:

```
$ source .env
$ go run main.go
```

Para comenzar con una base de datos limpia hay de dos sopas:
1. Borras las tablas usando `make db-down` y las vuelves a crear usando `make db-up`
2. Borras la base de datos por completo usando `make db-kill` y vuelves a recrear todo usando `make init`

## Para el futuro

La organización de paquetes no es necesariamente la mejor. Hay muy poco código para poder determinar la manera mas conveniente
de organizar el proyecto. Los paquetes `notifications` y `estimates` en particular tienen muy poco código, mi idea para ellos
era tener en `notifications` el código para enviar notificaciones por correo electrónico y también para determinar cuando se debía
de enviar una notificación, que mensajes deberían de tener, etc; en el caso del codigo de `estimates` pudiera expandirse para
manejar toda la logica y estadística de los estimados.

El resto del código carece de abstracciones a propósito. Conforme vaya creciendo se pueden crear funciones o interfaces que nos
permitan expresar mejor las ideas. Algunos lugares en donde esto ya es bastante obvio es en el manejo de requests y responses en el paquete `server`, el manejo de errores especializados (en todo el codigo) y la interacción con la base de datos (en los paquetes `users` y `tasks`).

Actualmente solo se implementan los siguientes endpoints:

```
GET "/users"
GET "/users/:user-id"
GET "/users/:user-id/tasks"
GET "/users/:user-id/tasks/:task-id"
```
