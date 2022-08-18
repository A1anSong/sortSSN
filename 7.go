package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"os"
	"sortSSN/person"
	"strings"
)

func main() {
	file, err := os.Open("./7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	ssnError, err := os.OpenFile("./7errors.txt", os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("7errors.txt error", err)
	}
	defer ssnError.Close()
	writeErrors := bufio.NewWriter(ssnError)
	line := bufio.NewReader(file)
	var people []person.Person
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF {
			break
		}
		available := true
		info := strings.Split(string(content), "|")
		name := info[0]
		addr := strings.Trim(strings.Replace(info[1], ",", "", -1), " ")
		city := strings.Trim(strings.Replace(info[2], ",", "", -1), " ")
		state := info[3]
		zip := info[4]
		ssn := info[5]
		birth := info[6]
		// 姓名筛选规则
		if available {
			if len(strings.Split(name, " ")) != 2 {
				available = false
			}
		}
		// 地址筛选规则
		if available {
			if addr == "" {
				available = false
			}
		}
		// 城市筛选规则
		if available {
			if city == "" {
				available = false
			}
		}
		// 州筛选规则
		if available {
			if len(state) != 2 {
				available = false
			}
		}
		// 邮编筛选规则
		if available {
			if len(zip) > 5 || len(zip) < 4 {
				available = false
			}
		}
		// SSN筛选规则
		if available {
			if len(strings.Replace(ssn, "-", "", -1)) != 9 {
				available = false
			}
		}
		// 生日筛选规则
		if available {
			if birth == "" || len(strings.Split(strings.Replace(birth, "-", "/", -1), "/")) != 3 {
				available = false
			}
		}
		// 不符合筛选条件的写入errors.txt
		if !available {
			writeErrors.WriteString(string(content) + "\n")
			continue
		}
		// 符合条件的写入内存
		toTitle := cases.Title(language.English)
		var person = new(person.Person)
		person.Name = toTitle.String(name)
		person.Address = toTitle.String(addr)
		person.City = toTitle.String(city)
		person.State = strings.ToUpper(state)
		if len(zip) == 4 {
			zip = "0" + zip
		}
		person.Zip = zip
		person.SSN = strings.Replace(ssn, "-", "", -1)
		birthday := strings.Split(strings.Replace(birth, "-", "/", -1), "/")
		person.Day = birthday[1]
		person.Month = birthday[0]
		person.Year = birthday[2]
		people = append(people, *person)
	}
	writeErrors.Flush()
	// 成功的开始写入csv
	csv, err := os.OpenFile("./7.csv", os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("7.csv error", err)
	}
	defer ssnError.Close()
	writeCsv := bufio.NewWriter(csv)
	for _, person := range people {
		writeCsv.WriteString("\n," + person.Name +
			"," + person.Address +
			"," + person.City +
			"," + person.State +
			"," + person.Zip +
			"," + person.SSN +
			"," + person.Day +
			"," + person.Month +
			"," + person.Year)
	}
	writeCsv.Flush()
}
