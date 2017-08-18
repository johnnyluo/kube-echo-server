build:
	GOOS=linux go build -o echo-server-linux
	docker build -t echo-server .
	rm -rf echo-server-linux

push:
	docker tag echo-server:latest johnnyluo/echo-server:latest
	docker push johnnyluo/echo-server:latest