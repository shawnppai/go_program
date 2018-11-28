package main

import (
	"fmt"
	"regexp"
)

func basic_regexes() {
	pattern := "[0-9]+"
	str := "The 12 monkeys ate 48 bananas"
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error comliling regex", err)
	}

	if re.MatchString(str) {
		fmt.Println("Match")
	} else {
		fmt.Println("Not Match")
	}

	result_two := re.FindString(str)
	fmt.Println("Number Matched", result_two)

	result_three := re.FindAllString(str, 2)
	for i, v := range result_three {
		fmt.Printf("Match %d: %s\n", i, v)
	}

	result_four := re.ReplaceAllString(str, "xx")
	fmt.Println("After Replace", result_four)
}

func case_insensitive() {
	ptn := `(?i)^t.`
	str := "To be or not to be"
	re, err := regexp.Compile(ptn)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	result := re.FindString(str)
	fmt.Println("Result:", result)

}

func sub_matches() {
	str1 := "Hello @world@ Match"
	sub_re, _ := regexp.Compile("@(.*)@")
	m := sub_re.FindStringSubmatch(str1)

	if len(m) > 1 {
		fmt.Println(m[1])
	}

	str2 := "A [word] here and [there]"
    esc_pattern := `\[(.*?)\]`
    esc_re, _ := regexp.Compile(esc_pattern)

    fmt.Println(esc_re.FindStringSubmatch(str2))
    fmt.Println(esc_re.FindAllStringSubmatch(str2, -1))
}

func main() {
	basic_regexes()
	case_insensitive()
	sub_matches()
}
