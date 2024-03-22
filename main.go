package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io"
	"os"

	"github.com/sirupsen/logrus"
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
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		logrus.Fatal("Fail to read stdin")
	}

	var obj interface{}
	err = json.Unmarshal(bytes, &obj)
	if err != nil {
		logrus.Fatal("Fail to parse JSON from stdin")
	}

	return bytes
}

func readTpl(tplFilename string) string {
	bytes, err := os.ReadFile(tplFilename)
	if err != nil {
		logrus.Fatal("Fail to read template file " + tplFilename)
	}

	return `{{define "tpl"}}` + string(bytes) + `{{end}}`
}
