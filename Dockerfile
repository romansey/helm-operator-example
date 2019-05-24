FROM golang:1.12-alpine as builder
WORKDIR /app
COPY ./cmd/ ./cmd/
RUN CGO_ENABLED=0 GOOS=linux go build cmd/examplesrv.go

FROM scratch 
WORKDIR /app
COPY --from=builder /app/examplesrv .
USER 1000
CMD ["./examplesrv"]