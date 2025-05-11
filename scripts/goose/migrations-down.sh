#!/bin/bash

cd ./database 
${GOPATH}/bin/goose down -env ./config.env
