package helper

import "fmt"

func CreateName(length int) string {
	name := fmt.Sprintf("%04d", length+1)
	return name
}
