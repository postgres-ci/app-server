html:
	build-html -i assets/templates_src -o assets/templates
tests:
	go test -v src/app/models/webhooks/common/*.go
	go test -v src/app/models/webhooks/github/*.go
	go test -v src/tools/limit/*.go
	go test -v src/tools/render/pagination/*.go
