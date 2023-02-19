FROM golang:1.19-alpine AS builder
RUN apk add --no-cache build-base go git
WORKDIR /src/api

COPY go.mod go.sum ./
RUN go mod download -x all

COPY . ./
RUN go build -o /app/api .

# ---

FROM alpine:edge
WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /src/api/data/migrations /app/data/migrations
COPY --from=builder /app/api /app

EXPOSE 4000
CMD ["/app/api", "serve"]