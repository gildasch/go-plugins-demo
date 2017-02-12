package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"path"
	"plugin"
)

type Transformer struct {
	name     string
	t        func(image.Image) image.Image
	priority int // priority can be 0 or 1
}

func loadPlugin(path string) (*Transformer, error) {
	fmt.Println("Loading", path)
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	transformSymbol, err := p.Lookup("Transform")
	if err != nil {
		return nil, err
	}
	transform, ok := transformSymbol.(func(image.Image) image.Image)
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s: Transform not a func(image.Image) image.Image", path))
	}

	var priority int
	prioritySymbol, err := p.Lookup("Priority")
	if err != nil {
		fmt.Println("Error looking up Priority:", err)
	} else {
		var ok bool
		priority, ok = prioritySymbol.(int)
		if !ok {
			fmt.Println("Error casting Priority to int")
		}
	}
	return &Transformer{path, transform, priority}, nil
}

func getPlugins(pluginDir string) (map[int][]*Transformer, error) {
	pDir, err := os.Open(pluginDir)
	if err != nil {
		return nil, err
	}
	pFiles, err := pDir.Readdir(0)
	if err != nil {
		return nil, err
	}

	ret := make(map[int][]*Transformer)
	for _, pFile := range pFiles {
		if pFile.IsDir() {
			continue
		}

		p, err := loadPlugin(path.Join(pluginDir, pFile.Name()))
		if err != nil {
			log.Printf("Failed to load pluging %s", p)
		} else {
			ret[p.priority] = append(ret[p.priority], p)
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

	for _, p := range plugins[1] {
		fmt.Println("Apply", p.name)
		i = p.t(i)
	}
	for _, p := range plugins[0] {
		fmt.Println("Apply", p.name)
		i = p.t(i)
	}

	jpeg.Encode(w, i, nil)
}

var pluginDir string

func main() {
	pluginDir = os.Args[1]

	http.HandleFunc("/a.jpg", handler)
	http.ListenAndServe(":8080", nil)
}
