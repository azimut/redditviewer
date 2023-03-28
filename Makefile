redditviewer: main.go request/main.go printer/main.go human/main.go format/main.go
	go build -v -ldflags="-s -w"
	ls -lh $@

.PHONY: clean
clean: ; $(GO) clean -x ./...

.PHONY: install
install: redditviewer
	mv redditviewer $(HOME)/go/bin/
