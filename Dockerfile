# Base Stage
FROM golang:1.22-alpine AS base
LABEL maintainer="YOUR_NAME <YOUR_EMAIL>"
# if use private libs uncomment this
#ARG GITHUB_TOKEN
#RUN apk update && apk add ca-certificates git openssh
#RUN git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN mkdir -p dist
RUN go mod download

# Development Stage
FROM base as dev
WORKDIR /app/
COPY . .
RUN go install github.com/air-verse/air@latest
RUN go build -o dist/app cmd/main.go
CMD ["air", "-c", ".air-unix.toml", "-d"]

# Debug stage
FROM base as debug
WORKDIR /
COPY . .
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go build -gcflags="all=-N -l" -o /app/app cmd/main.go
RUN mv /go/bin/dlv /
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app/app"]

# Build Production Stage
FROM base as build
WORKDIR /app
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN $GOPATH/bin/swag init -g cmd/main.go --output docs
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o dist/app cmd/main.go

# Production Stage
FROM cgr.dev/chainguard/busybox:latest-glibc as prod
WORKDIR /app/
RUN mkdir -p /app/assets \
    mkdir -p /app/docs
COPY --from=build /app/assets/* ./assets
COPY --from=build /app/docs/* ./docs
COPY --from=build /app/dist/app .
# Specify method fetch config
CMD ["./app"]