BUILD_VERSION	:=`git describe --tags --dirty="-dev"`
BUILD_TIME 		:= ${shell date "+%F%T%T%z"}
BUILD_NAME 		:= go-blog-`go env GOOS`-`go env GOARCH`-${BUILD_VERSION}
COMMIT_SHA1 	:= ${shell git rev-parse HEAD}
GIT_VERSION		:=`git --version`
GO_VERSION 		:=`go env GOVERSION`
PLATFORM 		:="`go env GOOS`/`go env GOARCH`"
SERVER_NAME 	:= go-blog-server-`go env GOOS`-`go env GOARCH`-${BUILD_VERSION}
CLIENT_NAME 	:= go-blog-cli-`go env GOOS`-`go env GOARCH`-${BUILD_VERSION}
BUILD_DIR		:= bin




server:
	GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/${SERVER_NAME} -ldflags  \
	" \
	-X 'blog/cmd.BuildVersion=${BUILD_VERSION}' \
	-X 'blog/cmd.GoVersion=${GO_VERSION}' \
	-X 'blog/cmd.GitVersion=${GIT_VERSION}' \
	-X 'blog/cmd.GitCommit=${COMMIT_SHA1}' \
	-X 'blog/cmd.Compiler=${BUILD_VERSION}' \
	-X 'blog/cmd.Platform=${PLATFORM}' \
	-X 'blog/cmd.ServerName=${SERVER_NAME}' \
    -X 'blog/cmd.ClientName=${CLIENT_NAME}' \
	" \
	main.go

client:server
	# GOOS=darwin GOARCH=amd64 CGO_ENABLED=0
	GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/${CLIENT_NAME} -ldflags  \
	" \
	-X 'blog/cmd.BuildVersion=${BUILD_VERSION}' \
	-X 'blog/cmd.GoVersion=${GO_VERSION}' \
	-X 'blog/cmd.GitVersion=${GIT_VERSION}' \
	-X 'blog/cmd.GitCommit=${COMMIT_SHA1}' \
	-X 'blog/cmd.Compiler=${BUILD_VERSION}' \
	-X 'blog/cmd.Platform=${PLATFORM}' \
	-X 'blog/cmd.ServerName=${SERVER_NAME}' \
	-X 'blog/cmd.ClientName=${CLIENT_NAME}' \
	" \
	cli.go


all:client

clean:
	rm ${BUILD_DIR}/${CLIENT_NAME} ${BUILD_DIR}/${SERVER_NAME}
