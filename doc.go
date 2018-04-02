//go:generate abigen --sol ./contracts/Crowdsale.sol --pkg gocontracts  --out ./gocontracts/Crowdsale.go
//go:generate abigen --sol ./contracts/Crowdsale.sol --pkg gotest  --out ./gotest/contract.go
//go:generate abigen --sol ./contracts/UsingFiatPrice.sol --pkg main  --out ./services/ExchangeRateService/contract.go

//go:generate protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --go_out=plugins=grpc:.                 api/buyer/buyer.proto
//go:generate protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --grpc-gateway_out=logtostderr=true:.   api/buyer/buyer.proto
//go:generate protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --swagger_out=logtostderr=true:.        api/buyer/buyer.proto

package main
