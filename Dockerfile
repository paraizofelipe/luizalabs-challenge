FROM golang:alpine as base
RUN apk add --no-cache make cmake

FROM base as build
ADD . /opt/luizalabs-challenge
WORKDIR /opt/luizalabs-challenge
RUN make deps && make build 

FROM scratch
COPY --from=build /opt/luizalabs-challenge /opt/luizalabs/
ENTRYPOINT ["/opt/luizalabs/luizalabs-challenge"]
