#!/bin/bash
GITHUB_USERNAME=lucasandre23
GITHUB_EMAIL=lucasandre152015@gmail.com

SERVICE_NAME=payment
RELEASE_VERSION=v1.2.3

echo "Instalando/Atualizando plugins gRPC..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH="$PATH:$(go env GOPATH)/bin"


echo "Gerando código fonte Go para: ${SERVICE_NAME}"

mkdir -p golang/${SERVICE_NAME}

protoc --proto_path=. \
  --go_out=./golang/${SERVICE_NAME} \
  --go_opt=paths=source_relative \
  --go-grpc_out=./golang/${SERVICE_NAME} \
  --go-grpc_opt=paths=source_relative \
  ./${SERVICE_NAME}.proto

echo "Arquivos gerados:"
ls -al ./golang/${SERVICE_NAME}

echo "Configurando Go Module..."
cd golang/${SERVICE_NAME}

rm -f go.mod 
go mod init github.com/${GITHUB_USERNAME}/microservices-proto/golang/${SERVICE_NAME}
go mod tidy

echo "Script finalizado com sucesso."