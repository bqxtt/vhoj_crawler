module github.com/ecnuvj/vhoj_crawler

go 1.14

require (
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/ecnuvj/vhoj_common v0.0.0-20210204125811-c22717ad12a6
	github.com/ecnuvj/vhoj_db v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)

replace (
	github.com/ecnuvj/vhoj_common => ../vhoj_common
	github.com/ecnuvj/vhoj_db => ../vhoj_db
)
