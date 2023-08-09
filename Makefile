default: install

generate:
	go generate ./...

install:
	# FML
	mkdir -p ~/.terraform.d/plugins/networknext.com/networknext/networknext/5.0.1/darwin_amd64
	GOBIN=~/.terraform.d/plugins/networknext.com/networknext/networknext/5.0.1/darwin_amd64 go install .

test:
	go test -count=1 -parallel=4 ./...

testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...
