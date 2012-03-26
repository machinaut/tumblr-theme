#!/bin/bash

# build css files
lessc tvmblr.less > tvmblr.css

# build go program
go build -o buildr

# run go program and generate output
./buildr > theme.html
