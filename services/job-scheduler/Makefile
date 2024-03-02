
.PHONY: generate build test lint build-docker compose compose-down migrate build-arm tag-and-push deploy
generate:
	go generate ./...

build: # build a server
	go build -a -o job-scheduler.bin job-scheduler/cmd/server

# Build base go container, used with `run` command to live reload binary for development
build-dev: 
	docker build --target build -f cmd/server/Dockerfile -t job-scheduler-base .
	docker run --name job-scheduler  --env-file local.env -v ${shell pwd}:/go/src/job-scheduler -v ${shell pwd}/config/local.yaml:/config/config.yaml -td job-scheduler-base
	
run:	
	docker exec -it --env-file local.env job-scheduler reflex -s go run job-scheduler/cmd/server
	

test:
	go clean -testcache
	go test ./... -v

lint:
	gofmt -l .

build-arm:
	docker build -f cmd/server/Dockerfile -t job-scheduler-arm --platform linux/amd64 .

tag-and-push: 
	docker tag job-scheduler-arm kirangershenfeld/youtube-metrics
	docker push kirangershenfeld/youtube-metrics

deploy: build-arm tag-and-push #build, tag, push build to dockerhub

