###################################################################################
## Multistage docker build for creating a smallest possible docker container
###################################################################################

ARG GO_VERSION=1.17

## Stage 1
## Prepare dev environment for building service

FROM golang:${GO_VERSION}-alpine AS dev

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

ENV APP_NAME="service" \
    APP_PATH="/var/app" \
    APP_PORT=8080

ENV APP_BUILD_NAME="${APP_NAME}"

RUN echo "test"
RUN echo $APP_BUILD_NAME

COPY . ${APP_PATH}
WORKDIR ${APP_PATH}

ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOFLAGS="-mod=vendor"

EXPOSE ${APP_PORT}
ENTRYPOINT ["sh"]

## Stage 2
## Downloading required modules and building go service in separate build environment

FROM dev as build

ENV USER=serviceuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

## Build
RUN (([ ! -d "${APP_PATH}/vendor" ] && go mod download && go mod vendor) || true)
RUN go build -ldflags="-s -w" -mod vendor cmd/main.go

RUN chmod +x main

## Stage 3
## Assemble final service container from an empty scratch image

FROM scratch AS service

ENV APP_BUILD_PATH="/var/app" \
    APP_BUILD_NAME="service"

WORKDIR ${APP_BUILD_PATH}

COPY --from=build ${APP_BUILD_PATH}/main ${APP_BUILD_PATH}/service
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group

EXPOSE ${APP_PORT}

USER serviceuser:serviceuser

ENTRYPOINT ["/var/app/service"]
CMD ""
