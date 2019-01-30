#! /usr/local/bin/bash -x

go get -u golang.org/x/tools/present
go get -u golang.org/x/tools/cmd/present

present -play -notes -orighost localhost

