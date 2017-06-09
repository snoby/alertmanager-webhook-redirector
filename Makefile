# Replace demo with your desired executable name
appname := spark-pivot

sources := $(wildcard *.go)

built = GOOS=$(1) GOARCH=$(2) go build -o built/$(appname)$(3)
tar = cd built && tar -cvzf $(1)_$(2).tar.gz $(appname)$(3) && rm $(appname)$(3)
zip = cd built && zip $(1)_$(2).zip $(appname)$(3) && rm $(appname)$(3)

.PHONY: all darwin linux clean

all: darwin linux

clean:
	rm -rf built

build: spark-pivot
	go build -v

##### LINUX BUILDS #####
linux: built/darwin_amd64.tar.gz built/linux_amd64.tar.gz


built/linux_amd64.tar.gz: $(sources)
	$(call built,linux,amd64,)
	$(call tar,linux,amd64)


##### DARWIN (MAC) BUILDS #####
darwin: built/darwin_amd64.tar.gz

built/darwin_amd64.tar.gz: $(sources)
	$(call built,darwin,amd64,)
	$(call tar,darwin,amd64)


