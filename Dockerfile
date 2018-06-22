# https://docs.docker.com/develop/develop-images/multistage-build/
FROM golang:1.10
WORKDIR /go/src/github.com/asicsdigital/dudewheresmy
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=0 /go/src/github.com/asicsdigital/dudewheresmy/app ./dudewheresmy
ENTRYPOINT ["/usr/bin/dudewheresmy"]
CMD ["--help"]
