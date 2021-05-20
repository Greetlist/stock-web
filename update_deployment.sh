#!/bin/bash

version=$(cat version.txt)
echo 'current version is : ', $version

cd web_application
go build cmd/main.go
cd -
sudo docker build -t stock-web:$version .
kubectl set image deployment/stock-web-deployment stock-web=stock-web:$version --record
