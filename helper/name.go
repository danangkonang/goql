/*
Copyright © 2022 DanangKonang danangkonang21@gmail.com
*/
package helper

import "fmt"

func CreateName(length int) string {
	name := fmt.Sprintf("%04d", length+1)
	return name
}

func CreateNextName(length int) string {
	name := fmt.Sprintf("%04d", length)
	return name
}
