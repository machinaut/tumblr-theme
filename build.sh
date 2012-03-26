#!/bin/bash

# build css files
lessc tvmblr.less > tvmblr.css

# build go program
go build

# run go program and generate output
./tvmblr > output.html
