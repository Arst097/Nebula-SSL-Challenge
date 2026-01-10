# SSL Labs CLI Analyzer
**Autor:** Sara Milena Arevalo Cristancho
**Lenguaje:** Go
**Descripción:** Herramienta de línea de comandos para consultar el estado SSL/TLS de un host usando la API de SSL Labs. Permite visualizar un resumen del escaneo, detalles por endpoint, y analizar múltiples hosts.
---
## Características
- Consulta la API de SSL Labs para cualquier dominio.
- Polling automático hasta que el análisis esté listo.
- Resumen de resultados: estado, puerto, fecha de escaneo y calificación de cada endpoint.
- Detalles de cada endpoint: IP, nombre del servidor, grado, advertencias, configuración especial, progreso y duración.
- Menú interactivo:
  1. Ver detalles de un endpoint.
  2. Analizar otro host.
  0. Salir.
- Validación de host antes de enviar la consulta y manejo de errores.

---

## Instalación
1. Clonar el repositorio:
```bash
git clone https://github.com/Arst097/Nebula-SSL-Challenge.git
cd challenge-ssl-labs
```
3. Ejecutar el programa:
```bash
go run main.go
```
## Uso
1. Ejecutar el programa y escribir el host a analizar
2. Esperar mientras el escaneo está en progreso
3. Revisar el resumen de resultados
4. Usar el menú para ver detalles de un endpoint o analizar otro host

## Ejemplo de salida
```
Host por analizar (sin http/https): facebook.com
Analizando host: facebook.com

✓ ---------- Resumen de Reporte SSL ---------- ✓
Host: facebook.com
Puerto: 443
Estado: READY
Escaneo iniciado: 2026-01-09 11:40:35
Escaneo terminado: 2026-01-09 11:44:25

Endpoints Totales: 2
[1] 57.144.252.1 | Grado: B | Progreso: 100%
[2] 2a03:2880:f37e:1:face:b00c:0:25de | Grado: B | Progreso: 100%
------------------------------------------------

▼ ▼ ------------ Menú ------------ ▼ ▼
1. Ver detalles de un endpoint
2. Analizar otro host
0. Salir
Seleccione una opción:
1
 ➡︎ Endpoint a revisar: 2
------------------------------------------------

● -------- Endpoint 2 -------- ●
IP: 2a03:2880:f37e:1:face:b00c:0:25de
Nombre del Servidor:  edge-star-mini6-shv-01-atl3.facebook.com
Grado: B
Advertencias detectadas: false
Config. Especial:  false
Progreso: 100%
Duración:  1m54.28s
------------------------------------------------

```

## Proceso de Desarrollo
Durante el desarrollo, se tomaron varias decisiones:
- Inicialmente, el JSON de la API se parseaba a un map[string]interface{}, pero luego se creó un struct para mapear solo los campos relevantes y facilitar la lectura de resultados.
- Se implementó polling para manejar análisis en progreso de forma automática, mostrando porcentajes de progreso por endpoint.
- Se creó un menú interactivo y funciones separadas (analizeHost, Summary, DetailsEndpoint) para mantener main.go limpio y organizado.
- Se añadió validación de host y manejo de errores para entradas inválidas o hosts que no existen.
- Se trabajó con tiempos de inicio y fin del escaneo (StartTime, TestTime) y se convirtió la duración de los endpoints a un formato legible (time.Duration).
- Este proyecto maneja un flujo completo desde la entrada del usuario, manejo de errores y validación, hasta la visualización de resultados de forma clara y práctica.

## Consideraciones
- El progreso de los endpoints se muestra en porcentaje durante el escaneo.
- La duración de los endpoints se muestra en time.Duration (milisegundos convertidos a minutos/segundos).
- Las advertencias (HasWarnings) y configuraciones especiales (IsExceptional) se muestran como true/false.
- El programa está diseñado para recibir solo hostnames válidos, no URLs completas con esquema http/https.

## Nota final
Este proyecto fue desarrollado íntegramente de forma local, usando Go y la API pública de SSL Labs. Se busco tener en cuenta tanto el manejo de datos de la API como la experiencia de usuario en CLI, priorizando la claridad de la información, la validación de entrada y la organización del código.

