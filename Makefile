
# code style recommendations, to enforce code conventions
lint:
	golint ./...

# to find subtle issues where code might not work as intended
vet:
	go vet ./...

# standard formatting
fmt:
	go fmt ./...