API_FILE=apiid.txt
CLIENT_FILE=clientid.txt

API_ID=`cat $(API_FILE)`
CLIENT_ID=`cat $(CLIENT_FILE)`

compile:
	echo "Building application binaries"
	CGO_ENABLED=0 go build -o bin/outdoor-api cmd/api/main.go
	CGO_ENABLED=0 go build -o bin/go-outdoors cmd/client/main.go

build:
	echo "Building docker images"
	docker build -f client/Dockerfile --iidfile clientid.txt -t go-outdoors:latest .
	docker build -f api/Dockerfile --iidfile apiid.txt -t outdoor-api:latest .

deploy:
	# docker login --username=_ --password=$(heroku auth:token) registry.heroku.com
	heroku container:login

	echo "Docker CLIENT Image ID is $(CLIENT_ID)"
	echo "Docker API Image ID is $(API_ID)"

	docker tag $(API_ID) registry.heroku.com/outdoor-api/web
	docker push registry.heroku.com/outdoor-api/web

	docker tag $(CLIENT_ID) registry.heroku.com/go-outdoors/web
	docker push registry.heroku.com/go-outdoors/web

	heroku login

	curl --netrc -X PATCH https://api.heroku.com/apps/outdoor-api/formation -H "Content-Type: application/json" -H "Accept: application/vnd.heroku+json; version=3.docker-releases" --data '{ "updates": [ { "type": "web", "docker_image": "'$(API_ID)'" } ] }'

	curl --netrc -X PATCH https://api.heroku.com/apps/go-outdoors/formation --header "Content-Type: application/json" --header "Accept: application/vnd.heroku+json; version=3.docker-releases" --data '{ "updates": [ { "type": "web", "docker_image": "'$(CLIENT_ID)'" } ] }'