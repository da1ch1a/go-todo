#!/bin/bash

docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "mysql://root:rootpassword@tcp(localhost:33060)/go_todo_local" up
