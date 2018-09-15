GOPATH:=$(GOPATH):$(PWD)

build:
	go build

init:
	git subtree add --prefix=ext git@github.com:enjoypi/ext.git master --squash
	git subtree add --prefix=god git@github.com:enjoypi/god.git master --squash

subtree:
	git subtree pull --prefix=ext git@github.com:enjoypi/ext.git master --squash
	git subtree pull --prefix=god git@github.com:enjoypi/god.git master --squash

push_subtree:
	git subtree push --prefix=ext git@github.com:enjoypi/ext.git master
	git subtree push --prefix=god git@github.com:enjoypi/god.git master

