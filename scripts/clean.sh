#!/usr/bin/env bash

# globals from makefile
PROJECT_ROOT=${PROJECT_ROOT}

BIN_DIR=${PROJECT_ROOT}/bin

find ${BIN_DIR} -type f -not -name '.gitignore' -delete