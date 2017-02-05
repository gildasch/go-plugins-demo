package main

import (
	"os"
	"path"
	"plugin"
)

func getPlugins(pluginDir string) ([]func(string), error) {
	pDir, err := os.Open(pluginDir)
	if err != nil {
		return nil, err
	}
	pFiles, err := pDir.Readdir(0)
	if err != nil {
		return nil, err
	}

	var ret []func(string)
	for _, pFile := range pFiles {
		if pFile.IsDir() {
			continue
		}
		p, err := plugin.Open(path.Join(pluginDir, pFile.Name()))
		if err != nil {
			return nil, err
		}
		sayHello, err := p.Lookup("SayHello")
		if err != nil {
			return nil, err
		}
		ret = append(ret, sayHello.(func(string)))
	}
	return ret, nil
}

func main() {
	pluginDir := os.Args[1]

	ps, err := getPlugins(pluginDir)
	if err != nil {
		panic(err)
	}
	for _, p := range ps {
		p("Gildas")
	}
}
