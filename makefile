doc:
	go install golang.org/x/tools/cmd/godoc@latest
	godoc -http=:6060

install:
	go build
	go install