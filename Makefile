
all: dep install

proto:
	protoc --go_out=. qwik.proto

dep:
	dep ensure

compile:
	go build cmd/qwik/qwik.go
	go build cmd/qwiklog/qwiklog.go

install:
	go install cmd/qwik/qwik.go
	go install cmd/qwiklog/qwiklog.go

.PHONY: build