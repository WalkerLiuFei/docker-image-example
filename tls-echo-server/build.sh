#!/bin/bash
git log --format="%H" -n 1
echo $logID
echo "Start build image"
docker build .
echo "Build image done"

