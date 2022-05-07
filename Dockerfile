ARG APP_DIR="/function"

### Build Stage ###

FROM golang:1.18.1-bullseye as build-stage

WORKDIR /workspace

COPY go.mod go.sum ./
COPY app/ ./

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

RUN go install
RUN go build -o app .

### Production Stage ###

FROM alpine:latest as prod-stage

ARG APP_DIR

COPY --from=build-stage /workspace/app /${APP_DIR}/app

WORKDIR ${APP_DIR}
CMD [ "./app" ]
