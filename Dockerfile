FROM golang:1.20.4 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
WORKDIR /app/cmd
RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o server -buildvsc=false
FROM scratch
COPY --from=builder /app/cmd/server /server
ENTRYPOINT [ "/server" ]