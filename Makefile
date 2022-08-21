cleare:
	rm goql_linux goql_windows.exe goql_macOS

build:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o goql_linux goql.go
	GOOS=windows GOARCH=amd64 go build -o goql_windows.exe goql.go
	OOS=darwin GOARCH=amd64 go build -o goql_macOS goql.go
upm:
	go run goql.go up migration --db postgres://postgres:postgres@localhost:5432/migration?sslmode=disable
downm:
	go run goql.go down migration --db postgres://postgres:postgres@localhost:5432/migration?sslmode=disable
ups:
	go run goql.go up seeder --db postgres://postgres:postgres@localhost:5432/migration?sslmode=disable
downs:
	go run goql.go down seeder --db postgres://postgres:postgres@localhost:5432/migration?sslmode=disable