build_proto: proto/%.pb.go proto/%_pb2.py

proto/%.pb.go: proto/*.proto
	protoc --go_out=paths=source_relative:. --go_opt=paths=source_relative $?

proto/%_pb2.py: proto/*.proto
	protoc --python_out=proto/ $?
