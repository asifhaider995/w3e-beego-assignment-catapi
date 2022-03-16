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

type CatCategoryController struct {
	beego.Controller
}

type Categories struct {

	// The `json` struct tag maps between the json name
	// and actual name of the field
	Name string `json:"name"`
}

type CategoryId struct {

	// The `json` struct tag maps between the json name
	// and actual name of the field
	Id int32 `json:"id"`
}

// This functions parses cat categories
func parseCatCategories(jsonBuffer []byte) ([]string, error) {

	// We create an empty array
	var cats []Categories

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

// This functions parses category ids
func parseCatCategoryIds(jsonBuffer []byte) ([]int32, error) {

	// We create an empty array
	var catIds []CategoryId

	// Unmarshal the json into it. this will use the struct tag
	//log.Println(jsonBuffer)
	err := json.Unmarshal(jsonBuffer, &catIds)
	if err != nil {
		return nil, err
	}
	x := len(catIds)
	i := 0
	var catSlice = make([]int32, x)
	for ; i < x; i++ {
		//log.Println(reflect.TypeOf(string(cats[i].Name)).Kind())
		//log.Println(strings.Trim(string(cats[i].Name), "{}"))
		//catSlice[i] = strings.Trim(string(catIds[i].Id), "{}")
		catSlice[i] = catIds[i].Id
	}

	return catSlice, nil

}

// This function delivers slice of cat categories
func getCategories() []string {
	categories, err := http.Get("https://api.thecatapi.com/v1/categories?limit=30")
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(categories.Body)
	catBreeds, err := parseCatCategories(body)
	return catBreeds
}

// This function delivers slice of cat category ids
func getCategoryIds() []int32 {
	categories, err := http.Get("https://api.thecatapi.com/v1/categories")
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(categories.Body)
	catBreedIds, err := parseCatCategoryIds(body)
	return catBreedIds
}

func (c *CatCategoryController) Get() {
	c.Data["Website"] = "thecatapi.com"
	c.Data["Email"] = "astaxie@gmail.com"

	c.Data["Categories"] = getCategories()
	c.Data["categoryName"] = ""
	c.Data["infoLengthCat"] = 0
	c.TplName = "category.tpl"
}

// Function to get cat by Category
func getCatByCategory(cId int32, mime string) string {
	var id string = strconv.Itoa(int(cId))
	cats, err := http.Get("https://api.thecatapi.com/v1/images/search?category_ids=" + id + "&mime_types=" + mime)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(cats.Body)
	return string(body)
}

func (c *CatCategoryController) Post() {
	c.TplName = "category.tpl"
	c.Data["Website"] = "thecatapi.com"
	c.Data["Email"] = "astaxie@gmail.com"
	// dealing with categories
	category := c.GetString("category")
	mime := c.GetString("mime_type")
	categories := getCategories()
	categoryIds := getCategoryIds()
	log.Println("Mime Type: ", mime)
	log.Println(category)
	log.Println(categories)
	log.Println(categoryIds)
	//c.Data["infoLengthCat"] = 0
	// index of category
	catInd, err := strconv.Atoi(category)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(categories[catInd])
	log.Println("category Id: ", categoryIds[catInd])
	//
	//// data payloads
	c.Data["Categories"] = categories
	c.Data["category"] = category
	c.Data["categoryName"] = categories[catInd]
	c.Data["infoCat"] = getCatByCategory(categoryIds[catInd], mime)
	c.Data["infoLengthCat"] = len(getCatByCategory(categoryIds[catInd], mime))

}
