#!/usr/bin/env bash

# globals from makefile
PROJECT_ROOT=${PROJECT_ROOT}

EXECUTABLE=smap
BIN_DIR=${PROJECT_ROOT}/bin
CUR_DIR=`pwd`

cd ${BIN_DIR}
go build -o ${EXECUTABLE} ${PROJECT_ROOT}/cmd/smap/main.go
chmod +x ${BIN_DIR}/${EXECUTABLE}

cd ${CUR_DIR}