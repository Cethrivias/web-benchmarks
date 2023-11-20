FROM rust:1.71

WORKDIR /usr/src/myapp
COPY . .

RUN cargo install --path .

CMD ["rust-rocket"]
