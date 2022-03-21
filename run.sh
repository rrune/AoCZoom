#!/bin/bash
cd ./ol && bash buildcopy.sh
cd ../src && go run cmd/main.go
