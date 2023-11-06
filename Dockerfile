FROM ubuntu:22.04
RUN apt update && apt install -y git protoc-gen-go golang-go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN git clone https://github.com/DenisKuzin/trade-api-docs.git
WORKDIR /trade-api-docs/contracts
RUN apt install -y curl
RUN curl -o google/type/date.proto --create-dirs https://raw.githubusercontent.com/googleapis/googleapis/master/google/type/date.proto
COPY gen.sh /trade-api-docs/contracts
CMD /trade-api-docs/contracts/gen.sh