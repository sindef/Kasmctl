BINRY_NAME=kasmctl
TEMP_FILE=false

tests:
	go test ./...
build:
##Check if /usr/local/bin exists
	if [ ! -d /usr/local/bin ]; then \
		mkdir /usr/local/bin;\
	fi
#Move old binary if it exists
	if [ -f /usr/local/bin/kasmctl ]; then\
		mv /usr/local/bin/kasmctl /tmp/kasmctl.bak;\
	fi

	go build -o /usr/local/bin/kasmctl kasmctl.go

	if [ -f /usr/local/bin/kasmctl ]; then\
		/usr/local/bin/kasmctl --version;\
		echo "Build successful";\
		if [ -f /tmp/kasmctl.bak ]; then\
			rm /tmp/kasmctl.bak;\
		fi;\
	else\
		echo "Build failed";\
		mv /tmp/kasmctl.bak /usr/local/bin/kasmctl;\
	fi

ctests:
	docker build -f Dockerfile -t kasmctl:testbuild .
	docker image rm kasmctl:testbuild --force
	echo -e "\n\nBuild tests successful"
run:
	/usr/local/bin/kasmctl

clean:
	rm /usr/local/bin/kasmctl
	go clean