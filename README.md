# TO DO LIST API

### Descripción: 

Una API que permita realizar operaciones CRUD (Create, Read, Update, Delete) sobre una lista de tareas o usuarios.

### Estructura

```bash
/to-do-api
│
├── main.go
└── go.mod
```

## 🚀 Instrucciones para correrlo:

1. Inicializa el proyecto

```bash
go mod init to-do-api
```

2. Instala el paquete `gorilla/mux`:

```bash
go get -u github.com/gorilla/mux
```

3. Para ejecutar la API:

```bash
go run main.go
```

## ✅ Probando la API

- GET: Obtener todas las tareas

```bash
curl http://localhost:8000/tasks
```

- POST: Crear una tarea nueva

```bash
curl -X POST -H "Content-Type: application/json" -d '{"id":"3","title":"Probar API","details":"Verificar endpoints con curl"}' http://localhost:8000/tasks
```

