#mac
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ./fsctl.darwin.amd64 -ldflags="-w -s" .
#linux64
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./fsctl.linux.amd64 -ldflags="-w -s" .
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o ./fsctl.linux.386 -ldflags="-w -s" .
#windows
#GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ./fsctl.windows.amd64 -ldflags="-w -s" .
#GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o ./fsctl.windows.386 -ldflags="-w -s" .