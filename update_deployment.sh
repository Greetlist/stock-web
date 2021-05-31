#!/bin/bash

old_version=$(cat version.txt)
echo 'current version is : ', $old_version

new_version=$(awk "BEGIN{print $old_version+0.001}")
echo $new_version > version.txt

sudo docker rmi -f stock-web:$old_version

cd server
go build cmd/main.go
cd -
sudo docker build -t stock-web:$new_version .
kubectl set image deployment/stock-web-deployment stock-web=stock-web:$new_version --record
