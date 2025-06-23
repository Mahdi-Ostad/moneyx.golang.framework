@echo off
for %%f in (*.proto) do (
    set "file=%%~nf"
    if not exist "%%~nf" mkdir "%%~nf"
    protoc --go_out=%%~nf --go_opt=paths=source_relative ^
           --go-grpc_out=%%~nf --go-grpc_opt=paths=source_relative ^
           %%f
)