package lib

import (
	"fmt"
	"io/ioutil"

	"github.com/goccy/go-yaml"
)

// RunYAML は ....
func RunYAML(jsonDir string, yamlFile string) error {

	// get list
	hosts, err := Ls(jsonDir)
	if err != nil {
		return err
	}
	if len(hosts) == 0 {
		return fmt.Errorf("No JSON files")
	}

	// loop by hosts
	m := make(map[string][]string, len(hosts))

	for _, host := range hosts {
		pkgs, err2 := ReadJSON(Host2File(host, jsonDir))
		if err2 != nil {
			return err2
		}
		m[host] = pkgs
	}

	// object to YAML
	bytes, err := yaml.Marshal(m)
	if err != nil {
		return err
	}

	// write YAML file
	return ioutil.WriteFile(yamlFile, bytes, 0644)
}
