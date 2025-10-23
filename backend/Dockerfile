# Build stage
FROM golang:1.25 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Module is "dia-manager-backend", entrypoint is main.go in this folder
RUN GIN_MODE=release CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/server ./.

# Runtime stage
FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=build /out/server /app/server
# Defaults (overridden by compose env)
ENV PORT=8369
EXPOSE 8369
USER 65532:65532
ENTRYPOINT ["/app/server"]