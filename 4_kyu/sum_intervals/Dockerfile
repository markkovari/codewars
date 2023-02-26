FROM rust:1.67 as builder
WORKDIR /usr/src/sum_intervals
COPY . .
RUN cargo install --path .

FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y extra-runtime-dependencies && rm -rf /var/lib/apt/lists/*
COPY --from=builder /usr/local/cargo/bin/sum_intervals /usr/local/bin/sum_intervals
CMD ["sum_intervals"]