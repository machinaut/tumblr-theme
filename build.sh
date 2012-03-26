#!/bin/bash

# build css files
lessc posts.less > posts.css

# build go program
go build

# run go program and generate output
./tvmblr > output.html
