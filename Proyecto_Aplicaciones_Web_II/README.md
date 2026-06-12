# Proyecto_Aplicaciones_Web_II

# Proyecto: Pesca-Directa Tarqui (Del Mar a la Mesa)

## Descripción

Pesca-Directa Tarqui es una plataforma orientada a digitalizar la comercialización de productos pesqueros artesanales en el Puerto de Tarqui (Manta, Ecuador).

El sistema conecta pescadores, clientes comerciales y operadores logísticos mediante una arquitectura basada en dominios de negocio, permitiendo gestionar capturas, inventario disponible, pedidos y distribución de productos pesqueros.

El objetivo principal es reducir la dependencia de intermediarios, mejorar la trazabilidad del producto y optimizar la conexión entre oferta y demanda.

---

## Problema

Actualmente la comercialización pesquera artesanal presenta varios inconvenientes:

* Los pescadores venden sus capturas a intermediarios que establecen precios poco favorables.
* Los compradores carecen de información actualizada sobre disponibilidad, cantidad y frescura de los productos.
* No existe una trazabilidad adecuada del producto desde la captura hasta la entrega final.
* La logística de distribución se gestiona de forma manual e informal.
* La información comercial depende de llamadas telefónicas, contactos personales y grupos de WhatsApp.

---

## Solución

Pesca-Directa Tarqui propone una plataforma API REST organizada en tres dominios principales:

### Gestión de Pesca

Administra la operación pesquera artesanal.

Entidades principales:

* Usuario
* Pescador
* Embarcación
* Especie
* Captura
* Bodega
* Stock

Funciones:

* Registro de pescadores.
* Administración de embarcaciones.
* Registro de capturas diarias.
* Control de especies disponibles.
* Gestión de inventario y stock.

---

### Gestión de Pedidos

Administra la demanda comercial.

Entidades principales:

* Cliente
* Pedido
* DetallePedido

Funciones:

* Registro de clientes.
* Gestión de pedidos.
* Control de productos solicitados.
* Seguimiento de compras.

Tipos de cliente:

* Restaurante
* Intermediario
* Mayorista

---

### Ruta de Distribución

Administra la logística de entrega.

Entidades principales:

* Ruta
* Punto
* Transportista
* EntregaPedido

Funciones:

* Definición de rutas de distribución.
* Administración de puntos de recorrido.
* Asignación de transportistas.
* Seguimiento de entregas.
* Control de estados logísticos.

---

## Regla de Negocio No-CRUD

### Sistema de Alertas por Escasez

El sistema analiza la disponibilidad histórica de especies pesqueras.

Ejemplo:

* Si la cantidad disponible de una especie es menor al 30% del promedio de los últimos 7 días se genera una alerta de escasez.
* Los clientes reciben notificaciones sobre disponibilidad limitada.
* Se puede sugerir un precio dinámico basado en la oferta disponible.

---

## Stack Tecnológico

* Lenguaje principal: Go (Golang)
* Framework HTTP: net/http
* Router: Chi Router
* Base de datos: SQLite
* ORM: GORM
* Generación de consultas tipadas: SQLC
* Autenticación: JWT
* Arquitectura: Clean Architecture
* Inyección de Dependencias
* Docker (planificado)
* Git + GitHub

---

## Arquitectura General

pesca-directa-tarqui/

├── cmd/
│ └── api/
│ └── main.go
│
├── internal/
│ ├── handlers/
│ ├── models/
│ ├── services/
│ ├── repositories/
│ └── storage/
│
├── migrations/
├── docs/
├── go.mod
├── go.sum
└── README.md

---

## Dominios del Proyecto

### Gestión de Pesca

Responsable: Anthony Macias

* Pescadores
* Embarcaciones
* Especies
* Capturas
* Bodegas
* Stock

---

### Gestión de Pedidos

Responsable: Michelle Salazar

* Clientes
* Pedidos
* DetallePedido

---

### Ruta de Distribución

Responsable: Madelyn Zambrano

* Rutas
* Puntos
* Transportistas
* EntregaPedido

---

## Flujo General del Sistema

Captura de Pesca
↓
Registro de Inventario (Stock)
↓
Creación de Pedido
↓
Asignación de Ruta
↓
Asignación de Transportista
↓
Seguimiento de Entrega
↓
Recepción por Cliente

---

## Objetivo Académico

Aplicar conceptos de:

* Desarrollo Backend en Go.
* Diseño de APIs REST.
* SQLC y GORM.
* Arquitectura por capas.
* Inyección de dependencias.
* Modelado de dominio.
* Diseño de bases de datos relacionales.
* Implementación de reglas de negocio reales.

---

## Autores

* Anthony Joel Macias Macias
* Ilaria Michelle Salazar Rezabala
* Madelyn Elisa Zambrano Cevallos

Universidad Laica Eloy Alfaro de Manabí (ULEAM)

Tecnologías de la Información

Aplicaciones Web II
