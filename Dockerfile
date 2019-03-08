###
# https://hub.docker.com/r/mcareysolstice/collect
##

FROM golang:1.12-alpine AS build

RUN mkdir -p $GOPATH/src/github.com/mcarey-solstice/collect

# Do this first for faster builds if the lock changed
COPY Gopkg.lock $GOPATH/src/github.com/mcarey-solstice/collect/
WORKDIR cd $GOPATH/src/github.com/mcarey-solstice/collect && dep ensure

COPY . $GOPATH/src/github.com/mcarey-solstice/collect/
RUN go build -o /srv/collect github.com/mcarey-solstice/collect

FROM alpine:latest

RUN mkdir -p /root/bin
COPY --from=build /srv/collect /root/bin/collect

ENV PATH $PATH:/root/bin

CMD /root/bin/collect

# docker build -t mcareysolstice/collect .
