#!/bin/sh

#compile checker

source="$1" # Polygon package path
dest="$2" # problem path

cd "$source/files"
g++ -o checker check.cpp
cp ./checker "$dest"

exit 0