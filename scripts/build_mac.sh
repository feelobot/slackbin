docker run --rm -it -v "$PWD":/go/src/ -w /go/src/ golang:1.4.2-cross sh -c '
go get
export GOOS=darwin
export GOARCH=amd64
echo "Building $GOOS-$GOARCH"
go build -o bin/slackbin-$GOOS-$GOARCH
'
