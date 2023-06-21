HOSTNAME=hashicorp.com
NAMESPACE=QuentinN42
NAME=file
BINARY=terraform-provider-${NAME}
VERSION=0.1
OS_ARCH=linux_amd64

default: install

.PHONY: build install run
build:
	go build -ldflags="-X main.providerName=${NAME}" -o ${BINARY}

install: build
	rm -r ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/ || true
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

run: install
	rm -r examples/.terraform/ || true
	rm -r examples/.terraform.lock.hcl || true
	( cd examples && terraform init && terraform apply -auto-approve )
