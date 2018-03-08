package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	reader := csv.NewReader(bufio.NewReader(os.Stdin))
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var out struct {
		Title     string
		Questions []string
		Data      [][]string
	}

	out.Title = "kubernetes ingress user survey 2018"
	out.Questions = records[0]
	out.Data = records[1 : len(records)-1]

	bytes, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
