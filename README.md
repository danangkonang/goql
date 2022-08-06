# goql

## Install

```bash
# linux
curl -L https://github.com/danangkonang/goql/releases/download/0.1.1/goql_linux -o goql
chmod +x goql
# testing
goql --version

# macOs
curl -L https://github.com/danangkonang/goql/releases/download/0.1.1/goql_macOs -o goql
chmod +x goql
# testing
goql --version

# windows
curl -L https://github.com/danangkonang/goql/releases/download/0.1.1/goql_windows -o goql.exe
# chmod +x goql.exe
# testing
goql.exe --version
```

## Usage
```bash
goql --help
```

# Data Seeder

goql use [https://github.com/brianvoe/gofakeit](https://github.com/brianvoe/gofakeit) for generate data

```bash
# create seeder
goql create seeder --table users --field "name:name,created_at:unixtime" --count 100

goql create seeder --table products --field "id:uuid" --count 10
```