ARG APP_DIR="/function"

### Build Stage ###

FROM golang:1.18.1-bullseye as build-stage

WORKDIR /workspace

COPY go.mod go.sum ./
COPY app/ ./

RUN go install
RUN go build -o app . && \
    cp ./app /${APP_DIR} && \
    rm -rf /workspace

### Production Stage ###

FROM alpine:latest

COPY --from=build-stage ${APP_DIR}/app ${APP_DIR}

WORKDIR /{APP_DIR}
CMD [ "/app" ]
