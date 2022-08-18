package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sortSSN/person"
)

func main() {
	file, err := os.Open("./originalSSN.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//file3, err := os.OpenFile("./3.txt", os.O_APPEND|os.O_CREATE, 0660)
	//if err != nil {
	//	fmt.Println("3.txt error", err)
	//}
	//defer file3.Close()
	//write3 := bufio.NewWriter(file3)
	line := bufio.NewReader(file)
	var people []person.Person
	for {
		_, _, err := line.ReadLine()
		if err == io.EOF {
			break
		}
		//info := strings.Split(string(content), "|")
		//if len(info) == 3 {
		//	write3.WriteString(string(content) + "\n")
		//}
		var person = new(person.Person)
		people = append(people, *person)
	}
	//write3.Flush()
	fmt.Println(len(people))
}
