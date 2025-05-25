#!/bin/sh

# build the server
go build -o bin/server ./server

# run the server
./bin/server