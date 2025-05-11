#!/bin/bash

cd ./database
${GOPATH}/bin/goose up -env ./config.env
