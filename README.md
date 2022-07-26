# goql

## Install

```bash
# linux
curl -L https://github.com/danangkonang/goql/releases/download/0.1.5/goql_linux -o goql
chmod +x goql
# testing
goql --version

# macOs
curl -L https://github.com/danangkonang/goql/releases/download/0.1.5/goql_macOs -o goql
chmod +x goql
# testing
goql --version

# windows
curl -L https://github.com/danangkonang/goql/releases/download/0.1.5/goql_windows -o goql.exe
# chmod +x goql.exe
# testing
goql.exe --version
```

## Usage
```bash
goql --help
```
# Data Migration

```bash
# generate sql file
goql create migration --table users

# execute sql file
goql up migration --db postgres://user:password@localhost:5432/migration?sslmode=disable
```

# Data Seeder

goql use [https://github.com/brianvoe/gofakeit](https://github.com/brianvoe/gofakeit) for generate data

```bash
# create seeder
goql create seeder --table users --field "id:uuid,name:name,created_at:unixtime" --count 100
```
