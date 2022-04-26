crypto:
	protoc -I proto --go_opt=module=github.com/lucassauro/klever-challenge --go_out=. --go-grpc_opt=module=github.com/lucassauro/klever-challenge --go-grpc_out=. proto/*.proto

# build server
buildServer:
	go build -o ./bin/ ./src/server

# Docker
runDockerServer:
	docker build -t klever . && docker run -p 50051:50051 klever

# run server
runServer:
	./bin/server

# build and run
buildRunServer:
	go build -o ./bin/ ./src/server && ./bin/server
	
# build client
buildRunClient:
	go build -o ./bin/ ./src/client && ./bin/client

# tests
tests:
	cd src/server && make tests

