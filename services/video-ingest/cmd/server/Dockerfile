#Build stage
FROM golang:1.21.6-alpine AS build

RUN mkdir -p /go/src/video-ingest ~/.ssh && \
    apk add --no-cache git openssh-client make gcc libc-dev && \
    go install github.com/cespare/reflex@latest

WORKDIR /go/src/video-ingest
COPY go.mod go.sum ./
RUN go mod download -x

#Compile stage
FROM build as binary
WORKDIR /go/src/video-ingest
COPY . .
RUN make build

#Package stage
FROM alpine:3
COPY --from=binary /go/src/video-ingestg/video-ingest.bin /bin/video-ingest

CMD /bin/video-ingest