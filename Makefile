BIN_PATH=bin/atap
INSTALL_PATH=/usr/local/bin/atap

all: atap

dependencies: vendor/
	git submodule init
	git submodule update

atap: dependencies atap.go
	go build -o ${BIN_PATH}

install: ${BIN_PATH}
	@cp ${BIN_PATH} ${INSTALL_PATH}
	@echo "Installed atap!"

uninstall:
	@rm -f ${INSTALL_PATH}
	@echo "Uninstalled atap"

clean: ${BIN_PATH}
	rm ${BIN_PATH}