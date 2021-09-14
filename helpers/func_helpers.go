package helpers

import "fmt"

func FailOnError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}
