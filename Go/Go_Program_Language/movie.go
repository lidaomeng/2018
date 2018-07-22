package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{"猛龙过江", 1980, false, []string{"李小龙", "邹文怀"}},
	{"少林足球", 2000, true, []string{"周星驰", "吴孟达", "张三"}},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)
}
