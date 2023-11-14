package config

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"myhealthcheckapp/dtos"

	"gopkg.in/yaml.v2"
)

const (
	DefaultMethod = "GET"
	DefaultBody   = ""
)

var DefaultHeaders = map[string]string{}

func ReadConfig(filePath string) ([]dtos.Endpoint, error) {
	var endpoints []dtos.Endpoint

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading input file")
		return endpoints, err
	}

	err = yaml.Unmarshal(data, &endpoints)
	if err != nil {
		return endpoints, err
	}

	//set default values
	for i := range endpoints {
		if endpoints[i].Method == "" {
			endpoints[i].Method = DefaultMethod
		}

		if endpoints[i].Headers == nil {
			endpoints[i].Headers = DefaultHeaders
		}

		if endpoints[i].Body == "" {
			endpoints[i].Body = DefaultBody
		}

		parsedUrl, err := url.Parse(endpoints[i].URL)
		if err != nil {
			fmt.Println("Invalid URL")
		}
		endpoints[i].Domain = parsedUrl.Hostname()

	}

	return endpoints, nil

}
