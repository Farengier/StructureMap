#!/usr/bin/env bash

# globals from makefile
PROJECT_ROOT=${PROJECT_ROOT}

CUR_DIR=`pwd`

cd ${PROJECT_ROOT}

find pkg -type f -name '*.go' -print > list
xgettext --language=C --keyword=T --files-from=list --sort-output --output-dir=localization
unlink list

cd ${CUR_DIR}