#!/bin/sh

docker build -t postgres-email-image -f docker/local.Dockerfile .
docker run -p 5435:5432 --name postgres-email-container -d postgres-email-image