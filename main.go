package main

import (
	"fmt"
	"os"
)

func main() {
	lesson := os.Args[1]

	createFile := func(name string, content string) {
		d := []byte(content)
		os.WriteFile(name, d, 0644)
	}

	err := os.Mkdir(lesson, 0755)

	if err != nil {
		panic(err)
	}

	createFile(fmt.Sprintf("%s/%s.go", lesson, lesson), fmt.Sprintf("package %s", lesson))
	createFile(fmt.Sprintf("%s/%s_test.go", lesson, lesson), fmt.Sprintf("package %s\n\nimport (\n  \"testing\"\n)\n\nfunc Test(t *testing.T) {}", lesson))
}
