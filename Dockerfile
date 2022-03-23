FROM golang:1.18 AS builder
WORKDIR /app
COPY spinner.go ./
RUN CGO_ENABLED=0 go build -o spinner spinner.go

FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/spinner ./
ENTRYPOINT ["/app/spinner"]
