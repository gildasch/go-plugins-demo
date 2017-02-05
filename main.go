package main

import (
	"log"
	"os"
	"path"
	"plugin"
)

func loadPlugin(path string) (func(string), error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}
	sayHello, err := p.Lookup("SayHello")
	if err != nil {
		return nil, err
	}
	return sayHello.(func(string)), nil
}

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

		p, err := loadPlugin(path.Join(pluginDir, pFile.Name()))
		if err != nil {
			log.Printf("Failed to load pluging %s", p)
		} else {
			ret = append(ret, p)
		}
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
