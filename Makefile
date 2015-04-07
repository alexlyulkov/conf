test: clean all
	go test -v github.com/alexlyulkov/conf/conf
coverage: clean
	gocov test github.com/alexlyulkov/conf/conf | gocov report
all:
	go install github.com/alexlyulkov/conf
deps:
	go get
clean:
	find . -name flymake_* -delete
run: all
	conf

