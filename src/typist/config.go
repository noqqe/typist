package typist

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Challenges struct {
	Description string   `description`
	Lines       []string `challenges`
}

// Read formatted yaml file
func (c *Challenges) ReadFile(path string) *Challenges {

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		LogMessage("Could not open file", "red")
		os.Exit(1)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		LogMessage(fmt.Sprintf("Unmarshal %v", err), "red")
		os.Exit(1)
	}

	return c
}
