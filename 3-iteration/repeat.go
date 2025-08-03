package iteration

import "strings"

// Refactor readable
// const repeatCount = 5

// func Repeat(character string) string {
// 	var repeated string
// 	for i:= 0; i < repeatCount ; i++ {
// 		repeated +=  character
// 	}
// 	return repeated
// }

// di test benchmark lebih rungan

const repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
