CUR_DIR = $(shell cd -P .;pwd)
ROOT_DIR = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
SCRIPTS_DIR = ${ROOT_DIR}/scripts

export PROJECT_ROOT=${ROOT_DIR}

build: clean
	@printf "%-60s\n" "    **** Building ****"
	@${SCRIPTS_DIR}/build.sh
	@echo "[OK]"

clean:
	@printf "%-60s\n" "    **** Clean ****"
	@${SCRIPTS_DIR}/clean.sh
	@echo "[OK]"

localize:
	@printf "%-60s\n" "    **** Localization ****"
	@${SCRIPTS_DIR}/localize.sh
	@echo "[OK]"

.PHONY: build clean localize