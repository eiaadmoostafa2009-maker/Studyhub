# syntax=docker/dockerfile:1

# Litestream is pinned so a future release cannot silently change production behavior.
FROM litestream/litestream:0.5.17 AS litestream

# ---------- Build stage ----------
FROM golang:1.26-bookworm AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# go-sqlite3 uses CGO, so the Debian Go image builds it normally.
ENV CGO_ENABLED=1
RUN go build -trimpath -ldflags="-s -w" -o /out/studyhub ./cmd/main.go

# ---------- Runtime stage ----------
FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /out/studyhub /app/studyhub
COPY --from=builder /src/internal/template /app/internal/template
COPY --from=builder /src/internal/static /app/internal/static
COPY --from=litestream /usr/local/bin/litestream /usr/local/bin/litestream
COPY litestream.yml /etc/litestream.yml

RUN useradd --system --uid 10001 --create-home studyhub \
    && chown -R studyhub:studyhub /app

# Persistent application data is mounted separately at /data. The VM grants
# UID 10001 ownership of that mount before starting the container.

USER studyhub

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/litestream"]
CMD ["replicate", "-config", "/etc/litestream.yml"]
