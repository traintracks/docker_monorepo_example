FROM golang:1.6
RUN apt-get update && apt-get install -y rsync
ADD . /go/src/traintracks/
WORKDIR /go/src/traintracks/coolapp
RUN go get ./...
