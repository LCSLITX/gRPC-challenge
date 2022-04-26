# klever-challenge
The Technical Challenge consists of creating an API with Golang using gRPC with stream pipes that exposes an Upvote service endpoints. The API will provide the user an interface to upvote or downvote a known list of the main Cryptocurrencies (Bitcoin, ethereum, litecoin, etc..).

Technical requirements:
- Keep the code in Github

API:
- The API must guarantee the typing of user inputs. If an input is expected as a string, it can only be received as a string.
- The structs used with your mongo model should support Marshal/Unmarshal with bson, json and struct
- The API should contain unit test of methods it uses

Extra:
- Deliver the whole solution running in some free cloud service



## Brief information 

| Object          | Technology   |
| --------------- |:------------:|
| Server          | Go           |
| Client          | Go           |
| API             | gRPC         |
| Database        | MongoDB      |
| Containers      | Docker       |
| Cloud Host      |     -        |


## Instructions to execute

  Requirements

 * You should have [GO (Golang)](https://go.dev/) installed to run both server and client.

 * You could use Docker to execute the server or database (or both) inside a container.

 * Additionally, you may use Make to execute commands easily with Makefile.



#### Clone this repository
 ```sh
# ssh
git@github.com:lucassauro/klever-challenge.git
# https
https://github.com/lucassauro/klever-challenge.git
```


After that, there are a few ways to execute the project;

#### Building a Go binary (easiest and faster):
1. Change to its directory and simply execute the command `make buildRunServer` to build and run it.
2. It'll be available to test it with Insonmnia, Postman or BloomRPC. The application is already connected with a cloud database.


#### Within a container:
1. Change to its directory and execute the command `make runDockerServer` to build and run it inside a container.
2. It'll be available to test it with Insonmnia, Postman or BloomRPC. The application is already connected with a cloud database.


#### With a set of containers, together with mongo and mongo-express:
1. Change to its directory and execute the command `docker-compose up` to build inside a container.
2. It'll be available to test it with Insonmnia, Postman or BloomRPC. The application is already connected with a cloud database.


#### Ports
The gRPC server is configured to be available at `localhost:50051`.

The Mongo server is configured to be available at `localhost:27017`.

The Mongo-Express is configured to be accessible through the browser at URL `localhost:8081`.



#### Requests

To create requests and use the application, it is recommended to use Insomnia. 

There is an exportation .JSON file in the root of this repository. It could be used to import request collection.


### Requirements explanation

As previously stated, the challenge consisted of creating an API with GoLang and gRPC. The main requirements were:

- [X] The API must guarantee the typing of user inputs. If an input is expected as a string, it can only be received as a string.

  It was decided to use a Go package called Validator to assure the inputs type.
        
- [X] The structs used with your mongo model should support Marshal/Unmarshal with bson, json and struct.

  The structs do support Marshall/Unmarshall with bson and json.
        
- [X] The API should contain unit test of methods it uses

  Almost all endpoints implemented contain unit tests. Total test coverage is 58% and could be ckecked with `make tests` command.
        
        
In addition to these essential requirements:
- [X] The API must have a read, insert, delete and update interfaces.

  The API implements the CRUD interfaces.
        
- [X] The API must have a method that stream a live update of the current sum of the votes from a given Cryptocurrency.
        
  The API implements LiveCryptoVotes, which is responsible to stream every vote related to a specific coin. It is not available if running database on Docker because mongo.Watch(), or $changestream stage is only supported on replica sets. Available if running database on cloud.

        
