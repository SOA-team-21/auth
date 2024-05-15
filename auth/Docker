FROM golang:alpine AS auth-builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o auth-webapp

FROM alpine
COPY --from=auth-builder /app/auth-webapp /usr/bin/auth-webapp
EXPOSE 90
ENTRYPOINT ["/usr/bin/auth-webapp"]
