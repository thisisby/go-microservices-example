FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o api_gateway ./cmd/main.go

RUN ls -l /app

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/api_gateway .

# Copy the configuration files
COPY --from=builder /app/pkg/config/envs/dev.env /app/pkg/config/envs/

# Verify the binary and configuration files are copied
RUN ls -l /app && ls -l /app/pkg/config/envs

EXPOSE 8080

CMD ["./api_gateway"]