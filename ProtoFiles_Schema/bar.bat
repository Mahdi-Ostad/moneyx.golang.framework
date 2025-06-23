@echo off
for %%f in (*.proto) do (
    set "file=%%~nf"
    if not exist "../integration/%%~nf" mkdir "../integration/%%~nf"
    protoc --go_out=../integration/%%~nf --go_opt=paths=source_relative ^
           --go-grpc_out=../integration/%%~nf --go-grpc_opt=paths=source_relative ^
           %%f
)