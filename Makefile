fumpt:
	gofumpt -l -w .

lint:
	golangci-lint run

imports:
	gci -w -local github.com/illfalcon/spotiyan cmd
	gci -w -local github.com/illfalcon/spotiyan internal
	gci -w -local github.com/illfalcon/spotiyan pkg

fmt: fumpt imports