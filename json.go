package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type details struct {
	Name    string    `json:"name"`
	Married bool      `json:"married"`
	Age     float64   `json:"age"`
	Address []Address `json:"address"`
	Marks   []int     `json:"marks"`
}
type Address struct {
	Street int    `json:"street"`
	City   string `json:"city"`
}

func main() {
	file, err := os.Open("details.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	jsonByteValues, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// var details details
	details := map[string]interface{}{}
	//json - key and value
	//key is a string
	//value - int, float, bool, string, json - interface

	json.Unmarshal(jsonByteValues, &details) //converting json data to struct
	log.Println(details)
	// fmt.Print(string(jsonByteValues)) //converting json data to string
	// jsonVal, err := json.Marshal(details)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// os.WriteFile("some.json",jsonVal)

}
func main4() {
	//var user map[int]string 
	user := map[int]interface{}{}
	user[1] = "name"
	user[2] = true
	// fmt.Println(user[2])
	//set
	users := map[string]bool{}
	users["golang"] = true

	// value, ok := users["rahul"]
	// if ok == false {
	// 	fmt.Println("value not there")
	// }

	//fmt.Println(users["rahul"])

}
