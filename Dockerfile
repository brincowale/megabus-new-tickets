FROM golang:1.14-buster as build

WORKDIR /go/src/megabus-new-tickets
ADD . /go/src/megabus-new-tickets
RUN go get -d -v ./...
RUN go build -ldflags "-s -w" -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
WORKDIR /
CMD ["/app"]