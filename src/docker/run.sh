#!/bin/sh

docker build -t postgres-email-image -f docker/local.Dockerfile .
ls -a
docker run -p 5436:5432 --name postgres-email-container --rm --detach --env-file .env/.postgre.env postgres-email-image