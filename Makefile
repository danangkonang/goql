compile: rm build

rm:
	rm goql_linux goql_windows.exe goql_macOS

build:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o goql_linux main.go
	GOOS=windows GOARCH=amd64 go build -o goql_windows.exe main.go
	OOS=darwin GOARCH=amd64 go build -o goql_macOS main.go