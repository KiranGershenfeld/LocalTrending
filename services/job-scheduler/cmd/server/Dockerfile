FROM golang:1.21.6-alpine AS build

RUN mkdir -p /go/src/job-scheduler ~/.ssh && \
    apk add --no-cache git openssh-client make gcc libc-dev
WORKDIR /go/src/job-scheduler
COPY . .
RUN make build

FROM alpine:3
COPY --from=build /go/src/job-scheduler/job-scheduler /bin/job-scheduler

CMD /bin/job-scheduler