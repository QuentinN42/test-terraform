HOSTNAME=hashicorp.com
NAMESPACE=QuentinN42
NAME=provider
BINARY=terraform-provider-${NAME}
VERSION=0.3.1
OS_ARCH=linux_amd64

default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

run:
	( cd examples && terraform init && terraform apply -auto-approve )
