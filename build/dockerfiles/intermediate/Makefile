ifeq ($(GO_VERSION),)
  GO_VERSION=1.21
endif
ifeq ($(RUST_VERSION),)
  RUST_VERSION=nightly-2023-12-03
endif
ifeq ($(PYTHON_VERSION),)
  PYTHON_VERSION=3.10
endif
ifeq ($(CUDA_VERSION),)
  CUDA_VERSION=11.7.1
endif

.PHONY: all go-alpine-builder rust-builder rust-alpine-builder go-rust-alpine-builder go-rust-builder cuda-go-rust-builder py-runner

cuda-go-rust-builder:
	docker build -t scrolltech/cuda-go-rust-builder:cuda-$(CUDA_VERSION)-go-$(GO_VERSION)-rust-$(RUST_VERSION) -f cuda-go-rust-builder.Dockerfile ./ --build-arg CUDA_VERSION=$(CUDA_VERSION) --build-arg GO_VERSION=$(GO_VERSION) --build-arg RUST_VERSION=$(RUST_VERSION)

go-rust-builder:
	docker build -t scrolltech/go-rust-builder:go-$(GO_VERSION)-rust-$(RUST_VERSION) -f go-rust-builder.Dockerfile ./ --build-arg GO_VERSION=$(GO_VERSION) --build-arg RUST_VERSION=$(RUST_VERSION)

go-alpine-builder:
	docker build -t scrolltech/go-alpine-builder:$(GO_VERSION) -f go-alpine-builder.Dockerfile ./ --build-arg GO_VERSION=$(GO_VERSION)

rust-builder:
	docker build -t scrolltech/rust-builder:$(RUST_VERSION) -f rust-builder.Dockerfile ./ --build-arg RUST_VERSION=$(RUST_VERSION)

rust-alpine-builder:
	docker build -t scrolltech/rust-alpine-builder:$(RUST_VERSION) -f rust-alpine-builder.Dockerfile ./ --build-arg RUST_VERSION=$(RUST_VERSION)

go-rust-alpine-builder:
	docker build -t scrolltech/go-rust-alpine-builder:go-$(GO_VERSION)-rust-$(RUST_VERSION) -f go-rust-alpine-builder.Dockerfile ./ --build-arg GO_VERSION=$(GO_VERSION) --build-arg RUST_VERSION=$(RUST_VERSION)

py-runner:
	docker build -t scrolltech/py-runner:$(PYTHON_VERSION) -f py-runner.Dockerfile ./ --build-arg PYTHON_VERSION=$(PYTHON_VERSION)

all: go-alpine-builder rust-builder rust-alpine-builder go-rust-alpine-builder go-rust-builder cuda-go-rust-builder py-runner

publish:
	docker push scrolltech/go-alpine-builder:$(GO_VERSION)
	docker push scrolltech/rust-builder:$(RUST_VERSION)
	docker push scrolltech/rust-alpine-builder:$(RUST_VERSION)
	docker push scrolltech/go-rust-alpine-builder:go-$(GO_VERSION)-rust-$(RUST_VERSION)
	docker push scrolltech/go-rust-builder:go-$(GO_VERSION)-rust-$(RUST_VERSION)
	docker push scrolltech/cuda-go-rust-builder:cuda-$(CUDA_VERSION)-go-$(GO_VERSION)-rust-$(RUST_VERSION)
	docker push scrolltech/py-runner:$(PYTHON_VERSION)
