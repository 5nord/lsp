// This program reads the TTCN-3 grammar file and generates a parser for it.
package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	modelFile = "../language-server-protocol/_specifications/lsp/3.17/metaModel/metaModel.json"
)

func main() {

	b, err := ioutil.ReadFile(modelFile)
	if err != nil {
		log.Fatal(err)
	}

	var model MetaModel
	if err := json.Unmarshal(b, &model); err != nil {
		log.Fatal(err)
	}

	files, err := filepath.Glob("./internal/gen/_templates/*")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, file := range files {
		t, err := template.New(file).Funcs(template.FuncMap{
			"type": func(t string) string {
				switch t {
				case "uinteger", "integer":
					return "int"
				default:
					return t
				}
			},
			"title": strings.Title,
			"comment": func(s string) string {
				var ret []string
				for _, line := range strings.Split(s, "\n") {
					ret = append(ret, "// "+line)
				}
				return strings.Join(ret, "\n")
			},
		}).ParseFiles(file)
		if err != nil {
			log.Fatal(err.Error())
		}
		var b bytes.Buffer
		if err := t.ExecuteTemplate(&b, filepath.Base(file), model); err != nil {
			log.Fatal(err.Error())
		}
		writeSource(filepath.Base(file), b.Bytes())
	}
}

func writeSource(file string, b []byte) {
	b2, err := format.Source(b)
	if err == nil {
		b = b2
	}
	// Always write to file to help debugging.
	if err := os.WriteFile(file, b, 0644); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatalf("%s: error: %s\n", file, err.Error())
	}
}
