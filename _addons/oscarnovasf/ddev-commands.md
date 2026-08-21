---
title: "oscarnovasf/ddev-commands"
github_url: "https://github.com/oscarnovasf/ddev-commands"
description: "Conjunto de comandos personalizados pensados principalmente para su uso con Drupal"
user: "oscarnovasf"
repo: "ddev-commands"
repo_id: 1151587694
default_branch: "main"
tag_name: "v1.3.6"
ddev_version_constraint: ""
dependencies: []
type: "contrib"
created_at: "2026-02-06"
updated_at: "2026-07-25"
workflow_status: "disabled"
stars: 2
---

DDEV - Herramientas para Drupal
===

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/oscarnovasf/ddev-commands/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/oscarnovasf/ddev-commands/actions/workflows/tests.yml?query=branch%3Amain)
[![version][version-badge]][changelog]
[![Licencia][license-badge]][license]
[![Código de conducta][conduct-badge]][conduct]
[![Donate][donate-badge]][donate-url]

---

## Información

Conjunto de comandos personalizados para [DDEV][ddev] que facilitan el flujo de
trabajo en proyectos [Drupal][drupal].

Este addon proporciona **varios comandos** que cubren las necesidades más
comunes del desarrollo en Drupal:
- limpieza de caché
- backup de base de datos
- ejecución de tests
- análisis estático de código
- refactorización automática
- compilación de SCSS
- traducciones
- comprobación de enlaces rotos
- métricas de código

---

## Requisitos

- [DDEV][ddev] instalado y configurado.
- Un proyecto Drupal con la estructura estándar (`/web/modules/custom`,
  `/web/themes/custom`).
- [Composer][composer] con las dependencias de desarrollo necesarias según los
  comandos que se vayan a utilizar:
  - `phpstan/phpstan` para el comando `phpstan`.
  - `phpunit/phpunit` para el comando `phpunit`.
  - `rector/rector` y `palantirnet/drupal-rector` para el comando `rector`.
  - `phpro/grumphp-shim` para el comando `grumphp`.
- Redis (opcional): el comando `cr` limpia Redis automáticamente si está
  disponible.

---

## Instalación

```bash
ddev add-on get oscarnovasf/ddev-commands
ddev restart
```

> Tras la instalación es necesario reiniciar DDEV para que se apliquen los
> cambios en el contenedor.

---

## Comandos disponibles

### `ddev backup` — Backup de base de datos

Crea un backup de la base de datos en la carpeta `_backups/` del proyecto.
Antes de exportar, ejecuta `drush cr` para volcar la caché.

```bash
ddev backup
```

El archivo generado sigue el formato `{proyecto}_{fecha}.sql.gz`:

```bash
# Resultado ejemplo
# drupal_2026.04.22_22.49.sql.gz
```

### `ddev cr` — Limpieza de caché

Ejecuta `drush cr` para reconstruir la caché de Drupal. Si el servicio Redis
está disponible, también vacía su caché de forma automática.

```bash
ddev cr
```

### `ddev phpunit` — Tests unitarios

Ejecuta [PHPUnit][phpunit] con soporte para Xdebug habilitado. Acepta como
argumento el nombre de la suite que se desea ejecutar.

```bash
# Ejecutar todas las suites
ddev phpunit

# Ejecutar una suite específica
ddev phpunit unit
```

### `ddev phpstan` — Análisis estático de código

Ejecuta [PHPStan][phpstan] para realizar análisis estático del código PHP.
Acepta opciones adicionales.

```bash
ddev phpstan
```

### `ddev rector` — Refactorización automática

Ejecuta [Rector][rector] con la configuración de
[Drupal Rector][drupal-rector] para modernizar y refactorizar el código
automáticamente.

```bash
ddev rector
```

### `ddev grumphp` — Análisis de calidad de código

Ejecuta [GrumPHP][grumphp] para comprobar la calidad del código según las
reglas configuradas en el proyecto.

```bash
ddev grumphp
```

### `ddev linkchecker` — Comprobación de enlaces rotos

Revisa todos los enlaces del sitio en busca de URLs rotas. Genera un informe
en HTML.

```bash
ddev linkchecker
```

Características:
- Comprueba enlaces externos.
- Timeout de 10 segundos por enlace.
- 30 hilos en paralelo.
- Genera un informe en `reporte-<nombre-del-sitio>.html` en la raíz del proyecto.
- Ignora enlaces `mailto:` y archivos CSS/JS.

### `ddev sass` — Compilación de SCSS

Compila los archivos SCSS de un theme custom a CSS comprimido. Detecta
automáticamente los themes que contienen una carpeta `scss/` y permite
seleccionar cuál compilar mediante un menú interactivo.

```bash
ddev sass
```

Opciones de modo:
- **Compilar** — compila una vez y termina.
- **Watch** — se queda escuchando cambios en los archivos SCSS y recompila
  automáticamente.

### `ddev translations_gen` — Traducciones

Genera archivos de traducción (.po) para módulos y temas personalizados que
tengan `interface translation` en su `.info.yml`. Detecta automáticamente los
idiomas instalados en Drupal y crea las traducciones para cada uno.

```bash
# Generar traducciones para todos los módulos y temas detectados
ddev translations_gen

# Generar traducciones para un módulo específico
ddev translations_gen -m mi_modulo

# Generar traducciones para un tema específico
ddev translations_gen -t mi_tema
```

### `ddev lineas` — Métricas de código

Cuenta las líneas de código de los módulos y temas custom del proyecto
utilizando [cloc][cloc].

```bash
ddev lineas
```

Analiza las carpetas:
- `/web/modules/custom`
- `/web/themes/custom`

## Skills para IAs

Este addon incluye skills en `assets/agents/` diseñadas para asistentes de IA
(como Cursor, Copilot, etc.). Durante la instalación se sincronizan
automáticamente en la carpeta `.agents/` del proyecto usando `rsync`, respetando
los archivos que hayas creado o modificado localmente.

#### Skills incluidas:

- **drupal-code-search**: Búsqueda de código en todos los módulos contrib de
  Drupal.org vía API de GitLab. Requiere un token de acceso personal
  (`DRUPALORG_GITLAB_TOKEN`).
- **drupal-migration**: Migraciones D7→D10/D11 con Migrate API, CSV, JSON API
  y plugins personalizados de source/process.
- **drupal-security**: Prevención de vulnerabilidades en Drupal (XSS, SQL
  injection, acceso no autorizado) con patrones seguros y checklist de revisión.
- **git-mr-description**: Genera descripciones de Merge Request para GitLab
  analizando el diff de la rama contra `develop`.
- **redmine-task-creator**: Crea definiciones detalladas de tareas Redmine en
  formato Textile con estimaciones, checklist y contexto del proyecto.

---

## Assets adicionales

### GitHub — Formato de commits

El addon incluye en `assets/github/instructions/` una plantilla de
instrucciones para generar commits con formato
[conventional commits](https://www.conventionalcommits.org/):

```
<type>: #<id> <descripción>
```

Durante la instalación se sincroniza a `.github/instructions/` del proyecto.

**Para usarlo desde VS Code**:

Debes añadir a tu `settings.json` esta configuración:

```json
"github.copilot.chat.commitMessageGeneration.instructions": [
    {
        "file": ".github/instructions/git-commit.instructions.md"
    }
],
```

---

## Estructura del proyecto

```
ddev-commands/
├── assets/
│   ├── agents/
│   │   └── skills/          # Skills para asistentes de IA
│   └── github/
│       └── instructions/    # Instrucciones para vscode
├── commands/
│   ├── host/                # Comandos para ejecutar en el anfitrión
│   └── web/                 # Comandos para ejecutar en el contenedor web
├── web-build/
│   └── Dockerfile.ddev-drupal-tools
├── config.drupal-tools.yaml
└── install.yaml
```

---

## Desinstalación

```bash
ddev add-on remove ddev-commands
ddev restart
```

---
⌨️ con ❤️ por [Óscar Novás][mi-web] 😊

[mi-web]: https://oscarnovas.com "for developers"

[version]: v1.3.6
[version-badge]: https://img.shields.io/badge/Versión-1.3.6-blue.svg

[license]: .github/LICENSE.md
[license-badge]: https://img.shields.io/badge/Licencia-GPLv3+-green.svg "Leer la licencia"

[conduct]: .github/CODE_OF_CONDUCT.md
[conduct-badge]: https://img.shields.io/badge/C%C3%B3digo%20de%20Conducta-2.0-4baaaa.svg "Código de conducta"

[changelog]: CHANGELOG.md "Histórico de cambios"

[donate-badge]: https://img.shields.io/badge/Donaci%C3%B3n-PayPal-red.svg
[donate-url]: https://paypal.me/oscarnovasf "Haz una donación"

[ddev]: https://ddev.readthedocs.io/
[drupal]: https://www.drupal.org/
[composer]: https://getcomposer.org/
[phpunit]: https://phpunit.de/
[phpstan]: https://phpstan.org/
[rector]: https://getrector.com/
[drupal-rector]: https://github.com/palantirnet/drupal-rector
[grumphp]: https://github.com/phpro/grumphp-shim
[cloc]: https://github.com/AlDanial/cloc
