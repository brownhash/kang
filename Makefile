build:
	go get
	go build -o bin/kang .
buildBasic:
	go get
	env GOOS=darwin GOARCH=amd64 go build -o bin/kang_darwin_amd64
	env GOOS=darwin GOARCH=arm64 go build -o bin/kang_darwin_arm64
	env GOOS=linux GOARCH=amd64 go build -o bin/kang_linux_amd64
	env GOOS=linux GOARCH=arm go build -o bin/kang_linux_arm
	env GOOS=linux GOARCH=arm64 go build -o bin/kang_linux_arm64