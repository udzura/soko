VERSION := $(shell go run cmd/soko/soko.go version | sed 's/version //')
.PHONY: solo test setup clean-zip all compress release

solo: test
	go build ./cmd/soko

test:
	go test ./...

setup:
	which gox || go get github.com/mitchellh/gox
	which ghr || go get github.com/tcnksm/ghr

clean-zip:
	find pkg -name '*.zip' | xargs rm

all: setup test
	gox \
	    -os="darwin linux windows" \
	    -arch="amd64" \
	    -output "pkg/{{.Dir}}_$(VERSION)-{{.OS}}-{{.Arch}}" \
	    ./cmd/soko

compress: all clean-zip
	cd pkg && ( find . -perm -u+x -type f -name 'soko*' | gxargs -i zip -m {}.zip {} )

release: compress
	git push origin master
	ghr $(VERSION) pkg
	git fetch origin --tags

