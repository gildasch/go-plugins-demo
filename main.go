package main

import (
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"path"
	"plugin"
)

func loadPlugin(path string) (func(image.Image) image.Image, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}
	transform, err := p.Lookup("Transform")
	if err != nil {
		return nil, err
	}
	return transform.(func(image.Image) image.Image), nil
}

func getPlugins(pluginDir string) ([]func(image.Image) image.Image, error) {
	pDir, err := os.Open(pluginDir)
	if err != nil {
		return nil, err
	}
	pFiles, err := pDir.Readdir(0)
	if err != nil {
		return nil, err
	}

	var ret []func(image.Image) image.Image
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

func handler(w http.ResponseWriter, r *http.Request) {
	fi, _ := os.Open("lca.jpg")
	i, _ := jpeg.Decode(fi)

	plugins, err := getPlugins(pluginDir)
	if err != nil {
		panic(err)
	}

	for _, p := range plugins {
		i = p(i)
	}

	jpeg.Encode(w, i, nil)
}

var pluginDir string

func main() {
	pluginDir = os.Args[1]

	http.HandleFunc("/a.jpg", handler)
	http.ListenAndServe(":8080", nil)
}
