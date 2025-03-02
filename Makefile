swag:
	swag init -g entry/main.go -d src -o docs --parseDependency --parseInternal

rundev:
	GO_ENV=development go run src/entry/main.go

runprod: swag
	GO_ENV=production go run src/entry/main.go