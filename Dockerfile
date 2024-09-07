FROM golang:1.23 AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p out
RUN CGO_ENABLED=0 go build -a -o out/app ./cmd/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=builder /workspace/out /out
COPY --from=builder /workspace/config /config

EXPOSE 8080

CMD ["/out/app"]