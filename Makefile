.PHONY: build release clean


build-all:
	@$(MAKE) clean
	@$(MAKE) build.linux
	@$(MAKE) build.darwin

release:
	@$(MAKE) build-all
	tar xzf filebeat-helper.tar.gz dist

clean:
	rm -rf dist

build.linux:
	mkdir -p dist
	GOOS=linux go build -o dist/filebeat-helper-linux -i .

build.darwin:
	mkdir -p dist
	GOOS=darwin go build -o dist/filebeat-helper-darwin -i .
