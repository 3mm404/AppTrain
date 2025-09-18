# Documentación del Backend de Trainify

## Arquitectura General

Trainify es una aplicación de gestión de gimnasios construida con Go y Fiber. La arquitectura sigue una estructura modular y limpia, facilitando la mantenibilidad y escalabilidad.

### Estructura de Directorios

```
Trainify/
├── client/              # Rutas de usuario
├── config/              # Configuración de la aplicación
├── database/            # Conexión y migraciones de la base de datos
├── gym/                 # Rutas y lógica relacionada con el gimnasio
├── middleware/          # Middleware personalizado
├── models/              # Modelos de datos
└── main.go              # Punto de entrada de la aplicación
```

## Configuración

### Variables de Entorno
El proyecto utiliza las siguientes variables de entorno (definidas en `.env`):

- `SUPABASE_HOST`: Host de la base de datos Supabase
- `SUPABASE_USER`: Usuario de la base de datos
- `SUPABASE_PASSWORD`: Contraseña de la base de datos
- `SUPABASE_DB`: Nombre de la base de datos
- `SUPABASE_PORT`: Puerto de la base de datos
- `PORT`: Puerto de la aplicación (por defecto 3000)

## Modelos de Datos

### 1. Usuario (`user.model.go`)
- ID
- Nombre
- Email
- Contraseña (hash)
- Teléfono
- Fecha de creación
- Fecha de actualización

### 2. Tipo de Membresía (`membershiptype.model.go`)
- ID
- Nombre
- Precio
- Duración
- Descripción
- Estado

### 3. Membresía de Usuario (`userMembership.model.go`)
- ID
- ID de Usuario
- ID de Tipo de Membresía
- Fecha de inicio
- Fecha de fin
- Estado

### 4. Gimnasio (`gym.model.go`)
- ID
- Nombre
- Dirección
- Teléfono
- Email
- Estado

### 5. Empleado de Gimnasio (`gymemployee.model.go`)
- ID
- ID de Usuario
- ID de Gimnasio
- Rol
- Estado

### 6. Asistencia (`attendance.model.go`)
- ID
- ID de Usuario
- Fecha
- Estado

### 7. Pago (`payment.model.go`)
- ID
- ID de Usuario
- ID de Membresía
- Monto
- Fecha
- Método de pago

## Rutas API

### Rutas de Usuario
- `POST /api/user/register`: Registro de nuevo usuario
- `POST /api/user/login`: Autenticación de usuario
- `GET /api/user/profile`: Obtener perfil de usuario
- `PUT /api/user/profile`: Actualizar perfil de usuario

### Rutas de Gimnasio
- `GET /api/gym`: Obtener lista de gimnasios
- `POST /api/gym`: Crear nuevo gimnasio
- `GET /api/gym/:id`: Obtener gimnasio específico
- `PUT /api/gym/:id`: Actualizar gimnasio
- `DELETE /api/gym/:id`: Eliminar gimnasio

## Middleware

- `logger`: Registro de todas las solicitudes HTTP
- `cors`: Soporte para peticiones cross-origin
- `auth`: Autenticación de usuarios

## Base de Datos

La aplicación utiliza Supabase como base de datos, proporcionando:
- Autenticación
- Base de datos PostgreSQL
- Almacenamiento
- Real-time subscriptions

## Ejecución del Proyecto

1. Configurar las variables de entorno en `.env`
2. Ejecutar migraciones de la base de datos
3. Iniciar el servidor con: `go run main.go`

El servidor escuchará en el puerto especificado en las variables de entorno (por defecto 3000)
