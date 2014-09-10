#!/bin/bash
mkdir -p build
touch build/index.html

cd js && npm install && npm run build
cd ..

if ! [ -f $GOPATH/bin/go-bindata ]; then
  go get github.com/jteeuwen/go-bindata/...
fi

cd assets

appendat=`grep -rnw "<head>" index.html | cut -f2 -d:`
total=`wc -l < index.html | tr -d ' '`
remaining=$((total-appendat))

> ../build/index.html
echo "`head -$appendat index.html`" >> ../build/index.html
echo "<script>" >> ../build/index.html
cat ../build/bundle.js >> ../build/index.html
echo "</script>" >> ../build/index.html
echo "`tail -$remaining index.html`" >> ../build/index.html

cd ..

rm build/bundle.js
go-bindata build
go get
go build
rm bindata.go
rm build/index.html
rmdir build
