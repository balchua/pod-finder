FROM quay.io/bitnami/golang:1.14 as build-env

WORKDIR /go/src/app

COPY go.* ./
COPY main.go ./
COPY cmd/* cmd/
COPY handler/* handler/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix nocgo -o /usr/local/bin/pod-finder .

FROM gcr.io/distroless/base:debug
COPY --from=build-env /usr/local/bin/pod-finder /
CMD ["/pod-finder"]
