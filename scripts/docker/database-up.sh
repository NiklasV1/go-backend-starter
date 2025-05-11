#!/bin/bash

docker compose build database
docker compose up database -d
