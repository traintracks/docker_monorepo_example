default: run 

builder:
	tar -zcf - ../  \
		| docker build -t coolapp-builder -f coolapp/coolapp.builder.dockerfile -

build:
	docker run -it  \
		-v ${PWD}/../:/mount \
		coolapp-builder /bin/sh -c "rsync -auiv --filter=\":- .gitignore\" /mount/ /go/src/traintracks/ && \
			go get ./... && \
			mkdir -p /mount/coolapp/output && \
			cp -v /go/bin/*  /mount/coolapp/output/"

runner: build
	docker build -t coolapp -f coolapp.dockerfile . 

run: runner
	docker run -it coolapp

all: builder run
