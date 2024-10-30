#!/bin/bash


echo "running migrations"

chmod +x ./wait-for-it.sh

./wait-for-it.sh $DB_HOST:$DB_PORT --timeout=2 -- goose -dir=/migrations status
./wait-for-it.sh $DB_HOST:$DB_PORT --timeout=2 -- goose -dir=/migrations up

echo "migrations done"