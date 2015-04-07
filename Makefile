test: clean all
	go test -v github.com/alexlyulkov/conf/conf
	go test -v github.com/alexlyulkov/conf/server
coverage: clean
	gocov test github.com/alexlyulkov/conf/conf | gocov report
	gocov test github.com/alexlyulkov/conf/server | gocov report
all:
	go install github.com/alexlyulkov/conf
deps:
	go get
clean:
	find . -name flymake_* -delete
run: all
	conf -address "0.0.0.0:8080" -workdir "/var/tmp/alex_config"

