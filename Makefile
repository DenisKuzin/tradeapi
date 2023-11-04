all:
	docker build -t protoc .
	docker run --mount src=$(shell pwd),target=/tradeapi,type=bind protoc
