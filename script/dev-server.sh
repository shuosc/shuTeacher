#!/usr/bin/env bash
export DB_ADDRESS="postgres://postgres@localhost:5432/postgres?sslmode=disable"
export PORT="8001"
cd ./web
gin -p 8000 run main.go