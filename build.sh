#!/bin/bash
mkdir -p build
cd js && npm install && npm run build
cd ..
go get github.com/jteeuwen/go-bindata/...
go get
go-bindata build && go build
