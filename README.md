# Comic Handler [![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/)

Automated Comic Book downloader greatly inspired by <https://github.com/mylar3/mylar3>

Currently, it is an exercise for me to learn Rust and Angular.

## How to develop

### Prerequisites

- [Node.js](https://nodejs.org/en/download/)
- [Rust](https://www.rust-lang.org/tools/install)

### Frontend

Make sure you have the Angular CLI installed and the dependencies installed

```shell
npm install -g @angular/cli
npm install --include=dev
```

#### Running the angular frontend

On a separate terminal run the backend

```shell
cd backend
cargo run
```

Then run the frontend

```shell
npm watch
```

### Backend

#### Running the backend for development

Install [cargo-watch](https://github.com/watchexec/cargo-watch) and use the following command

```shell
cargo watch -x clippy -x run
```
