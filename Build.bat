REM protoc -I "pkg/protoc/include" -I "src/common/models" --go_out=. src/common/models/*.proto
REM cd ../../..

protoc -I "pkg/protoc/include" -I "src/common/models" --go_out=src/common/models --go_opt=paths=source_relative --go-grpc_out=src/common/models --go-grpc_opt=paths=source_relative src/common/models/*.proto

REM cd ../../..