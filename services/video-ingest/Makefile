
.PHONY: generate build test lint build-docker compose compose-down migrate build-arm tag-and-push deploy
generate:
	go generate ./...

# Build binary
build: 
	go build -a -o video-ingest.bin video-ingest/cmd/server

# Build base go container, used with `run` command to live reload binary for development
build-dev: 
	docker build --target build -f cmd/server/Dockerfile -t video-ingest-base .
	docker run --name video-ingest  --env-file local.env -v ${shell pwd}:/go/src/video-ingest -v ${shell pwd}/migrations:/config/migrations -v ${shell pwd}/config/local.yaml:/config/config.yaml -td video-ingest-base
	
run:	
	docker exec -it --env-file local.env video-ingest reflex -s go run video-ingest/cmd/server
	
test:
	go clean -testcache
	go test ./... -v

lint:
	gofmt -l .
	
# Build for amd64, tag, push build to dockerhub
deploy:
	docker build -f cmd/server/Dockerfile -t video-ingest-amd --platform linux/amd64 .
	docker tag video-ingest-amd kirangershenfeld/youtube-metrics:video-ingest
	docker push kirangershenfeld/youtube-metrics:video-ingest

migrate:
	docker run --rm -v migrations:/migrations --network host migrate/migrate -path=/migrations/ \
	-database mysql://root:password@localhost:3306/local_db?charset=utf8&parseTime=True&multiStatements=true up 2

