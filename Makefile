.DEFAULT_GOAL := build

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -rf output/*

build: vet
	mkdir -p output/
	CGO_ENABLED=0 go build ${LDFLAGS} -o output/sensor_vl53l1x .

build-linux: vet
	mkdir -p output/
	CGO_ENABLED=1 GOOS=linux go build ${LDFLAGS} -o output/sensor_vl53l1x .