html:
	build-html -i assets/templates_src -o assets/templates
tests:
	go test -v src/app/models/webhooks/common/*.go
	go test -v src/app/models/webhooks/github/*.go