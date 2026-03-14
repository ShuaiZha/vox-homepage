# AdaVox Homepage

A beautiful, lightweight homepage for your server, NAS, or browser.

## Features

- 🎨 Clean and modern UI
- 📱 Responsive design
- 🔧 Easy configuration
- 🚀 Lightweight (78MB)
- 🐳 Docker support

## Quick Start

### Docker Run

```bash
docker run -d \
  --name vox-homepage \
  -p 3002:3002 \
  -v ./conf:/app/conf \
  -v ./uploads:/app/uploads \
  -v ./database:/app/database \
  yexingmiren/vox-homepage:latest
```

### Docker Compose

```yaml
version: "3.2"

services:
  vox-homepage:
    image: yexingmiren/vox-homepage:latest
    container_name: vox-homepage
    volumes:
    - ./conf:/app/conf
    - ./uploads:/app/uploads
    - ./database:/app/database
    ports:
    - 3002:3002
    restart: always
```

```bash
docker-compose up -d
```

## Default Credentials

- Username: `admin`
- Password: `adminadmin`

## Access

- Web UI: http://localhost:3002
- Default login: http://localhost:3002/#/login

## Build from Source

```bash
# Build Docker image
docker build -t yexingmiren/vox-homepage:latest .
```

## Docker Hub

Image available at: https://hub.docker.com/r/yexingmiren/vox-homepage
