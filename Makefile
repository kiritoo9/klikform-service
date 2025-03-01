rundev:
	GO_ENV=development go run src/entry/main.go
runprod:
	GO_ENV=production go run src/entry/main.go