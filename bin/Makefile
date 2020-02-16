
client:
	go build -x -o ./binary/ves-client ./ves-client

server:
	go build -x -o ./binary/ves-server ../ves
	go build -x -o ./binary/central-ves-server ../central-ves
	if [ ! -f "./binary/ves-config.toml" ]; then cp ../ves/config.toml ./binary/ves-config.toml; fi;
	if [ ! -f "./binary/cves-config.toml" ]; then cp ../central-ves/config.toml ./binary/cves-config.toml; fi;

up:
	./binary/ves-server --config ./binary/ves-server-config.toml

up-central:
	./binary/central-ves-server

login:
	./binary/ves-client

.PHONY: client server up up-central
