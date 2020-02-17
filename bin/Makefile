
client:
	go build -x -o ./binary/ves-client ./ves-client

ves-server:
	go build -x -o ./binary/ves-server ../ves
	if [ ! -f "./binary/ves-config.toml" ]; then cp ../ves/config.toml ./binary/ves-config.toml; fi;

cves-server:
	go build -x -o ./binary/central-ves-server ../central-ves
	if [ ! -f "./binary/cves-config.toml" ]; then cp ../central-ves/config.toml ./binary/cves-config.toml; fi;

server: ves-server cves-server

up: up-ves up-cves

up-ves: ves-server
	nohup ./binary/ves-server --config ./binary/ves-config.toml >> ves.out &

down-cves:
	kill -15 `pidof central-ves-server`

up-cves: cves-server
	nohup ./binary/central-ves-server --config ./binary/cves-config.toml >> cves.out &

login:
	./binary/ves-client

repl:
	./ves-client/repl.py

.PHONY: client server ves-server cves-server up up-ves up-cves down-cves repl
