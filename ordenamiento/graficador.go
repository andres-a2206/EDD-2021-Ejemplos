package main

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./diagrama.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("outfile.png", cmd, os.FileMode(mode))
}
