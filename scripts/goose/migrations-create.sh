#!/bin/bash

echo "Please enter a name for the new migration:"
read migration_name

cd ./database
${GOPATH}/bin/goose create ${migration_name} sql -env ./config.env
