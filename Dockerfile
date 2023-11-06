FROM ubuntu:22.04
RUN apt update && apt install -y git protoc-gen-go golang-go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN git clone https://github.com/DenisKuzin/trade-api-docs.git
WORKDIR /trade-api-docs
RUN apt install -y curl
RUN curl -o /usr/bin/include/google/type/date.proto --create-dirs https://raw.githubusercontent.com/googleapis/googleapis/master/google/type/date.proto
COPY gen.sh /
CMD /gen.sh