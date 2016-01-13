package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

import "os/exec"
import "log"

func main() {
	// run()
	// start()
	// stdoutpipe()
	newline()
}

func run() {
	cmd := exec.Command("ls")

	var out bytes.Buffer
	fmt.Println("run function")
	fmt.Println("bytes.Buffer is zeroed out", out)
	fmt.Println("cmd.Stdout is", cmd.Stdout)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("out is\n", out.String())
}

func start() {
	cmd := exec.Command("ls")

	var out bytes.Buffer
	fmt.Println("start function")
	fmt.Println("bytes.Buffer is zeroed out", out)
	fmt.Println("cmd.Stdout is", cmd.Stdout)
	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// if err := cmd.Wait(); err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("out is\n", out.String())

}

func stdoutpipe() {
	cmd := exec.Command("cat", "file.txt")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	// defer cmd.Wait()  //I want to understand this commmand before using it!
	buf := bufio.NewScanner(stdout)
	// allText := []string{} // same as var allText []string

	for buf.Scan() {
		fmt.Println(buf.Text())
		text := buf.Text()
		result := strings.Fields(text)
		for i := range result {
			fmt.Println("occurence ", i, "is", result[i])
		}
	}
	// fmt.Println(allText)

	cmd.Wait()
}

func newline() {
	dat, _ := ioutil.ReadFile("file.txt")
	// fmt.Println(strings.Contains(string(dat), "\n"))
	result := strings.Replace(string(dat), "\n", "<br>", -1)
	fmt.Println(result)

}
