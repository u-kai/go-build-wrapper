package main

import (
	"os"
	"os/exec"
	"strings"
)

var targetDir = os.Getenv("GOPATH")

var defaultTargetDir = os.Getenv("HOME") + "/go/bin"

func main() {
	dir := targetDir + "/bin"
	if dir == "/bin" {
		dir = defaultTargetDir
	}
	thisDir, _ := os.Getwd()
	projectName := thisDir[strings.LastIndex(thisDir, "/")+1:]
	if len(os.Args) > 1 {
		projectName = os.Args[1]
	}
	target := dir + "/" + projectName

	output, err := exec.Command("go", "build", "-o", target).CombinedOutput()
	if err != nil {
		println(string(output))
		panic(err)
	}
	println(string(output))
}
