//usr/bin/go run $0 $@ ; exit
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"text/template"

	"gopkg.in/yaml.v2"
)

var data = `
repos:
- name: repo1
  repo: https://github.com/aminjam/repo1
  branch: master
- name: repo2
  repo: https://github.com/john/repo2
  branch: develop
`

type Structure struct {
	Repos []struct {
		Name   string `yaml:"name"`
		Repo   string `yaml:"repo"`
		Branch string `yaml:"branch"`
	}
}

func main() {
	input, err := ioutil.ReadFile("./fixtures/pipeline.yml")
	checkError(err)
	pipeline := string(input)

	tmpl := template.New("template")
	pipeline = preserveConfigValue(pipeline)
	fmt.Println(pipeline)
	tmpl, err = tmpl.Parse(pipeline)
	checkError(err)

	s := Structure{}
	err = yaml.Unmarshal([]byte(data), &s)
	checkError(err)

	var output bytes.Buffer
	err = tmpl.Execute(&output, s)
	checkError(err)
	pipeline = revertConfigValue(string(output.Bytes()))
	fmt.Println(pipeline)
}

func preserveConfigValue(input string) string {
	//keep config value e.g. {{my-private-key}}
	var re = regexp.MustCompile(`(\w*:\s*)({{)([^.][a-zA-Z_\-\s]*)(}})`)
	out := re.ReplaceAllString(input, "${1}(${3})")
	return out
}
func revertConfigValue(input string) string {
	var re = regexp.MustCompile(`(\w*:\s*)(\()(.*)(\))`)
	out := re.ReplaceAllString(input, "${1}{{${3}}}")
	return out
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
