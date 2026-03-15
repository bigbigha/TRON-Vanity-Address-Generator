# TRON Vanity Address Generator

Un generador de direcciones TRON vanity (con patrón personalizado) ligero y **de código abierto**.  
Permite crear direcciones TRON con prefijo/sufijo personalizado de forma **segura** y **totalmente offline**.  
Sencillo, rápido y amigable para desarrolladores y usuarios finales.

## Idiomas

- [English](../README.md))
- [Русский](README.ru.md)
- [简体中文](README.zh-CN.md)
- [Türkçe](README.tr.md)
- [ไทย](README.th.md)
- [Español](README.es.md)

***

## Características

- **100% offline**: sin peticiones de red, sin telemetría, sin cargas automáticas.
- **Código abierto**: base de código pequeña y enfocada, fácil de auditar.
- **Solo CPU**: binario estático único para Linux, macOS y Windows.
- **Multihilo**: utiliza todos los núcleos disponibles de la CPU para acelerar la búsqueda.
- **Aleatoriedad segura**: usa un generador de números aleatorios criptográficamente seguro.

***

## Descarga

No necesitas instalar Go ni Git.  
Puedes descargar los binarios ya compilados desde la página de GitHub Releases:

- Última versión:  
  https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases/latest

En la página de Releases, descarga el archivo correspondiente a tu sistema:

- **Windows**: `tronvanity-windows-amd64.exe`
- **macOS (Intel)**: `tronvanity-darwin-amd64`
- **macOS (Apple Silicon)**: `tronvanity-darwin-arm64`
- **Linux**: `tronvanity-linux-amd64`

Se recomienda descargar también `SHA256SUMS.txt` para verificar la integridad de los archivos descargados.

***

## Scripts de ayuda (para quienes no quieren usar la consola)

Si prefieres **no** escribir comandos manualmente en la terminal o en el Command Prompt, puedes usar los scripts de ayuda incluidos en la carpeta `scripts`.

- Explicación detallada: [`scripts/README.ru.md`](../scripts/README.ru.md)

### Resumen rápido

- **Windows**:  
  Descarga `tronvanity-windows-amd64.exe` y `windows-run.bat`, colócalos en la misma carpeta y haz doble clic en `windows-run.bat`.

- **macOS / Linux**:  
  Descarga el binario para tu sistema y `unix-run.sh`, colócalos en la misma carpeta, da permisos con `chmod +x unix-run.sh` y ejecuta `./unix-run.sh`.

Los scripts son simples “wrappers” alrededor del binario local; no acceden a Internet ni modifican archivos del sistema.

***

## Instalación para desarrolladores

```bash
# Clonar el repositorio
git clone https://github.com/bigbigha/TRON-Vanity-Address-Generator.git
cd TRON-Vanity-Address-Generator

# Compilar el binario
go build -o tronvanity ./cmd/tronvanity

# Ejemplo de ejecución
./tronvanity --prefix T888
```

***

## Inicio rápido para usuarios no técnicos

### Windows

1. Descarga `tronvanity-windows-amd64.exe`.
2. Coloca el archivo en una carpeta (o úsalo junto con `windows-run.bat`).
3. Abre Command Prompt y navega hasta esa carpeta.
4. Ejemplos:

```cmd
:: Generar una dirección que empiece por T888
tronvanity-windows-amd64.exe --prefix T888

:: Generar una dirección que termine en 9999
tronvanity-windows-amd64.exe --suffix 9999
```

### macOS

1. Descarga el binario adecuado:
    - `tronvanity-darwin-amd64` (Intel)
    - `tronvanity-darwin-arm64` (Apple Silicon)
2. Abre Terminal y navega hasta la carpeta del archivo.
3. Concede permisos de ejecución (solo la primera vez):

```bash
chmod +x tronvanity-darwin-amd64      # Intel
chmod +x tronvanity-darwin-arm64      # Apple Silicon
```

4. Ejemplos:

```bash
# Dirección que empiece por T888
./tronvanity-darwin-amd64 --prefix T888

# Dirección que termine en 9999
./tronvanity-darwin-arm64 --suffix 9999
```

### Linux

1. Descarga `tronvanity-linux-amd64`.
2. En la terminal, navega hasta la carpeta del archivo.
3. Concede permisos de ejecución:

```bash
chmod +x tronvanity-linux-amd64
```

4. Ejemplos:

```bash
# Dirección que empiece por T888
./tronvanity-linux-amd64 --prefix T888

# Dirección que termine en 9999
./tronvanity-linux-amd64 --suffix 9999
```

> Recomendación: para grandes cantidades de fondos, usa este programa en una máquina desconectada de Internet (air‑gapped) para mayor seguridad.

***

## Uso (CLI)

Ejemplos básicos:

```bash
# Coincidir solo prefijo
tronvanity --prefix T888

# Coincidir solo sufijo
tronvanity --suffix 9999

# Coincidir prefijo y sufijo, 8 workers y salir tras 3 coincidencias
tronvanity --prefix T666 --suffix 888 --workers 8 --count 3

# Generar 5 direcciones con prefijo T123
tronvanity --prefix T123 --count 5
```

### Opciones de línea de comandos

- `--prefix`: Prefijo a coincidir (por ejemplo, `T888`).
- `--suffix`: Sufijo a coincidir (por ejemplo, `9999`).
- `--workers`: Número de goroutines worker (por defecto: número de núcleos de la CPU).
- `--count`: Número de coincidencias a encontrar antes de salir (por defecto: 1).

***

## Consejos de seguridad

1. **Genera direcciones offline**: para fondos de alto valor, usa una máquina sin conexión a Internet.
2. **Verifica las direcciones**: antes de enviar fondos, asegúrate de que la dirección generada tiene el prefijo/sufijo que deseas.
3. **Protege tus claves privadas**: guárdalas en un monedero hardware, en un gestor de contraseñas o en una copia de seguridad cifrada y offline. Nunca pegues tus claves privadas en sitios web o aplicaciones de dudosa confianza.
4. **Empieza con cantidades pequeñas**: realiza pruebas con montos pequeños antes de mover grandes cantidades.

***

## Cómo funciona (resumen corto)

La herramienta realiza una búsqueda por fuerza bruta (brute‑force) de direcciones TRON vanity de forma local:

1. Genera una clave privada de 32 bytes usando el generador de números aleatorios seguro de Go.
2. Calcula la clave pública sin comprimir (65 bytes) en la curva elíptica secp256k1 (la misma que usan TRON y Ethereum).
3. Elimina el primer byte `0x04`, aplica Keccak‑256 al resto y toma los últimos 20 bytes.
4. Añade delante el prefijo de la red principal de TRON `0x41`, obteniendo 21 bytes de datos de dirección.
5. Codifica esos datos en Base58Check para obtener una dirección TRON estándar que empieza por `T...`.
6. Varios workers repiten este proceso en paralelo y comprueban si cada dirección cumple las condiciones de prefijo/sufijo. Cuando hay coincidencia, se muestran la dirección y su clave privada en la consola.

Todo el proceso se ejecuta **solo en tu máquina**: no se realizan peticiones de red y las claves privadas nunca se envían a ningún servidor.

***

## Licencia

MIT

***

## Descargo de responsabilidad

Este software se distribuye “tal cual”, sin ningún tipo de garantía.  
Lo usas **bajo tu propia responsabilidad**.  
Antes de utilizarlo con fondos reales, verifica cuidadosamente las direcciones y las claves privadas, y asegúrate de contar con buenas prácticas de seguridad y almacenamiento.
