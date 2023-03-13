# syntax=docker/dockerfile:1
FROM golang:1.20-alpine AS go_builder
RUN apk add build-base git musl-dev
WORKDIR /go/src/app
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download && go mod verify
COPY ./backend .
RUN CGO_ENABLED=1 go build -o /go/bin/ -v -a -ldflags "-linkmode external -extldflags '-static' -s -w" ./...

FROM node:18-alpine AS angular_builder
ARG BUILD_TYPE=production
RUN npm install -g npm
RUN npm install -g @angular/cli
COPY ./frontend /webapp
WORKDIR /webapp
RUN npm install && ng build --configuration ${BUILD_TYPE}

FROM golang:1.20-alpine
ARG GIN_MODE=release
ENV GIN_MODE=${GIN_MODE}
COPY --from=go_builder /go/bin/comic-handler /app/comic-handler
COPY --from=angular_builder /webapp/dist/frontend/* /app/static/
COPY ./backend/config /config
WORKDIR /app

ENTRYPOINT ["/app/comic-handler"]
LABEL Name=comic-handler Version=0.0.1
EXPOSE 9999
