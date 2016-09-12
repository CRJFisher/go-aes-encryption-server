# Golang encryption server with MongoDB microservice connection

## Setup

#### Encryption server
If you do not want this to store/retrieve data from REST/MongoDB microservice change the 'UseRemoteStore' flag to false

Once the repository is downloaded, install the encryption server:
    In the 'server' folder, run: go install

Then run the server:
    Make sure there is no process at port 8080
    Then run: $GOPATH/bin/server

#### Database server
This requires two mongo packages to compile:
    Run:     go get gopkg.in/mgo.v2
    ... and: go get gopkg.in/mgo.v2/bson

To install the database server:
    In the 'database' folder, type: go install

To run the database server:
    Make sure there is no process at port 8082
    Then run: $GOPATH/bin/database

## Run the client
To run the client script:
    In the client folder, run: go run main.go
