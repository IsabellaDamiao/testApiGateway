workdir: $GOPATH/src/github.com/testApiGateway
observe: *.go *.js
ignore: /vendor
formation: web=2
build-server: make server
web: restart=fail waitfor=localhost:8888 ./server serve