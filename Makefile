test:
	go test

build:
	#Check if /usr/local/bin exists
	if [ ! -d /usr/local/bin ]; then
		mkdir /usr/local/bin
	fi
	#Remove old binary if it exists
	if [ -f /usr/local/bin/kasmctl ]; then
		rm /usr/local/bin/kasmctl
	fi
	go build -o /usr/local/bin/kasmctl kasmctl.go
	
run:
	go run kasmctl.go