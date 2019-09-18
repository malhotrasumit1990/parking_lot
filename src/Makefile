# defining environment variables
TARGETNAME=src
REPONAME=parking_lot
TARGETPATH=$(GOPATH)/bin/$(TARGETNAME)
SERVICEPACKNAME=$(REPONAME).zip
BUILDNUMBER = $(BUILD_NUMBER)
ifndef $(BUILDNUMBER)
	BUILDNUMBER=`git rev-parse HEAD`
endif
BUILDCOMMITSHA = `git rev-parse HEAD`

LDFLAGBUILDVERSION = "-X github.com/parking_lot/src/web.BuildCommitSHA=$(BUILDCOMMITSHA)\
			 -X github.com/ContinuumLLC/platform-itsm-sync-service/src/web.BuildNumber=$(BUILDNUMBER)"

# global command
.PHONY: all
#all: clean dependencies platform-itsm-sync-service  pack-service
all: clean dependencies build cover

.PHONY: local
local: clean local-dep build cover pack-service

.PHONY: clean
clean:
	@echo "\n*********Cleaning in progress*************"
	rm -rf cover-all.out
	rm -rf cover-all.html
	rm -rf $(SERVICEPACKPATH)
	rm -rf $(SERVICEPACKPATH)/*.zip

.PHONY: local-dep
local-dep:
	@echo "\n*********Installing dependencies on local*************"
	glide install


.PHONY: dependencies
dependencies:
	@echo "\n*********Installing dependencies*************"
	rm -f glide.lock
	glide cc
	glide install


.PHONY: build
build:
	@echo "\n*************Generating Binary*************"
	GOOS=linux GOARCH=amd64 go build -ldflags ${LDFLAGBUILDVERSION} -o $(GOPATH)/bin/linux_amd64/$(TARGETNAME)

#Target to run test suits
.PHONY: test
test:
	@echo "\n*************Executing testcases*************"
	go test -coverprofile ./businesslayer/cover.out -covermode=count ./businesslayer


.PHONY: cover
cover: test
	@echo "\n*************Executing testcases and store coverage in cover-all.out*************"
	@echo "mode: count" >  cover-all.out
	tail -n +2 ./businesslayer/cover.out >> cover-all.out

	
	gocover-cobertura < cover-all.out > cover-cobertura.xml
