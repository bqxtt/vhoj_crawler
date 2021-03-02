cd pkg/sdk
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative crawlerpb/service.proto
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative base/base.proto
cd ../..

xcopy pkg\sdk\crawlerpb\*.go ..\vhoj_rpc\model\crawlerpb\ /s /y