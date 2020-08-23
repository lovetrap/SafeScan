package main

import (
	"SafeScan/config"
	"SafeScan/dom"
	"SafeScan/yaml"
	"io/ioutil"
	"net/url"
	"regexp"
	"strings"
	"sync"
)

var urls []string
var fuzzurl []string

func main() {
	config.NewConfig()
	uri := "http://127.0.0.1:8080"
	var html string
	dom.CdpBody(uri, &html)
	dir := "./plugins/dirscan/"

	pocYaml, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, fi := range pocYaml {
		if !fi.IsDir() {
			wg.Add(2)
			go yaml.DirScan(dir + fi.Name(), &wg)
			go yaml.PocCheck(dir + fi.Name(), &wg)
		}
	}
	wg.Wait()

	//yaml.Parse(dir + "wordpress-ext-mailpress-rce.yml")

	//var wg sync.WaitGroup
	//wg.Add(2)
	//go findUrls(uri, "[href]", "href", &wg)
	//wg.Wait()
	//for _, value := range urls {
	//	fmt.Println(value)
	//}
}

func findUrls(uri,selectors,attr string, wg *sync.WaitGroup) {
	defer wg.Done()
	Dom := dom.NewGobySelector(uri)
	URL, _ := url.Parse(uri)
	for _, value := range Dom.FindLink(selectors,attr) {
		var temp string
		if reg := regexp.MustCompile(`[\w-]+=((?i)[\w%]*)`).FindAllStringSubmatch(value, -1); reg != nil {
			for _, k := range reg {
				temp = strings.ReplaceAll(value, k[1], "")
			}
		}
		if checkArray(fuzzurl, temp) {
			fuzzurl = append(fuzzurl, temp)
			if strings.Contains(uri, URL.Host) && checkArray(urls, value) && regexp.MustCompile(`(?i)(\.doc|\.xml|\.xls|\.rtf|\.ppt)`).FindString(value) == ""{
				urls = append(urls, value)
				wg.Add(1)
				findUrls(value, selectors, attr, wg)
			}
		}
	}
}

func checkArray(a []string, str string) bool{

	for _, value := range a {
		if strings.Contains(value, str) {
			return false
		}
	}
	return true
}
