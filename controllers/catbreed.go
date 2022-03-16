package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CatBreedController struct {
	beego.Controller
}

type Cat struct {

	// The `json` struct tag maps between the json name
	// and actual name of the field
	Name string `json:"name"`
}

type CatId struct {

	// The `json` struct tag maps between the json name
	// and actual name of the field
	Id string `json:"id"`
}

// This functions accepts a byte array containing a JSON
func parseCatBreeds(jsonBuffer []byte) ([]string, error) {

	// We create an empty array
	var cats []Cat

	// Unmarshal the json into it. this will use the struct tag
	//log.Println(jsonBuffer)
	err := json.Unmarshal(jsonBuffer, &cats)
	if err != nil {
		return nil, err
	}
	x := len(cats)
	i := 0
	var catSlice = make([]string, x)
	for ; i < x; i++ {
		//log.Println(reflect.TypeOf(string(cats[i].Name)).Kind())
		//log.Println(strings.Trim(string(cats[i].Name), "{}"))
		catSlice[i] = strings.Trim(string(cats[i].Name), "{}")
	}

	return catSlice, nil

}

// This functions accepts a byte array containing a JSON
func parseCatBreedIds(jsonBuffer []byte) ([]string, error) {

	// We create an empty array
	var catIds []CatId

	// Unmarshal the json into it. this will use the struct tag
	//log.Println(jsonBuffer)
	err := json.Unmarshal(jsonBuffer, &catIds)
	if err != nil {
		return nil, err
	}
	x := len(catIds)
	i := 0
	var catSlice = make([]string, x)
	for ; i < x; i++ {
		//log.Println(reflect.TypeOf(string(cats[i].Name)).Kind())
		//log.Println(strings.Trim(string(cats[i].Name), "{}"))
		catSlice[i] = strings.Trim(string(catIds[i].Id), "{}")
	}

	return catSlice, nil

}

// This function delivers slice of cat breeds
func getBreeds() []string {
	breeds, err := http.Get("https://api.thecatapi.com/v1/breeds?limit=30")
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(breeds.Body)
	catBreeds, err := parseCatBreeds(body)
	return catBreeds
}

// This function delivers slice of cat breeds ids
func getBreedIds() []string {
	breeds, err := http.Get("https://api.thecatapi.com/v1/breeds?limit=30")
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(breeds.Body)
	catBreedIds, err := parseCatBreedIds(body)
	return catBreedIds
}

func (c *CatBreedController) Get() {
	c.Data["Website"] = "thecatapi.com"
	c.Data["Email"] = "astaxie@gmail.com"

	//categories, err := http.Get("https://api.thecatapi.com/v1/categories?limit=30")

	//bodyCategories, err := ioutil.ReadAll(categories.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//	fmt.Println(err)
	//}

	//catCategories, err := parseCatCategories(bodyCategories)

	c.Data["Breeds"] = getBreeds()
	c.Data["breedName"] = ""
	c.Data["infoLength"] = 0
	c.TplName = "cat.tpl"
}

// Function to get cat by breed
func getCatByBreed(bId string) string {
	cats, err := http.Get("https://api.thecatapi.com/v1/breeds/search?name=" + bId)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(cats.Body)
	return string(body)
}

func (c *CatBreedController) Post() {
	c.TplName = "cat.tpl"
	c.Data["Website"] = "thecatapi.com"
	c.Data["Email"] = "astaxie@gmail.com"
	// Dealing with breeds
	breed := c.GetString("breed") // Getting from tpl
	breeds := getBreeds()
	breedIds := getBreedIds()

	// index of breed
	breedInd, err := strconv.Atoi(breed)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(breeds[breedInd])
	log.Println("breed Id: ", breedIds[breedInd])

	// data payloads
	c.Data["Breeds"] = breeds
	c.Data["breed"] = breed
	c.Data["breedName"] = breeds[breedInd]
	c.Data["info"] = getCatByBreed(breedIds[breedInd])
	c.Data["infoLength"] = len(getCatByBreed(breedIds[breedInd]))

}
