# Scripts de ayuda para TRON Vanity Address Generator

Estos scripts proporcionan una interfaz sencilla para usuarios no técnicos, permitiéndoles generar direcciones TRON vanity **sin tener que usar directamente la línea de comandos**.

## Script para Windows

### Uso

1. Descarga `tronvanity-windows-amd64.exe` desde la página de [GitHub Releases](https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases).
2. Descarga el archivo `windows-run.bat` desde la carpeta `scripts` de este repositorio.
3. Coloca ambos archivos en **la misma carpeta**.
4. Haz doble clic en `windows-run.bat` para ejecutar el script.
5. Sigue las instrucciones en pantalla e introduce:
    - El prefijo (prefix) que quieres que tenga la dirección al inicio.
    - El sufijo (suffix) que quieres que tenga la dirección al final.
    - El número de direcciones que quieres que el programa encuentre.

### Características

- Interfaz de texto simple y fácil de entender.
- No requiere conocimientos de la línea de comandos.
- Se repite automáticamente hasta que pulses `Ctrl + C` para salir.

## Script para macOS / Linux

### Uso

1. Descarga el binario adecuado para tu sistema desde la página de [GitHub Releases](https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases):

    - `tronvanity-darwin-amd64` (Intel Mac)
    - `tronvanity-darwin-arm64` (Apple Silicon Mac)
    - `tronvanity-linux-amd64` (Linux)

2. Descarga el archivo `unix-run.sh` desde la carpeta `scripts`.
3. Coloca el binario y `unix-run.sh` en **la misma carpeta**.
4. Abre una terminal y navega hasta esa carpeta.
5. Da permisos de ejecución al script: `chmod +x unix-run.sh`
6. Ejecuta el script: `./unix-run.sh`
7. Sigue las instrucciones en pantalla.

### Características

- Detecta automáticamente tu sistema operativo y arquitectura.
- Interfaz de texto sencilla.
- No requiere conocimientos avanzados de la línea de comandos.
- Hace que el binario principal sea ejecutable de forma automática.
- Se repite automáticamente hasta que pulses `Ctrl + C` para salir.

## Nota de seguridad

Estos scripts son simples “wrappers” que solo llaman al binario local. **No**:

- Acceden a Internet.
- Suben ni envían ningún tipo de datos.
- Instalan software adicional.
- Modifican archivos del sistema.

Tanto los scripts como el binario principal funcionan de forma completamente **offline** en tu propio equipo.