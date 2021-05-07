all: install run startSW
	
install:
	go install ${GOPATH}/src/github.com/ueihvn/learnWebService/chapter1/mirrorFinder
run:
	${GOBIN}/mirrorFinder &
startSW:
	docker run --rm -d -p 80:8080 --name sw_mirrorFinder -e SWAGGER_JSON=/app/openapi.json -v ${GOPATH}/src/github.com/ueihvn/learnWebService/chapter1:/app swaggerapi/swagger-ui:v3.48.0
clean:
	pkill mirrorFinder
	docker kil sw_mirrorFinder
