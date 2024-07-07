doc:
	go install golang.org/x/tools/cmd/godoc@latest
	godoc -http=:6060
