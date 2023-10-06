package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

// App struct
type App struct {
	ctx context.Context
}

type RandomImage struct {
	Message string
	Status  string
}

type AllBreeds struct {
	Message map[string]map[string][]string
	Status  string
}

type ImagesByBreed struct {
	Message []string
	Status  string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) GetRandomImageUrl() string {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data RandomImage
	json.Unmarshal(responseData, &data)

	return data.Message
}

func (a *App) GetBreedList() []string {
	var breeds []string

	response, err := http.Get("https://dog.ceo/api/breeds/list/all")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data AllBreeds
	json.Unmarshal(responseData, &data)

	for k := range data.Message {
		breeds = append(breeds, k)
	}

	sort.Strings(breeds)

	return breeds
}

func (a *App) GetImageUrlsByBreed(breed string) []string {

	url := fmt.Sprintf("%s%s%s%s", "https://dog.ceo/api/", "breed/", breed, "/images")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data ImagesByBreed
	json.Unmarshal(responseData, &data)

	return data.Message
}
