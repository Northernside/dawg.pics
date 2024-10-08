package main

import (
	"io/ioutil"
	"log"

	"dawg.pics/api"
	"dawg.pics/api/routes"
	"dawg.pics/modules/env"
)

func main() {
	env.LoadEnvFile()

	loadImages()
	api.StartWebServer()
}

func loadImages() {
	files, err := ioutil.ReadDir(env.GetEnv("PICTURES_DIR"))
	if err != nil {
		panic(err)
	}

	log.Printf("Loading %d images into cache\n", len(files))

	index := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		data, err := ioutil.ReadFile(env.GetEnv("PICTURES_DIR") + "/" + file.Name())
		if err != nil {
			panic(err)
		}

		routes.FileCache[index] = data
		index++
	}

	log.Println("Images loaded into cache")
}
