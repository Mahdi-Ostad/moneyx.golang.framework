if not exist "../integration/%1" mkdir "../integration/%1"
protoc --go_out=../integration/%1 --go_opt=paths=source_relative ^
    --go-grpc_out=../integration/%1 --go-grpc_opt=paths=source_relative ^
    %1.proto