#!/usr/bin/env bash

docker run --name postgres \
 -p 5432:5432 \
 -e POSTGRES_USER=dev \
 -e POSTGRES_PASSWORD=dev \
 -e POSTGRES_DB=sts \
 -v ./init.sql:/docker-entrypoint-initdb.d/init.sql \
 -d postgres