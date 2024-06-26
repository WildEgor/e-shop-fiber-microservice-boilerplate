# Project boilerplate: Go + Fiber + Docker + Air
A containerised eShopFiberMicroserviceBoilerplate (mmm Java-style naming)

| Component        | Link                                   |
|------------------|----------------------------------------|
| Language         | [Go](https://go.dev/)                  |
| Framework        | [Fiber](https://docs.gofiber.io/)      |
| Hot Reloading    | [Air](https://github.com/cosmtrek/air) |
| Containerisation | [Docker](https://www.docker.com/)      |
| Profiler         | [Pyroscope](https://pyroscope.io/)      |

## Features
- Hot-reloading in dev mode;
- Containerized debug mode;
- Profiler in debug mode;
- Health checks;
- Configuration;

## Structure

- [/api](api/README.md)
- [/assets](assets/README.md)
- [/cmd](cmd/README.md)
- [/internal/adapters](internal/adapters/README.md)
- [/internal/configs](internal/configs/README.md)
- [/internal/handlers](internal/handlers/README.md)
- [/internal/mappers](internal/mappers/README.md)
- [/internal/middlewares](internal/middlewares/README.md)
- [/internal/models](internal/models/README.md)
- [/internal/repositories](internal/repositories/README.md)
- [/internal/routers](internal/routers/README.md)

## Quick start

Install the dependencies:

> Note: this project uses [Go mod](https://blog.golang.org/using-go-modules), the official module manager, to handle Go modules in a portable way without having to worry about GOPATH.

```bash
go mod download
go mod vendor
go mod verify
```

Define configs using `yaml`:

```bash
cp config.example.yaml config.yaml
```

Run locally:

> Note: this builds the Docker image and runs it automatically with the config defined in `docker-compose.yaml`. This saves you having to build the docker image and then run a manual `docker run` command with all the flags (for environment variables, ports, etc).

Local:
```bash
task wire && air
```

Using Docker (change `target` to `dev` in `.env`):
```bash
docker compose up --build -d app
```

## Debug

> Note: in debug mode available port for break points and profiler

Using Docker (change `target` to `debug` in `.env`):
```bash
docker compose up --build -d app
```

## Production

> Note: Environment variables are never baked into the image, or they wouldn't be _environment_ variables. The production environment will start a Docker container based on this image, but it will have to pass the environment variables to the container when it runs it.

Using Docker (change `target` to `prod` in `.env`):
```bash
docker compose up --build --remove-orphans app
```