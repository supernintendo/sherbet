#!/bin/bash
function cleanup {
    rm build/bundle.js
    rm bindata.go
    rmdir build
}
cleanup
mkdir -p build
cd js && npm install && npm run build
cd ..
go get
go-bindata build
go build
cleanup
