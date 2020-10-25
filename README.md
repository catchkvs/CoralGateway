# CoralGateway
A gateway server for Apps local storage to keep it in sync with server

### Starting coral Gateway
#### Getting started
1. Clone the repo: `git clone https://github.com/catchkvs/CoralGateway.git`
2. Go to the root : `cd CoralGateway`
3. Run the build: `go build -o server pkg/coralgateway.go`
4. Run `./server`
#### Using docker container
1. Build the docker image: `docker build -t coral-gateway:0.1 .`
2. Run the docker container: `docker run -p 3030:3030 coral-gateway:0.1`

### Pebble DB as storage
By default currently the data is stored in pebbleDB.
Plan is to add support for MongoDB and Google cloud firestore.