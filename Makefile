VERSION=$(shell git describe --tags)
BUILD_COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
PROJECTNAME := $(shell basename $(shell pwd))

# Use linker flags to provide version/build/buildTime settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD_COMMIT_HASH) -X=main.BuildTime=$(BUILD_TIME)"

hello:
	@echo "Use \`make help\` para obter instruções."

run:
	@echo "Executando main.go"
	cd cmd && \
	go run main.go

clean:
	@echo "Limpando compilacoes anteriores"
	rm -rf bin && cd cmd && rm -rf temp-bin

build: clean
	@echo "Compilando para a plataforma atual"
	
	cd cmd && \
	go build -o temp-bin/main main.go && \
	cd .. && \
	mv cmd/temp-bin bin && \
	ls -lh bin

compile: clean
	@echo "Compilando projeto para todas as plataformas"
	
	# 32-Bit Systems
	# FreeBDS
	cd cmd && \
	GOOS=freebsd GOARCH=386 go build $(LDFLAGS) -o temp-bin/main-freebsd-386 main.go
	
	# Linux
	@echo " Linux 32-Bit"
	cd cmd && \
	GOOS=linux GOARCH=386 go build $(LDFLAGS) -o temp-bin/main-linux-386 main.go
	
	# Windows
	@echo " Windows 32-Bit"
	cd cmd && \
	GOOS=windows GOARCH=386 go build $(LDFLAGS) -o temp-bin/main-windows-386 main.go
	
	# 64-Bit
	# FreeBDS
	cd cmd && \
	GOOS=freebsd GOARCH=amd64 go build $(LDFLAGS) -o temp-bin/main-freebsd-amd64 main.go
	
	# MacOS
	cd cmd && \
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o temp-bin/main-darwin-amd64 main.go
	
	# Linux
	cd cmd && \
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o temp-bin/main-linux-amd64 main.go
	
	# Windows
	cd cmd && \
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o temp-bin/main-windows-amd64 main.go

	# Movendo pasta cmd/temp-bin para bin
	mv cmd/temp-bin bin

	# Listando arquivos compilados
	ls -lh bin

all: compile

.PHONY: help
help: Makefile
	@echo
	@echo " Comandos disponiveis em "$(PROJECTNAME)":"
	@echo
	@echo " make run"
	@echo " make clean"
	@echo " make build"
	@echo " make compile"
	@echo " make all"
	@echo
