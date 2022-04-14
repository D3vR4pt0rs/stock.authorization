FROM golang:latest AS builder
ENV PROJECT_PATH=/app/auth_server
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY . ${PROJECT_PATH}
WORKDIR ${PROJECT_PATH}
RUN go build cmd/auth_server/main.go

FROM golang:alpine
WORKDIR /app/cmd/auth_server
COPY --from=builder /app/auth_server/main .
EXPOSE 5000
CMD ["./main"]