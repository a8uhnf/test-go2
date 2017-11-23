package test_go2

import (
	"fmt"
	// a "github.com/a8uhnf/test-go3"
	// "golang.org/x/tools/go/gcimporter15/testdata"
)

func HelloWorld()  {
	fmt.Println("Hello World from test-go2")
	// a.HelloWorld()
}

func HelloWorldTestDep()  {
	fmt.Println("Hello World from test-go2. Branch: test-dep")
	// a.HelloWorld()
}
