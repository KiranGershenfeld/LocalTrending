
.PHONY: generate build test lint build-docker compose compose-down migrate build-arm tag-and-push deploy
generate:
	go generate ./...

# Build binary
build: 
	go build -a -o video-ingest.bin video-ingest/cmd/server

# Build base go container, used with `run` command to live reload binary for development
build-dev: 
	docker build --target build -f cmd/server/Dockerfile -t local-trending-base .
	docker run --name local-trending  --env-file local.env -v ${shell pwd}:/go/src/local-trending -v ${shell pwd}/migrations:/config/migrations -v ${shell pwd}/config/local.yaml:/config/config.yaml -td local-trending-base
	
run:	
	docker exec -it --env-file local.env local-trending reflex -s go run video-ingest/cmd/server
	
test:
	go clean -testcache
	go test ./... -v

lint:
	gofmt -l .
	
# Build for amd64, tag, push build to dockerhub
deploy:
	docker build -f cmd/server/Dockerfile -t local-trending-amd --platform linux/amd64 .
	docker tag local-trending-amd kirangershenfeld/youtube-metrics
	docker push kirangershenfeld/youtube-metrics

migrate:
	docker run --rm -v migrations:/migrations --network host migrate/migrate -path=/migrations/ \
	-database mysql://root:password@localhost:3306/local_db?charset=utf8&parseTime=True&multiStatements=true up 2

