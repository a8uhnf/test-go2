package test_go2

import (
	"fmt"
	b "github.com/a8uhnf/test-go1"
)

func HelloWorld()  {
	fmt.Println("Hello World from test-go2")
	b.HelloWorld()
}

func HelloWorldMaster()  {
	fmt.Println("Hello World From master branch")
	b.HelloWorld()
}
