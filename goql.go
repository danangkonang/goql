/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com

*/
package main

import (
	"github.com/danangkonang/goql/cmd"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	cmd.Execute()
}
