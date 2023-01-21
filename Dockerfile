# syntax=docker/dockerfile:1
FROM golang:1.19-alpine AS go_builder
WORKDIR /go/src/app
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download && go mod verify
COPY ./backend .
RUN CGO_ENABLED=0 go build -o /go/bin/app -v -ldflags="-w -s" ./...

FROM node:18-alpine AS angular_builder
RUN npm install -g npm
RUN npm install -g @angular/cli
COPY ./frontend /webapp
WORKDIR /webapp
RUN npm install && ng build

FROM scratch
COPY --from=go_builder /go/bin/app /app/comichandler
COPY --from=angular_builder /webapp/dist/frontend/* /app/static/
WORKDIR /app

ENTRYPOINT ["/app/comichandler"]
LABEL Name=comichandler Version=0.0.1
EXPOSE 8080
