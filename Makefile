
# Informing Make that test-cover and build are not files,
# but rather commands that should always be run when they are called, 
# regardless of whether a file or directory with the same name exists.
.PHONY: test-cover build

test-cover:
	go tool cover -html=coverage-all.out

build:
	docker build -f build/Dockerfile -t microforge .
