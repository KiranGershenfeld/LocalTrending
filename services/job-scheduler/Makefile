
.PHONY: generate build test lint build-docker compose compose-down migrate build-arm tag-and-push deploy
generate:
	go generate ./...

build: # build a server
	go build -a -o job-scheduler.bin job-scheduler/cmd/server

test:
	go clean -testcache
	go test ./... -v

lint:
	gofmt -l .

build-docker: # build docker image
	docker build -f cmd/server/Dockerfile -t job-scheduler .

build-arm:
	docker build -f cmd/server/Dockerfile -t job-scheduler-arm --platform linux/amd64 .

tag-and-push: 
	docker tag job-scheduler-arm kirangershenfeld/youtube-metrics
	docker push kirangershenfeld/youtube-metrics

deploy: build-arm tag-and-push #build, tag, push build to dockerhub
	
run:
	docker run --name job-scheduler job-scheduler


