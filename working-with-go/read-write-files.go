package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {
	file_name := "./create.sql"

	content, err := ioutil.ReadFile(file_name)
	if err != nil {
		log.Fatalln("Error reading file: ", file_name)
	}

	fmt.Println(string(content))

	if _, err := os.Stat("junk.foo"); os.IsNotExist(err) {
		fmt.Println(">>")
		fmt.Println("File: junk.foo does not exists!")
	}

	outfile := "outfile.sql"
	err = ioutil.WriteFile(outfile, content, 0644)
	if err != nil {
		log.Fatalln("Error writing file: ", err)
	} else {
		fmt.Println(">>")
		fmt.Println("Created: outfile.sql")
	}

	af, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error appending to file: ", err)
	}
	defer af.Close()
	if _, err = af.WriteString("Appending this text"); err != nil {
		log.Fatalln("Error writing to file: ", err)
	}
}