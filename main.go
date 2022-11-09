package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
)

var filepath string

func main() {
	flag.StringVar(&filepath, "t", "target", "文件路径，默认为空")
	flag.Parse()
	content, _ := ioutil.ReadFile(filepath)
	reg1 := regexp.MustCompile(`flag.*}`)
	result1 := reg1.FindAllString(string(content), -1)
	//fmt.Println(result1)
	for _, value := range result1 {
		fmt.Println(value)
	}
}
