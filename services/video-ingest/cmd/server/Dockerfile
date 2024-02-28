#Build stage
FROM golang:1.21.6-alpine AS build

RUN mkdir -p /go/src/local-trending ~/.ssh && \
    apk add --no-cache git openssh-client make gcc libc-dev && \
    go install github.com/cespare/reflex@latest

WORKDIR /go/src/local-trending
COPY go.mod go.sum ./
RUN go mod download -x

#Compile stage
FROM build as binary
COPY . .
RUN make build

#Package stage
FROM binary as pkg
COPY --from=binary /go/src/local-trending/video-ingest /bin/video-ingest

CMD /bin/video-ingest