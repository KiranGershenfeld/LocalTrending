
.PHONY: generate build test lint build-docker compose compose-down migrate build-arm tag-and-push deploy
generate:
	go generate ./...

build: # build a server
	go build -a -o video-ingest.bin video-ingest/cmd/server

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

build-arm:
	docker build -f cmd/server/Dockerfile -t local-trending-arm --platform linux/amd64 .

tag-and-push: 
	docker tag local-trending-arm kirangershenfeld/youtube-metrics
	docker push kirangershenfeld/youtube-metrics

deploy: build-arm tag-and-push #build, tag, push build to dockerhub
	
compose.%:
	$(eval CMD = ${subst compose.,,$(@)})
	tools/script/compose.sh $(CMD)

migrate:
	docker run --rm -v migrations:/migrations --network host migrate/migrate -path=/migrations/ \
	-database mysql://root:password@localhost:3306/local_db?charset=utf8&parseTime=True&multiStatements=true up 2

