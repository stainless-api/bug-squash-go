make:
	go build ./...

.PHONY: test
test:
	./scripts/test.sh
