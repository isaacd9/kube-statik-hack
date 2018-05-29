package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/rakyll/statik/fs"

	_ "github.com/isaacd9/statiktest/statik"
)

type kubeModelProp struct {
	Type        string `json:"type"`
	Description string
}

type kubeModel struct {
	ID          string
	Description string
	Required    []string
	Properties  map[string]kubeModelProp
}

type kubeAPI struct {
	Models map[string]kubeModel
}

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	f, err := statikFS.Open("/batch_v1.json")
	if err != nil {
		log.Fatal(err)
	}

	d, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	api := &kubeAPI{}
	if err := json.Unmarshal(d, api); err != nil {
		log.Fatal(err)
	}

	for k, _ := range api.Models {
		log.Printf(
			"Model Found: %s",
			strings.Replace(k, ".", "/", -1),
		)

		//log.Printf("%v", v.ID)
		//log.Printf("%v", v.Description)
		//log.Printf("%v", v.Required)
		//for p, info := range v.Properties {
		//	//log.Printf("%v - %s - %s", snake.Snake(p), info.Type, info.Description)
		//}
	}
}
