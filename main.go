package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type pageData struct {
	Stats    string
	Meminfo  string
	Hostname string
}

func main() {
	http.HandleFunc("/", stats)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func stats(w http.ResponseWriter, r *http.Request) {
	dat, _ := ioutil.ReadFile("/sys/fs/cgroup/memory/memory.stat")
	result := strings.Replace(string(dat), "\n", "<br>", -1)

	dat2, _ := ioutil.ReadFile("/etc/hosts")
	result2 := strings.Replace(string(dat2), "\n", "<br>", -1)

	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	p := pageData{
		Stats:    result,
		Meminfo:  result2,
		Hostname: os.Getenv("HOSTNAME"),
	}

	t.Execute(w, p)

	var out bytes.Buffer
	cmd := exec.Command("cat", "/sys/fs/cgroup/memory/memory.stat")
	if err != nil {
		log.Fatal(err)
	}

	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
