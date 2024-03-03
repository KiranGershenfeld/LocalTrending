
.PHONY: build build-dev run deploy

build: 
	docker build -f Dockerfile -t job-scheduler-base .
	
run:	
	docker run --name job-scheduler --env-file local.env -v ${shell pwd}/src:/src -td job-scheduler-base
	
deploy: 
	docker build  -f Dockerfile -t job-scheduler-amd --platform linux/amd64 .
	docker tag job-scheduler-amd kirangershenfeld/youtube-metrics:job-scheduler
	docker push kirangershenfeld/youtube-metrics:job-scheduler
