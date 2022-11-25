#!/bin/sh

for f in $(find -name '*.proto')
do
    dir=`dirname $f`
    file=`basename $f`
    cd $dir
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $file;
    cd ..
  done
