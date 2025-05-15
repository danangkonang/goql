cleare:
	rm -f goql_linux goql_windows.exe goql_macOS

build:
	echo "Compiling for every OS and Platform"
	rm -f goql_linux goql_windows.exe goql_macOS
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

mysql:
	go run goql.go up migration --dir schema/migration --db "mysql://danang:danang@tcp(localhost:3306)/simcard?parseTime=true&loc=Asia%2FJakarta"
	# ./goql up migration --dir migration --db "mysql://danang:danang@(localhost:3306)/simcard?parseTime=true"

	# go run goql.go down migration --dir migration --db "mysql://root:root@(localhost:3306)/db_hr?parseTime=true"
