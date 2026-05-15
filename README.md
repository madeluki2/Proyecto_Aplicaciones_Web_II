# Proyecto_Aplicaciones_Web_II

# Proyecto: Pesca-Directa Tarqui (Del Mar a la Mesa)

## Descripción
Plataforma web que conecta directamente a pescadores de Tarqui(Manta) con restaurantes y cevicherías locales.
El sistema busca eliminar intermediarios, mejorar los ingresos de los pescadores y garantizar pescado fresco para los negocios gastronómicos.

## Problema
- Los pescadores venden a intermediarios que imponen precios bajos sin negociación.
- Los restaurantes no saben qué pescado llega cada mañana al puerto.
- La falta de transparencia genera pérdidas para ambos lados.

### Solución
Un sistema con:
- **Módulo de Especies y Catálogo del Mar**: se encarga de registrar y administrar las especies disponibles en el puerto.
- **Módulo de Pescadores y Embarcaciones**: gestiona la informaciín de los pescadores artesanales y sus embarcaciones.
- **Módulo de Restaurantes y Cevicherías**: administra los perfiles de los restaurantes y sus preferencias de compra.

## Regla de Negocio No-CRUD
Sistema de alertas por *"Captura del Día"* con precios dinámicos.
Ejemplo:
- Si la cantidad disponible de una especie es menor al 30% del promedio de los últimos 7 días el sistema envía una alerta de escasez.
- Precios varían según hora de llegada: antes de las 06:00 premium, después de las 09:00 descuento.
- Restaurantes reciben notificación con especie, cantidad y precio sugerido.

---

## Stack Tecnológico
- **Lenguaje principal**: Go (Golang)
-  **Framework web**: net/http (posible integración futura con Chi router)
-  **Base de datos**: SQLite, PostgreSQL
-  **ORM**: GORM (librería más usada en Go para hacer ORM).
-  **Autenticación**: JWT
-  **Arquitectura**: Clean Architecture con separación en `cmd/` y `internal/`, será hadler --> service --> repository, con interfaces e inyeccion de dependencias.
-  **Control de versiones**: Git + GitHub

---

## Estructura inicial del proyecto
pesca-directa-tarqui/
 ├── cmd/
 │   └── api/
 │       └── main.go
 ├── internal/
 │   ├── models/         # Entidades de dominio (Especie, Pescador, Restaurante, User)
 │   ├── handlers/       # Endpoints REST
 │   └── storage/        # Conexión a BD y repositorios
 ├── go.mod              # Módulo inicializado
 ├── .gitignore          # Configuración para Go
 └── README.md           # Descripción del proyecto

## Diagrama de módulos
cmd/api
   └── Punto de entrada
        └── internal/ Lógica de negocio
             ├── especies/ → Catálogo del mar
             ├── pescadores/ → Perfil y embarcaciones
             └── restaurantes/ → Perfil y preferencias

Relaciones:
- `especies` provee catálogo y precios base.
- `pescadores` registran capturas y embarcaciones.
- `restaurantes` consultan disponibilidad y preferencias.
- Todos se integran con la lógica de precios dinámicos (alertas de Captura del Día).
