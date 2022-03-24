BUILD_VERSION:=`git describe --tags`
BUILD_TIME := ${shell date "+%F%T%T%z"}
BUILD_NAME := go-blog-${shell date "+%Y%m%d%H"}
SOURCE := ./main.go
COMMIT_SHA1 := ${shell git rev-parse HEAD}


build:
	# GOOS=linux GOARCH=amd64 CGO_ENABLED=0
	go build -ldflags \
	" \
	-X 'cmd.BuildVersion=${BUILD_VERSION}' \
	" \
 	-o ${$BUILD_NAME} .

clean:
	rm ${BUILD_NAME}
