fumpt:
	gofumpt -l -w .

lint:
	golangci-lint run

imports:
	gci -w -local github.com/illfalcon/spotiyan

fmt: fumpt imports