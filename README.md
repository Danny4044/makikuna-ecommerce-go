# MAKIKUNA – Sistema de Gestión de E-commerce

**Tienda online de pulseras y collares masculinos**

Proyecto académico de planificación y desarrollo inicial de un sistema de comercio electrónico, elaborado como parte del **Aprendizaje Autónomo 1 de la Unidad 1: Programación con Go**.

> **Estado actual:** este repositorio contiene una estructura inicial y un prototipo funcional de algunas reglas del negocio. No representa todavía una tienda en línea completa ni procesa pagos reales.

---

## 1. Descripción del proyecto

**MAKIKUNA** es una propuesta de tienda virtual especializada en la comercialización de accesorios masculinos, principalmente pulseras y collares. El sistema busca organizar las operaciones principales del negocio, facilitar la administración del catálogo y del inventario, y permitir que los clientes puedan seleccionar productos y generar pedidos.

La solución se plantea con **Go (Golang)** para el desarrollo del backend. El código se organiza en paquetes y funciones con responsabilidades específicas, procurando que las reglas del negocio sean reutilizables y fáciles de probar.

## 2. Objetivo general

Diseñar y desarrollar progresivamente un sistema de gestión de e-commerce para MAKIKUNA que permita administrar productos, inventario, clientes, carritos de compra, pedidos y operaciones relacionadas con la venta de pulseras y collares masculinos.

## 3. Objetivos específicos

- Organizar el catálogo de pulseras y collares por producto y categoría.
- Definir reglas para validar la disponibilidad de los productos y las cantidades solicitadas.
- Establecer el proceso de registro de clientes y generación de pedidos.
- Proponer una estructura modular mediante paquetes de Go.
- Aplicar funciones, condicionales, ciclos y pruebas para implementar y verificar reglas del negocio.

## 4. Módulos contemplados

| Módulo | Funcionalidad prevista |
|---|---|
| Productos | Registrar, consultar, actualizar y organizar pulseras y collares. |
| Inventario | Controlar existencias y validar disponibilidad. |
| Clientes | Registrar y consultar información de compradores. |
| Carrito de compras | Agregar artículos y calcular subtotales. |
| Pedidos | Registrar compras y controlar sus estados. |
| Pagos | Registrar pagos simulados en la primera versión. |
| Envíos | Administrar información y seguimiento de entregas. |
| Reportes | Consultar información básica de ventas y pedidos. |

Estos módulos representan el alcance propuesto del sistema. No todos se encuentran implementados en el prototipo actual.

## 5. Tecnologías y herramientas

- **Go 1.22:** lenguaje definido en el archivo `go.mod`.
- **Biblioteca estándar de Go:** incluye `net/http` para el servidor, `fmt` para mensajes y errores, y `log` para el registro de actividad.
- **Paquete `testing`:** utilizado para las pruebas unitarias.
- **Git y GitHub:** propuestos para el control de versiones y el alojamiento del repositorio.
- **PostgreSQL:** motor de base de datos considerado para una siguiente etapa; todavía no está conectado al prototipo.

El prototipo actual no utiliza dependencias externas de Go. Los paquetes `pgx` para PostgreSQL y `bcrypt` para el manejo de contraseñas son alternativas consideradas para futuras etapas, pero no están instaladas ni integradas en esta versión.

## 6. Funcionalidades que ya contiene el prototipo

- **Servidor HTTP básico:** se inicia en el puerto `8080`.
- **Comprobación de estado:** expone la ruta `GET /health`, que responde indicando que la API de MAKIKUNA está disponible.
- **Modelo de producto:** define datos básicos como identificador, nombre, categoría, precio, stock y estado activo.
- **Cálculo del carrito:** la función `CalcularSubtotal` recorre los artículos, verifica que el producto exista, que esté activo, que la cantidad sea válida y que exista stock suficiente. El subtotal se maneja en centavos para evitar trabajar con valores monetarios de punto flotante.
- **Estados de pedidos:** la función `PuedeCambiarEstado` valida transiciones permitidas entre estados mediante una estructura `switch`.
- **Pruebas unitarias:** existen pruebas para el cálculo del subtotal, la validación de stock y las transiciones de estado de pedidos.

### Funcionalidades pendientes

La administración completa del catálogo, el registro y autenticación de clientes, la persistencia en PostgreSQL, la creación integral de pedidos, los pagos simulados, la gestión de envíos y los reportes todavía requieren desarrollo.

## 7. Estructura del repositorio

```text
makikuna-ecommerce-go/
├── cmd/
│   └── server/
│       └── main.go              # Punto de entrada del servidor HTTP
├── internal/
│   ├── producto/
│   │   └── modelo.go            # Modelo básico de producto
│   ├── carrito/
│   │   ├── calculo.go           # Cálculo y validación del subtotal
│   │   └── calculo_test.go      # Pruebas del carrito
│   └── pedido/
│       ├── estado.go            # Reglas de transición de estados
│       └── estado_test.go       # Pruebas de estados del pedido
├── docs/
│   ├── alcance.md               # Alcance inicial del sistema
│   └── arquitectura.png         # Diagrama de arquitectura
├── .env.example                 # Ejemplo de variables para futuras etapas
├── .gitignore                  # Archivos excluidos de Git
├── go.mod                       # Módulo y versión de Go
└── README.md                    # Documentación del proyecto
```

## 8. Requisitos previos

- Tener instalado **Go 1.22 o una versión compatible**.
- Tener Git instalado si se desea trabajar con control de versiones.
- Contar con una terminal o editor de código, por ejemplo Visual Studio Code.

Puedes comprobar la instalación de Go con:

```bash
go version
```

## 9. Cómo ejecutar el prototipo

1. Abrir una terminal en la carpeta raíz `makikuna-ecommerce-go`.
2. Ejecutar el servidor:

```bash
go run ./cmd/server
```

3. Abrir en el navegador la siguiente dirección:

```text
http://localhost:8080/health
```

La respuesta esperada es:

```text
MAKIKUNA API: disponible
```

Para detener el servidor, presionar `Ctrl + C` en la terminal.

## 10. Cómo ejecutar las pruebas

Desde la carpeta raíz del proyecto, ejecutar:

```bash
go test ./...
```

Para compilar los paquetes del proyecto:

```bash
go build ./...
```

## 11. Aplicación de los contenidos de la Unidad 1

El prototipo relaciona los fundamentos estudiados en la Unidad 1 con necesidades concretas de una tienda de comercio electrónico:

- **Sintaxis y tipos de datos:** estructuras para representar productos, artículos del carrito y estados de pedidos.
- **Condicionales `if` y `switch`:** validación de productos, cantidades, existencias y cambios de estado de pedidos.
- **Ciclo `for`:** recorrido de los artículos del carrito para calcular el subtotal.
- **Funciones:** encapsulación de reglas de negocio en operaciones reutilizables, como `CalcularSubtotal` y `PuedeCambiarEstado`.
- **Paquetes:** separación del código por responsabilidades (`producto`, `carrito` y `pedido`).
- **Pruebas:** verificación de los resultados y de las condiciones de error.

## 12. Publicación del proyecto en GitHub

Crear primero un repositorio vacío en GitHub con el nombre `makikuna-ecommerce-go`. Después, desde la carpeta local del proyecto, ejecutar:

```bash
git init
git add .
git commit -m "chore: estructura inicial de MAKIKUNA"
git branch -M main
git remote add origin https://github.com/USUARIO/makikuna-ecommerce-go.git
git push -u origin main
```

Reemplazar `USUARIO` por el nombre real de la cuenta de GitHub. Si el repositorio remoto ya tiene un primer commit o README, revisar las instrucciones de GitHub antes de realizar el primer `push`.

**Seguridad:** no subir contraseñas, claves privadas, tokens ni archivos `.env` con credenciales reales. El archivo `.env.example` es únicamente una plantilla.

## 13. Alcance y trabajo futuro

El propósito de esta primera etapa es establecer la planificación, la arquitectura inicial y una base de código que demuestre la aplicación de los contenidos de Go. Como siguientes pasos se plantea completar los módulos, conectar una base de datos, implementar autenticación y controles de acceso, ampliar las pruebas y desarrollar la interfaz de la tienda.

La integración con pasarelas de pago reales, los servicios automáticos de empresas de envío y el despliegue productivo quedan fuera del alcance de esta versión inicial.

## 14. Información académica

- **Proyecto:** MAKIKUNA – Sistema de Gestión de E-commerce.
- **Actividad:** Aprendizaje Autónomo 1 – Selección de Sistemas de Gestión Empresarial.
- **Unidad:** Unidad 1 – Programación con Go.
- **Autor:** [Escribir nombre del estudiante].
- **Institución:** [Escribir nombre de la institución].
- **Docente:** [Escribir nombre del docente].
- **Fecha:** [Escribir fecha de entrega].
