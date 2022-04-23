crypto:
	protoc -I proto --go_opt=module=github.com/lucassauro/klever-challenge --go_out=. --go-grpc_opt=module=github.com/lucassauro/klever-challenge --go-grpc_out=. proto/*.proto

# build server
bserver:
	go build -o ./bin/ ./server && ./bin/server
# build client
bclient:
	go build -o ./bin/ ./client && ./bin/client