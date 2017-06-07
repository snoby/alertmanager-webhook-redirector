# Replace demo with your desired executable name
appname := spark-pivot

sources := $(wildcard *.go)

build = GOOS=$(1) GOARCH=$(2) go build -o build/$(appname)$(3)
tar = cd build && tar -cvzf $(1)_$(2).tar.gz $(appname)$(3) && rm $(appname)$(3)
zip = cd build && zip $(1)_$(2).zip $(appname)$(3) && rm $(appname)$(3)

.PHONY: all darwin linux clean

all: darwin linux

clean:
	rm -rf build

##### LINUX BUILDS #####
linux: build/darwin_amd64.tar.gz build/linux_amd64.tar.gz


build/linux_amd64.tar.gz: $(sources)
	$(call build,linux,amd64,)
	$(call tar,linux,amd64)
	cp -rf build/linux_amd64.tar.gz /Users/snoby/work/Official/spark-pivot-docker/spark-pivot.tgz


##### DARWIN (MAC) BUILDS #####
darwin: build/darwin_amd64.tar.gz

build/darwin_amd64.tar.gz: $(sources)
	$(call build,darwin,amd64,)
	$(call tar,darwin,amd64)
