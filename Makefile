crypto:
	protoc -I proto --go_opt=module=github.com/lucassauro/klever-challenge --go_out=. --go-grpc_opt=module=github.com/lucassauro/klever-challenge --go-grpc_out=. proto/*.proto

cserver:
	go build -o bin ./server

cclient:
	go build -o ./bin ./client