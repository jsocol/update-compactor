GSOURCE=../googleapis

build_proto: proto/%.pb.go proto/%_pb2.py

proto/%.pb.go: proto/*.proto
	protoc -I$(GSOURCE) -I. --go_out=paths=source_relative:. --go_opt=paths=source_relative $?

proto/%_pb2.py: proto/*.proto
	protoc -I$(GSOURCE) -I. --python_out=proto/ $?

.PHONY: test
test: test_go test_python

.PHONY: test_go
test_go:
	go test -v ./...

.PHONY: test_python
test_python:
	python -m unittest compactor_test.py


.PHONY: virtualenv
virtualenv:
	python -m venv venv
	venv/bin/pip install -r requirements.txt
