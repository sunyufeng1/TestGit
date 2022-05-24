package main

import (
	"github.com/sunyufeng1/CommTool/fileTool"
	"github.com/sunyufeng1/LeeCode/uniteTest"
)

func main() {
	testObj := new(uniteTest.Obj)
	testObj.Test(3)
	fileTool.IsExist("l")

}
