FROM golang:1.14.4-alpine AS builder

WORKDIR /go/app

RUN apk add --no-cache ca-certificates git make protobuf \
		&& update-ca-certificates \
		&& go get -u google.golang.org/grpc \
		&& go get -u github.com/golang/protobuf/protoc-gen-go \
		&& go get github.com/oxequa/realize 

COPY . /go/app

RUN GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o dist/api ./cmd/api

FROM busybox AS app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder dist/app /bin/app

EXPOSE 8080

CMD ["/bin/app"]
