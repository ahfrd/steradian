# Guide

## Build image docker 
```
docker build -t ahfrd/steradian:v1 .
```

## Run steradian image on docker
```
docker run -d -p 9018:9018 -v config:/app/config --name steradian-v1 ahfrd/steradian:v1
```# AQ
# steradian
