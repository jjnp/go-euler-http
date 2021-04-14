FROM golang
WORKDIR /go/src/github.com/jjnp/go-euler-http
COPY . .
RUN go get ./... && CGO_ENABLED=0 GOOS=linux go build -o app
FROM alpine
WORKDIR /app/
COPY --from=0 /go/src/github.com/jjnp/go-euler-http/app .
CMD ["./app"]

