GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ./fsctl.Darwin.x86_64 -ldflags="-w -s" .
sudo mv ./fsctl.Darwin.x86_64 /usr/local/bin/fsctl
sudo chmod +x /usr/local/bin/fsctl