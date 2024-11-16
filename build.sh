#mac
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ./fsctl.Darwin.x86_64 -ldflags="-w -s" .
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o ./fsctl.Darwin.arm64 -ldflags="-w -s" .
sudo \cp ./fsctl.Darwin.arm64 /usr/local/bin/fsctl

#linux64
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./fsctl.Linux.x86_64 -ldflags="-w -s" .
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o ./fsctl.Linux.i686 -ldflags="-w -s" .
#windows
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ./fsctl.Windows.x86_64 -ldflags="-w -s" .
GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o ./fsctl.Windows.i686 -ldflags="-w -s" .
