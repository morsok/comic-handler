# syntax=docker/dockerfile:1
FROM rust:1.81 as rust_builder
WORKDIR /comichandler
COPY ./backend .
RUN cargo build --release

FROM node:22-alpine AS angular_builder
ARG BUILD_TYPE=production
RUN npm install -g npm
RUN npm install -g @angular/cli
COPY ./frontend /webapp
WORKDIR /webapp
RUN npm install && ng build --configuration ${BUILD_TYPE}

FROM rust:1.81-slim
COPY --from=rust_builder /comichandler/target/release/comichandler /app/comichandler
COPY --from=angular_builder /webapp/dist/frontend/* /app/static/
COPY ./backend/config /config
RUN mkdir /watch /comics
WORKDIR /app

ENTRYPOINT ["/app/comichandler"]
LABEL Name=comichandler Version=0.0.1
EXPOSE 9999
