
all: proto compile install

proto:
	protoc --go_out=. qwik.proto

compile:
	go build cmd/qwik/qwik.go
	go build cmd/qwiklog/qwiklog.go

.PHONY: build