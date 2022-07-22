compile:
	cleare
	build

cleare:
	rm goql_linux goql_windows.exe goql_macOS

build:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o goql_linux goql.go
	GOOS=windows GOARCH=amd64 go build -o goql_windows.exe goql.go
	OOS=darwin GOARCH=amd64 go build -o goql_macOS goql.go
	