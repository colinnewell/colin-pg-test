FROM golang:1.17 AS build

COPY go.* /pg-test/
WORKDIR /pg-test
# allow things to be cached.
RUN go mod download

COPY pgx.go /pg-test/
RUN CGO_ENABLED=0 GOOS=linux go build -o pg-test

FROM scratch

COPY --from=build /pg-test/pg-test /pg-test

CMD ["/pg-test"]
