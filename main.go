package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	tplFilename string
)

func main() {
	flag.StringVar(&tplFilename, "f", "", "Template filename")
	flag.Parse()

	if tplFilename == "" {
		logrus.Fatal("Template filename required")
	}

	tpl := readTpl(tplFilename)
	bytes := readJSONFromEnv()

	var obj interface{}
	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		logrus.Fatal(err)
	}

	t, err := template.New("tpl").Parse(tpl)
	if err != nil {
		logrus.Fatal("Fail to execute template, err:")
	}
	err = t.ExecuteTemplate(os.Stdout, "tpl", obj)
	if err != nil {
		logrus.Fatal("Fail to execute template, err:", err)
	}
}

func readJSONFromEnv() []byte {
	data := os.Getenv("JTPL_JSON_DATA")
	bytes := []byte(data)

	var obj interface{}
	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		logrus.Fatal("Fail to parse JSON from JTPL_JSON_DATA env var")
	}

	return bytes
}

func readTpl(tplFilename string) string {
	bytes, err := ioutil.ReadFile(tplFilename)
	if err != nil {
		logrus.Fatal("Fail to read template file " + tplFilename)
	}

	return `{{define "tpl"}}` + string(bytes) + `{{end}}`
}
