# User API Go

Este proyecto implementa una API RESTful para la gestión de usuarios, construida con Go y GORM, utilizando MySQL como base de datos. Sigue un patrón de arquitectura limpia, separando las responsabilidades en controladores, modelos, rutas y conexión a la base de datos.

## Características

*   **Gestión de Usuarios**: Crear, leer, actualizar y eliminar registros de usuarios.
*   **Borrado Lógico (Soft Delete)**: Los usuarios pueden ser eliminados lógicamente, lo que significa que se marcan como eliminados pero permanecen en la base de datos.
*   **Borrado Permanente**: Opción para eliminar usuarios de forma permanente, incluyendo aquellos eliminados lógicamente.
*   **Integración con Base de Datos**: Utiliza GORM para operaciones ORM con MySQL.
*   **Enrutamiento**: Gestionado por `gorilla/mux`.
*   **Recarga en Vivo (Live Reload)**: Configurado con `Air` para la recompilación automática y el reinicio del servidor al realizar cambios en el código.
*   **Patrón Strategy**: Implementado para manejar diferentes estrategias de eliminación de usuarios (borrado lógico y permanente).

## Patrones de Diseño Implementados

### Patrón Strategy

Este proyecto implementa el **Patrón Strategy** para manejar diferentes formas de eliminar usuarios:

*   **¿Qué es?**: El patrón Strategy es un patrón de comportamiento que permite definir una familia de algoritmos, encapsular cada uno de ellos y hacerlos intercambiables. Esto permite que el algoritmo varíe independientemente de los clientes que lo utilizan.

*   **¿Por qué se eligió?**: Se eligió este patrón porque tenemos dos estrategias diferentes para eliminar usuarios (borrado lógico y borrado permanente) y queremos desacoplar estos algoritmos de los controladores que los utilizan.

*   **¿Cómo se implementó?**: 
    - Se creó una interfaz `DeleteStrategy` que define el método `Delete`.
    - Se implementaron dos estrategias concretas: `SoftDeleteStrategy` (borrado lógico) y `HardDeleteStrategy` (borrado permanente).
    - Se creó un contexto `DeleteContext` que utiliza una estrategia para realizar la operación de eliminación.
    - Los controladores utilizan el contexto con la estrategia apropiada según la operación requerida.

*   **Beneficios obtenidos**:
    - **Desacoplamiento**: Los algoritmos de eliminación están desacoplados de los controladores.
    - **Extensibilidad**: Es fácil añadir nuevas estrategias de eliminación sin modificar los controladores existentes.
    - **Mantenibilidad**: El código es más limpio y fácil de mantener al separar las responsabilidades.
    - **Testabilidad**: Es más fácil probar cada estrategia de forma aislada.

## Requisitos

*   Go (versión 1.24.4 o superior)
*   Base de Datos MySQL
*   [Air](https://github.com/cosmtrek/air) (para recarga en vivo, opcional pero recomendado para desarrollo)

## Configuración

### 1. Configuración de la Base de Datos

Este proyecto utiliza MySQL. Asegúrate de tener un servidor MySQL en ejecución. La cadena de conexión a la base de datos está definida en `db/connection.go`.

```go:db%2Fconnection.go
var DSN = "root:@tcp(localhost:3306)/user-api-go?charset=utf8mb4&parseTime=True&loc=Local"
```
Puedes modificar la variable DSN para que coincida con tu configuración de MySQL.

El archivo db/users.sql contiene el volcado SQL para la tabla users , que puedes importar en tu base de datos MySQL.

### 2. Variables de Entorno
No se requieren variables de entorno específicas más allá de la cadena de conexión a la base de datos, que está codificada para simplificar en db/connection.go .

### 3. Instalar Dependencias
Navega al directorio raíz del proyecto e instala los módulos de Go:
```bash
go mod tidy
go mod download
```
## Estructura del Proyecto

```bash
.air.toml             # Configuración de Air para recarga en vivo
├── controllers/      # Contiene las funciones de controlador para los endpoints de la API
│   ├── index.controller.go
│   └── user.controller.go
├── db/               # Conexión a la base de datos y scripts SQL
│   ├── connection.go
│   └── users.sql
├── go.mod            # Dependencias del módulo Go
├── go.sum            # Sumas de verificación del módulo Go
├── main.go           # Punto de entrada principal de la aplicación
├── models/           # Define las estructuras de datos (ej. estructura User)
│   └── User.go
└── routes/           # Define las rutas de la API y las enlaza a los controladores
    └── user.routes.go
```
## Endpoints de la API
La API expone los siguientes endpoints para la gestión de usuarios:

- GET / : Endpoint de inicio.
- GET /users : Recupera todos los usuarios.
- GET /users/{id} : Recupera un solo usuario por ID.
- POST /users : Crea un nuevo usuario.
- PUT /users/{id} : Actualiza un usuario existente por ID.
- DELETE /users/sofd/{id} : Elimina lógicamente un usuario por ID.
- DELETE /users/{id} : Elimina permanentemente un usuario por ID (incluidos los eliminados lógicamente).

## Cómo Ejecutar
### Desarrollo (con Recarga en Vivo)
Si tienes Air instalado, puedes ejecutar la aplicación con recarga en vivo:

```
air
```
Esto iniciará el servidor, y cualquier cambio que realices en los archivos Go activará automáticamente una reconstrucción y un reinicio del servidor.

### Producción
Para construir y ejecutar la aplicación sin recarga en vivo:

1. Construye el ejecutable:
go build -o main.exe .
```

2. Ejecuta el ejecutable:
.\main.exe
```

La API será accesible en http://localhost:3000 .

## Consideraciones de Seguridad
- Validación de Entrada : Asegúrate de que todos los datos entrantes sean validados para prevenir vulnerabilidades comunes como la inyección SQL o el cross-site scripting.
- Autenticación y Autorización : Implementa mecanismos adecuados de autenticación y autorización para controlar el acceso a los endpoints de la API.
- Manejo de Errores : Evita exponer información sensible en los mensajes de error.
- HTTPS : Utiliza siempre HTTPS en entornos de producción para cifrar la comunicación.
- Limitación de Tasa (Rate Limiting) : Considera implementar la limitación de tasa para prevenir abusos y ataques de denegación de servicio.
Este README.md proporciona una visión general completa de tu proyecto, su configuración y cómo ejecutarlo.

API RESTful desarrollada por:
- [Santiago Orjuela](
- [Santiago Orjuela](https://github.com/sago-code)
