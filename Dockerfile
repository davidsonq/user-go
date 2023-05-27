FROM golang:1.20.4 as builder
WORKDIR /app
COPY .  .
RUN go mod tidy

WORKDIR /app/cmd/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server -buildvcs=false

FROM scratch
COPY --from=builder /app/cmd/server /server
ENTRYPOINT [ "/server" ]