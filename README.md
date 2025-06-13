# User API Go

Este proyecto implementa una API RESTful para la gestión de usuarios, construida con Go y GORM, utilizando MySQL como base de datos. Sigue un patrón de arquitectura limpia, separando las responsabilidades en controladores, modelos, rutas y conexión a la base de datos.

## Características

*   **Gestión de Usuarios**: Crear, leer, actualizar y eliminar registros de usuarios.
*   **Borrado Lógico (Soft Delete)**: Los usuarios pueden ser eliminados lógicamente, lo que significa que se marcan como eliminados pero permanecen en la base de datos.
*   **Borrado Permanente**: Opción para eliminar usuarios de forma permanente, incluyendo aquellos eliminados lógicamente.
*   **Integración con Base de Datos**: Utiliza GORM para operaciones ORM con MySQL.
*   **Enrutamiento**: Gestionado por `gorilla/mux`.
*   **Recarga en Vivo (Live Reload)**: Configurado con `Air` para la recompilación automática y el reinicio del servidor al realizar cambios en el código.

## Requisitos

*   Go (versión 1.24.4 o superior)
*   Base de Datos MySQL
*   [Air](https://github.com/cosmtrek/air) (para recarga en vivo, opcional pero recomendado para desarrollo)

## Configuración

### 1. Configuración de la Base de Datos

Este proyecto utiliza MySQL. Asegúrate de tener un servidor MySQL en ejecución. La cadena de conexión a la base de datos está definida en `db/connection.go`.

```go:db%2Fconnection.go
// ... existing code ...
var DSN = "root:@tcp(localhost:3306)/user-api-go?charset=utf8mb4&parseTime=True&loc=Local"
// ... existing code ...
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
Trae

Claro, aquí tienes el archivo README.md generado nuevamente para tu proyecto:

```
# User API Go

This project implements a RESTful API for 
managing users, built with Go and GORM, 
using MySQL as the database. It follows a 
clean architecture pattern, separating 
concerns into controllers, models, routes, 
and database connection.

## Features

*   **User Management**: Create, read, 
update, and delete user records.
*   **Soft Deletion**: Users can be 
soft-deleted, meaning they are marked as 
deleted but remain in the database.
*   **Permanent Deletion**: Option to 
permanently delete users, including those 
soft-deleted.
*   **Database Integration**: Uses GORM for 
ORM operations with MySQL.
*   **Routing**: Handled by `gorilla/mux`.
*   **Live Reload**: Configured with `Air` 
for automatic recompilation and server 
restart on code changes.

## Requirements

*   Go (version 1.24.4 or higher)
*   MySQL Database
*   [Air](https://github.com/cosmtrek/air) 
(for live reloading, optional but 
recommended for development)

## Setup

### 1. Database Configuration

This project uses MySQL. Ensure you have a 
MySQL server running. The database 
connection string is defined in `db/
connection.go`.

```go:db%2Fconnection.go
// ... existing code ...
var DSN = "root:@tcp(localhost:3306)/
user-api-go?charset=utf8mb4&parseTime=True&
loc=Local"
// ... existing code ...
```
You can modify the DSN variable to match your MySQL configuration.

The db/users.sql file contains the SQL dump for the users table, which you can import into your MySQL database.

### 2. Environment Variables
No specific environment variables are required beyond the database connection string, which is hardcoded for simplicity in db/connection.go .

### 3. Install Dependencies
Navigate to the project root directory and install the Go modules:

```
go mod tidy
```
## Project Structure
```
.air.toml             # Air configuration 
for live reloading
├── controllers/      # Contains handler 
functions for API endpoints
│   ├── index.controller.go
│   └── user.controller.go
├── db/               # Database connection 
and SQL scripts
│   ├── connection.go
│   └── users.sql
├── go.mod            # Go module 
dependencies
├── go.sum            # Go module checksums
├── main.go           # Main application 
entry point
├── models/           # Defines data 
structures (e.g., User struct)
│   └── User.go
└── routes/           # Defines API routes 
and links them to controllers
    └── user.routes.go
```
## API Endpoints
The API exposes the following endpoints for user management:

- GET / : Home endpoint.
- GET /users : Retrieve all users.
- GET /users/{id} : Retrieve a single user by ID.
- POST /users : Create a new user.
- PUT /users/{id} : Update an existing user by ID.
- DELETE /users/sofd/{id} : Soft delete a user by ID.
- DELETE /users/{id} : Permanently delete a user by ID (including soft-deleted ones).
## How to Run
### Development (with Live Reload)
If you have Air installed, you can run the application with live reloading:

```
air
```
This will start the server, and any changes you make to the Go files will automatically trigger a rebuild and server restart.

### Production
To build and run the application without live reloading:

1. Build the executable:
go build -o main.exe .
```

2. Run the executable:
.\main.exe
```

The API will be accessible at http://localhost:3000 .

## Security Considerations
- Input Validation : Ensure all incoming data is validated to prevent common vulnerabilities like SQL injection or cross-site scripting.
- Authentication and Authorization : Implement proper authentication and authorization mechanisms to control access to API endpoints.
- Error Handling : Avoid exposing sensitive information in error messages.
- HTTPS : Always use HTTPS in production environments to encrypt communication.
- Rate Limiting : Consider implementing rate limiting to prevent abuse and denial-of-service attacks.
This README.md provides a comprehensive overview of your project, its setup, and how to run it.

sago-code

Copy

Delete

en español mano

Trae

Claro, aquí tienes el archivo README.md generado nuevamente para tu proyecto, en español:

Puedes modificar la variable DSN para que coincida con tu configuración de MySQL.

El archivo db/users.sql contiene el volcado SQL para la tabla users , que puedes importar en tu base de datos MySQL.

### 2. Variables de Entorno
No se requieren variables de entorno específicas más allá de la cadena de conexión a la base de datos, que está codificada para simplificar en db/connection.go .

### 3. Instalar Dependencias
Navega al directorio raíz del proyecto e instala los módulos de Go:

## Estructura del Proyecto
```
.air.toml             # Configuración de 
Air para recarga en vivo
├── controllers/      # Contiene las 
funciones de controlador para los endpoints 
de la API
│   ├── index.controller.go
│   └── user.controller.go
├── db/               # Conexión a la base 
de datos y scripts SQL
│   ├── connection.go
│   └── users.sql
├── go.mod            # Dependencias del 
módulo Go
├── go.sum            # Sumas de 
verificación del módulo Go
├── main.go           # Punto de entrada 
principal de la aplicación
├── models/           # Define las 
estructuras de datos (ej. estructura User)
│   └── User.go
└── routes/           # Define las rutas de 
la API y las enlaza a los controladores
    └── user.routes.go
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
