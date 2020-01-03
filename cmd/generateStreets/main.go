package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	res, err := http.Get("http://www.1562.kharkov.ua/uk/ppr/street/?q=&limit=10000")
	if err != nil {
		log.Fatalln("Could not get list")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Could not get list. Status=%d\n", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Could not read body")
	}
	bodyStr := string(body)
	re := regexp.MustCompile(`(\d+)\|([^\n]+)\n`)
	results := re.FindAllStringSubmatch(bodyStr, 1e6)
	fmt.Println("package my1562api")
	fmt.Println("")
	fmt.Println("// Streets list of all streets from 1562 database")
	fmt.Print("var Streets = StreetsList{\n")

	streetIds := make([]int, len(results))

	for i, item := range results {
		id, err := strconv.Atoi(item[1])
		if err != nil {
			log.Fatalf("Could not parse %s\n", item[1])
		}
		streetIds[i] = id
		name := strings.TrimSpace(item[2])
		fmt.Printf("    {%d, %#v},\n", id, name)
	}
	fmt.Print("}\n")

	fmt.Println("")
	fmt.Print("var streetIndexById = StreetIndexMap{\n")
	for i, id := range streetIds {
		fmt.Printf("    %d: %d,\n", id, i)
	}
	fmt.Print("}\n")

}
