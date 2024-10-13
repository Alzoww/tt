package main

import (
	"fmt"
	"strings"
)

//func main() {
//	path := "/      "
//
//	trimmed := strings.TrimSpace(path)
//
//	fmt.Println(path)
//	fmt.Println(strings.TrimSpace(path))
//	fmt.Println(trimmed != "" || trimmed == "/")
//}

func main() {
	path := "/012/my pc/my  drivers/common /"

	trimmed := strings.TrimSpace(path[strings.LastIndex(path, "/")+1:])

	fmt.Println(trimmed != "")
}
