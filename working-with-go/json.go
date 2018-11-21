package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name, City string
}

//重新定义输出格式
func (p Person) String() string {
	return fmt.Sprintf("%s:%s", p.Name, p.City)
}

var person Person

var people []Person

func main() {
	//json
	json_str := "{\"name\": \"Sunny\", \"city\": \"China\"}"
	//解析json，到person
	if err := json.Unmarshal([]byte(json_str), &person); err != nil {
		fmt.Println("Error Reading File ", err)
	}
	//打印解析后结构体
	fmt.Printf("Name: %v, City: %v \n", person.Name, person.City)

	//定义一个map
	peopleMap := map[string]string{
		"Jack":     "USA",
		"Sunny":    "UK",
		"Simmon":   "China",
		"Shawnpai": "German",
	}
	//写入json
	people_json, err := json.Marshal(peopleMap)
	if err != nil {
		fmt.Println("Marshal People Error: ", err)
	}
	fmt.Printf("(Not String) People Json is %v\n", people_json)
	fmt.Printf("People Json is %v\n", string(people_json))

	//写入文件
	err = ioutil.WriteFile("./people_map.json", people_json, 0644)
	if err != nil {
		fmt.Println("Write File Err: ", err)
	}

	//创建你一个文件
	f, err := os.Create("./people.json")
	if err != nil {
		fmt.Println("Create Err", err)
	}
	//关闭
	defer f.Close()

	//将key-value放入array
	var personList []Person
	for k, v := range peopleMap {
		personList = append(personList, Person{k, v})
	}
	//转化为json
	personListJson, err := json.Marshal(personList)
	if err != nil {
		fmt.Println("Marshal Error", err)
	}
	//将json写入文件
	f.Write(personListJson)

	//读取文件
	file, err := ioutil.ReadFile("./people.json")
	if err != nil {
		fmt.Println("Read Json Err", err)
	}
	//解析
	if err = json.Unmarshal(file, &people); err != nil {
		fmt.Println("Unmarshal File err", err)
	}
	fmt.Printf("People Json are %v\n", people)

}
