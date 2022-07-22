```
go run goql.go create seeder --table users --field "name:name,created_at:unixtime" --count 100

go run goql.go create seeder --table products --field "id:uuid" --count 10
```