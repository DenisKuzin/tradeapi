FROM ubuntu:22.04
RUN apt update && apt install -y git protoc-gen-go golang-go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN git clone https://github.com/DenisKuzin/trade-api-docs.git
WORKDIR /trade-api-docs
COPY gen.sh /
CMD /gen.sh