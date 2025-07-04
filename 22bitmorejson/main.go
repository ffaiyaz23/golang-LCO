package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              // this field will not be reflected
	Tags     []string `json:"tags,omitempty"` //omitempty says if value is nil just dont throw that
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs", 599, "Learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern", 299, "Learncodeonline.in", "def123", []string{"fullstack", "js"}},
		{"Angular", 199, "Learncodeonline.in", "ghi123", nil},
	}

	finalJson, _ := json.MarshalIndent(lcoCourses, "", "\t")

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "coursename": "ReactJs",
                "price": 599,
                "website": "Learncodeonline.in",
                "tags": ["web-dev", "js"]
        }
	`)

	// to check if is json data is valid or not
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {

		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) // passing refernce just to make sure we are not sending copy (could work without & also)
		fmt.Printf("%#v\n", lcoCourse)              // %#v to print interfaces

	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// Some cases where you just want to add data to key value
	var myonlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlineData)
	fmt.Printf("%#v\n", myonlineData)

	for k, v := range myonlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}

func main() {
	fmt.Println("Welcome to json in Golang")

	// EncodeJson()
	DecodeJson()
}
