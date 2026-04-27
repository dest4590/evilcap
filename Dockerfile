# build stage
FROM golang:1.24-alpine AS build-env

RUN apk add --no-cache ca-certificates
RUN apk add --no-cache bash gcc g++ binutils-gold iptables wireless-tools build-base libpcap-dev libusb-dev linux-headers libnetfilter_queue-dev git

WORKDIR $GOPATH/src/github.com/dest4590/evilcap
ADD . $GOPATH/src/github.com/dest4590/evilcap
RUN make

# get caplets
RUN mkdir -p /usr/local/share/evilcap
RUN git clone https://github.com/dest4590/evilcap /usr/local/share/evilcap/caplets

# final stage
FROM alpine
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache bash iproute2 libpcap libusb-dev libnetfilter_queue wireless-tools iw
COPY --from=build-env /go/src/github.com/dest4590/evilcap/evilcap /app/
COPY --from=build-env /usr/local/share/evilcap/caplets /app/
WORKDIR /app

EXPOSE 80 443 53 5300 8080 8081 8082 8083 8000
ENTRYPOINT ["/app/evilcap"]
