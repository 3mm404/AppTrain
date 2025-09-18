📘 API GYMGO – Documentación para Pruebas

✅ Registro de Usuario (Cliente)

Endpoint: POST  http://127.0.0.1:3000/users/register
Descripción: Crea un nuevo usuario cliente.
Requiere contraseña obligatoria.
🧪 Ejemplo de JSON:

{
  "nombre": "Carlos Pérez",
  "email": "carlos@test.com",
  "telefono": "9998887777",
  "password": "miclave123"
}

🔐 Login de Usuario

Endpoint: POST http://127.0.0.1:3000/login/user
Descripción: Autentica a un usuario por teléfono o email.
🧪 Ejemplo de JSON:

{
  "Identificador": "9611234567",
  "password": "micontraseña123"
}

🏋️ Registro de Gimnasio

Endpoint: POST http://127.0.0.1:3000/register/gym
Descripción: Registra un nuevo gimnasio con ubicación y contraseña.
🧪 Ejemplo de JSON:

{
  "nombre": "Gym MaxPower",
  "direccion": "Av. Reforma #123",
  "telefono": "9611234567",
  "email": "gymmax@example.com",
  "foto": "https://ejemplo.com/foto.jpg",
  "latitud": 16.7533,
  "longitud": -93.1156,
  "aprobado": false,
  "creadoEn": "2025-05-03T10:00:00Z",
  "password": "supersecreto123"
}

👥 Registro de Usuario por el Gimnasio (Empleado o Cliente)


📊 Estado de Membresía del Usuario

Endpoint: GET http://127.0.0.1:3000/users/membership
Descripción: Obtiene el estado actual de la membresía del usuario, incluyendo si está asociado a un gimnasio, si su membresía está activa y las fechas de inicio y fin de la membresía.
Requiere autenticación: Sí (JWT)
Requiere rol: Cliente

Ejemplo de respuesta exitosa (membresía activa):
```
{
    "gym_id": 1,
    "gym_name": "",
    "status": "activo",
    "start_date": "2025-06-13T00:00:00Z",
    "end_date": "2025-07-13T00:00:00Z"
}
```

Ejemplo de respuesta (sin membresía activa):
```
{
    "gym_id": 1,
    "gym_name": "",
    "status": "inactivo",
    "start_date": "",
    "end_date": ""
}
```

Ejemplo de respuesta (sin gimnasio asociado):
```
{
    "gym_id": 0,
    "gym_name": "",
    "status": "inactivo",
    "start_date": "",
    "end_date": ""
}
```

Ejemplo de error:
```
{
    "error": "Error al obtener la información del usuario"
}
```


👥 Registro de Usuario por el Gimnasio (Empleado o Cliente)

Endpoint: POST http://127.0.0.1:3000/register/gym/user
Descripción: Permite al gimnasio registrar usuarios internos.
Nota: El campo password se deja vacío y será actualizado por correo o SMS.
🧪 Ejemplo de JSON:

{
  "nombre": "Lucía Méndez",
  "email": "lucia@gymadmin.com",
  "telefono": "9617654321",
  "fecha_nacimiento": "1995-04-03",
  "password": "",                            <------ Aqui el password es null por defauld
  "foto": "https://miapp.com/fotos/lucia.jpg",
  "tipo_usuario": "cliente",
  "codigo_verificacion": "123456"
}

{
  "nombre": "Membresía Semestral",
  "descripcion": "Acceso completo durante seis meses con descuentos especiales.",
  "precio": 600.00,
  "duracion_dias": 180,
  "status": "activo"
}


📋 Listar Membresías del Gimnasio

Endpoint: GET http://127.0.0.1:3000/gyms/mymemberships
Descripción: Muestra todas las membresías asociadas al gimnasio autenticado.
🧪 Ejemplo de respuesta:

[
  {
    "id": 3,
    "gym": "Gym FitLive",
    "descripcion": "Acceso a todas las áreas con clases ilimitadas.",
    "precio": 3200,
    "duracion_dias": 90,
    "status": "activo"
  }
]

✏️ Actualizar Membresía

Endpoint: PUT http://127.0.0.1:3000/membership-types/{id}
Descripción: Actualiza los datos de una membresía existente por su ID.
📌 Reemplaza {id} por el ID de la membresía.

Ejemplo de URL: http://127.0.0.1:3000/membership-types/3
🧪 Ejemplo de JSON:

{
  "nombre": "Membresía Premium",
  "descripcion": "Acceso a todas las áreas con clases ilimitadas.",
  "precio": 3200,
  "duracion_dias": 90,
  "status": "activo"
}


❌ Eliminar Membresía

Endpoint: DELETE http://127.0.0.1:3000/eliminar/membresias/{id}
Descripción: Elimina una membresía por su ID.
📌 Reemplaza {id} por el ID correspondiente.

Ejemplo de URL: http://127.0.0.1:3000/eliminar/membresias/3









FUNCIONES DE LA APP

1. Usuarios
Tipos de usuarios:
• Cliente: Persona que usa la app para gestionar su membresía, pagar, y asistir al gimnasio.

• Empleado del gimnasio: Persona encargada de gestionar el gimnasio (registro de asistencias, pagos, etc.).

• Administrador del gimnasio: Tiene acceso completo a todas las funciones del gimnasio (agregar membresías, ver pagos, gestionar usuarios, etc.).

• Super Admin (opcional): Control total sobre toda la plataforma, puede supervisar y gestionar todos los gimnasios y sus administradores.


Funciones principales:
◇ Registro/Login: Acceso mediante correo y contraseña.

◇ Visualización: El usuario puede ver sus pagos, membresías activas, historial de asistencias.

◇ Entrada y salida: Los usuarios pueden registrar su entrada y salida del gimnasio por QR o manualmente desde la app.

◇ Consultar gimnasios cercanos: Basado en la ubicación del usuario (latitud y longitud).

◇ Pagar membresías: Métodos de pago: tarjeta, transferencia o efectivo.



2. Gimnasios
Funciones del gimnasio:
◇ Registro del gimnasio: Un gimnasio se registra en la plataforma, pero está en un estado de "pendiente" hasta que sea aprobado.

◇ Agregar tipos de membresía: El gimnasio puede crear varios tipos de membresías (diaria, mensual, anual, etc.).

◇ Gestionar usuarios: Ver qué usuarios han comprado qué membresías.

◇ Registrar asistencias: Puede registrar asistencias manualmente o utilizando un sistema de QR para verificar la entrada y salida de los usuarios.

◇ Ver historial de pagos: Los gimnasios pueden ver todos los pagos realizados, incluyendo el estado (pendiente, completado, rechazado).

◇ Agregar pagos manualmente: Si el usuario paga en efectivo, los administradores pueden registrar el pago manualmente desde el panel.



3. Pagos y Membresías
Métodos de pago:
◇ Tarjeta: Integración con una pasarela de pago.

◇ Transferencia: El pago por transferencia requiere revisión manual.

◇ Efectivo: El pago en efectivo se activa manualmente desde el panel de administración.


Lógica de pagos:
◇ Cada pago realizado crea un registro en la tabla payments.

◇ Activación de membresía: Solo si el pago es completado (estado "completado") se activa la membresía del usuario en la tabla user_memberships.

◇ Si el usuario no está registrado, el gimnasio puede ingresar los datos manualmente y asociar la membresía a ese usuario.



4. Asistencias
Lógica de asistencias:
◇ Los usuarios pueden registrar su entrada y salida en el gimnasio mediante un sistema de QR o manualmente desde la app.

◇ La asistencia es consultable tanto por el usuario como por el gimnasio (empleados).

◇ Propósito de las asistencias: Servir como control de presencia y generar estadísticas (por ejemplo, número de visitas por usuario).



5. Ubicación y Descubrimiento de Gimnasios
Funciones de localización:
◇ Lista de gimnasios cercanos: Los gimnasios pueden ser listados por proximidad utilizando latitud y longitud.

◇ Filtros de búsqueda: Los usuarios pueden filtrar los gimnasios por ciudad, nombre o tipo de gimnasio.

◇ Información del gimnasio: Al seleccionar un gimnasio, los usuarios pueden ver detalles como su información, tipos de membresía, precios y la opción de realizar pagos.



Resumen de Interacciones entre Componentes:
1. Usuarios pueden registrarse, pagar membresías, hacer consultas y registrar asistencias.

2. Gimnasios gestionan sus usuarios (clientes y empleados), controlan asistencias y pagos, y ofrecen tipos de membresías.

3. Pagos son gestionados por la plataforma, con validación de estado antes de activar membresías.

4. Asistencias se registran por QR o manualmente para control y estadísticas.

5. Ubicación y descubrimiento de gimnasios se facilita con un sistema de búsqueda basado en proximidad geográfica.

