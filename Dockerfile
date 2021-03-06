FROM golang:1.9.2 as builder
ADD . /go/src/github.com/mattpaletta/AbilitySoftwareGroup468
WORKDIR /go/src/github.com/mattpaletta/AbilitySoftwareGroup468
RUN go get ./...
RUN go install github.com/mattpaletta/AbilitySoftwareGroup468
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags netgo -installsuffix netgo -o app .

#FROM alpine:3.7 as certs
#ENV PATH=/bin
#RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*

FROM alpine:latest
RUN apk add --no-cache bash
RUN apk add --no-cache iputils
ENV PATH=/bin
#COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY templates/test.html /root/templates/test.html
WORKDIR /root/
ADD config config
COPY --from=builder /go/src/github.com/mattpaletta/AbilitySoftwareGroup468/app .
ENTRYPOINT ["./app"]
