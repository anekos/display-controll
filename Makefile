.PHONY: build
build:
	go build
	sudo chown root:root display-controll
	sudo chmod u+s display-controll
