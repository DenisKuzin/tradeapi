#!/bin/bash
set -x
export PATH=$PATH:$(go env GOPATH)/bin
protoc --go_out=/tradeapi --go_opt=paths=import \
       --go_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       proto/tradeapi/v1/common.proto

protoc --go_out=/tradeapi --go_opt=paths=import \
       --go_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go_opt=Mproto/tradeapi/v1/candles.proto=github.com/DenisKuzin/tradeapi \
       proto/tradeapi/v1/candles.proto

protoc --go_out=/tradeapi --go_opt=paths=import \
       --go_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go_opt=Mproto/tradeapi/v1/portfolios.proto=github.com/DenisKuzin/tradeapi \
       --go_opt=Mproto/tradeapi/v1/orders.proto=github.com/DenisKuzin/tradeapi \
       --go_opt=Mproto/tradeapi/v1/events.proto=github.com/DenisKuzin/tradeapi \
       proto/tradeapi/v1/events.proto

protoc --go_out=/tradeapi --go_opt=paths=import \
       --go_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go_opt=Mproto/tradeapi/v1/orders.proto=github.com/DenisKuzin/tradeapi \
       proto/tradeapi/v1/orders.proto

protoc --go_out=/tradeapi --go_opt=paths=import \
       --go_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go_opt=Mproto/tradeapi/v1/portfolios.proto=github.com/DenisKuzin/tradeapi \
       proto/tradeapi/v1/portfolios.proto

protoc --go_out=/tradeapi --go_opt=paths=import \
       --go_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go_opt=Mproto/tradeapi/v1/security.proto=github.com/DenisKuzin/tradeapi \
       proto/tradeapi/v1/security.proto

protoc --go_out=/tradeapi --go_opt=paths=import \
	--go_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
	--go_opt=Mproto/tradeapi/v1/stops.proto=github.com/DenisKuzin/tradeapi \
	proto/tradeapi/v1/stops.proto

protoc --go-grpc_out=/tradeapi --go-grpc_opt=paths=import \
       --go-grpc_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/candles.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/candles.proto=github.com/DenisKuzin/tradeapi \
       grpc/tradeapi/v1/candles.proto

protoc --go-grpc_out=/tradeapi --go-grpc_opt=paths=import \
       --go-grpc_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/portfolios.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/orders.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/events.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/events.proto=github.com/DenisKuzin/tradeapi \
       grpc/tradeapi/v1/events.proto

protoc --go-grpc_out=/tradeapi --go-grpc_opt=paths=import \
       --go-grpc_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/portfolios.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/orders.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/events.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/events.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/orders.proto=github.com/DenisKuzin/tradeapi \
       grpc/tradeapi/v1/orders.proto

protoc --go-grpc_out=/tradeapi --go-grpc_opt=paths=import \
       --go-grpc_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/portfolios.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/orders.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/events.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/events.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/portfolios.proto=github.com/DenisKuzin/tradeapi \
       grpc/tradeapi/v1/portfolios.proto

protoc --go-grpc_out=/tradeapi --go-grpc_opt=paths=import \
       --go-grpc_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/security.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/portfolios.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/orders.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/events.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/stops.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/events.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/portfolios.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/securities.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/stops.proto=github.com/DenisKuzin/tradeapi \
       grpc/tradeapi/v1/stops.proto

protoc --go_out=/tradeapi --go_opt=paths=import \
       --go_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go_opt=Mproto/tradeapi/v1/security.proto=github.com/DenisKuzin/tradeapi \
       --go_opt=Mgrpc/tradeapi/v1/securities.proto=github.com/DenisKuzin/tradeapi \
       grpc/tradeapi/v1/securities.proto

protoc --go-grpc_out=/tradeapi --go-grpc_opt=paths=import \
       --go-grpc_opt=Mproto/tradeapi/v1/common.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/security.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/portfolios.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/orders.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mproto/tradeapi/v1/events.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/events.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/portfolios.proto=github.com/DenisKuzin/tradeapi \
       --go-grpc_opt=Mgrpc/tradeapi/v1/securities.proto=github.com/DenisKuzin/tradeapi \
       grpc/tradeapi/v1/securities.proto

mv /tradeapi/github.com/DenisKuzin/tradeapi/* /tradeapi/
rm -rf /tradeapi/github.com
