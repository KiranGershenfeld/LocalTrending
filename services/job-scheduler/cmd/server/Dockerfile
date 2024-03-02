#Build stage
FROM golang:1.21.6-alpine AS build

RUN mkdir -p /go/src/job-scheduler ~/.ssh && \
    apk add --no-cache git openssh-client make gcc libc-dev python3-dev py3-pip py3-virtualenv libpq-dev && \
    go install github.com/cespare/reflex@latest

WORKDIR /go/src/job-scheduler
COPY go.mod go.sum ./
COPY go.mod go.sum ./
RUN go mod download -x

COPY jobs ./jobs
COPY jobs_common ./jobs_common

RUN python -m venv .venv
RUN source .venv/bin/activate
RUN python -m pip install -r jobs/requirements.txt --break-system-packages
RUN pip install -e jobs_common/ --break-system-packages

#Compile stage
FROM build as binary
WORKDIR /go/src/job-scheduler
COPY . .
RUN make build

#Package stage
FROM alpine:3 as deploy
RUN apk --no-cache add build-base gcc python3-dev py3-pip py3-virtualenv libpq-dev
COPY --from=binary /go/src/job-scheduler/job-scheduler.bin /bin/job-scheduler
COPY --from=binary /go/src/job-scheduler/jobs /jobs
COPY --from=binary /go/src/job-scheduler/jobs_common /jobs_common

RUN python -m venv .venv
RUN source .venv/bin/activate
RUN python -m pip install -r jobs/requirements.txt --break-system-packages
RUN pip install -e jobs_common/ --break-system-packages

CMD /bin/job-scheduler