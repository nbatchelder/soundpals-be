OUTPUT = bin/soundpals

## Build soundpals api
.PHONY: build
build:
	CGO_ENABLED=${CGO_ENABLED} go build -o ${OUTPUT} *.go

	## Run soundpals api
run:
	${OUTPUT}