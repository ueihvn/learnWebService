# Find fatest mirror site from a list (debian)
    ##mirrorFinder/main.go
    function  findFatest(): make request to all mirrors and calculate the fatest. Using goroutines to request URLs
    When 1st goroutine send data to channels(urlChan,latencyChan), two channels receiving data and return response struct, others goroutines block
    ##mirrors/data.go
    list URL mirror link in array of string
    ## openapi.json
    using SwaggerUI, an API desciption formar for REST APIs
    ##Makefile
        all: to run HTTP server and SwaggerUI using Docker
        install: complies packages mirrors and place copy in the $GOPATH/pkg.Places a binary in $GOPATH/bin
        run: run binary in bg
        startSW: run docker container SwaggerUI
            -v mount openapi.json to container
        test: curl to test Server
            -i: include protocol header in output
            -X: Specify request command to use (GET)
        clean: terminate server and SwaggerUI container
